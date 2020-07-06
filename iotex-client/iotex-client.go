// Copyright (c) 2020 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package iotex_client

import (
	"context"
	"crypto/tls"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"sync"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"

	"github.com/iotexproject/iotex-core-rosetta-gateway/config"
)

const (
	rewardingProtocolID      = "rewarding"
	stakingProtocolID        = "staking"
	availableBalanceMethodID = "AvailableBalance"
	Transfer                 = "transfer"
	Execution                = "execution"
	DepositToRewardingFund   = "depositToRewardingFund"
	ClaimFromRewardingFund   = "claimFromRewardingFund"
	StakeCreate              = "stakeCreate"
	StakeWithdraw            = "stakeWithdraw"
	StakeAddDeposit          = "stakeAddDeposit"
	CandidateRegister        = "candidateRegister"
	StatusSuccess            = "success"
	StatusFail               = "fail"
	ActionTypeFee            = "fee"
	// NonceKey is the name of the key in the Metadata map inside a
	// ConstructionMetadataResponse that specifies the next valid nonce.
	NonceKey = "nonce"
)

var (
	RewardingAddress string
	StakingAddress   string
)

func init() {
	h := hash.Hash160b([]byte(rewardingProtocolID))
	addr, _ := address.FromBytes(h[:])
	RewardingAddress = addr.String()
	h = hash.Hash160b([]byte(stakingProtocolID))
	addr, _ = address.FromBytes(h[:])
	StakingAddress = addr.String()
}

type (
	// IoTexClient is the IoTex blockchain client interface.
	IoTexClient interface {
		// GetChainID returns the network chain context, derived from the
		// genesis document.
		GetChainID(ctx context.Context) (string, error)

		// GetBlock returns the IoTex block at given height.
		GetBlock(ctx context.Context, height int64) (*types.Block, error)

		// GetLatestBlock returns latest IoTex block.
		GetLatestBlock(ctx context.Context) (*types.Block, error)

		// GetGenesisBlock returns the IoTex genesis block.
		GetGenesisBlock(ctx context.Context) (*types.Block, error)

		// GetAccount returns the IoTex staking account for given owner address
		// at given height.
		GetAccount(ctx context.Context, height int64, owner string) (*types.AccountBalanceResponse, error)

		// SubmitTx submits the given encoded transaction to the node.
		SubmitTx(ctx context.Context, tx *iotextypes.Action) (txid string, err error)

		// GetStatus returns the status overview of the node.
		GetStatus(ctx context.Context) (*iotexapi.GetChainMetaResponse, error)

		// GetVersion returns the server's version.
		GetVersion(ctx context.Context) (*iotexapi.GetServerMetaResponse, error)

		// GetTransactions returns transactions of the block.
		GetTransactions(ctx context.Context, height int64) ([]*types.Transaction, error)

		// GetConfig returns the config.
		GetConfig() *config.Config
	}

	Account struct {
		Nonce   uint64
		Balance string
	}

	// grpcIoTexClient is an implementation of IoTexClient using gRPC.
	grpcIoTexClient struct {
		sync.RWMutex

		grpcConn *grpc.ClientConn
		client   iotexapi.APIServiceClient
		cfg      *config.Config
	}
)

// NewIoTexClient returns an implementation of IoTexClient
func NewIoTexClient(cfg *config.Config) (cli IoTexClient, err error) {
	gcli := &grpcIoTexClient{cfg: cfg}
	if err = gcli.connect(); err != nil {
		return
	}
	cli = gcli
	return
}

func (c *grpcIoTexClient) GetChainID(ctx context.Context) (string, error) {
	return c.cfg.NetworkIdentifier.Network, nil
}

func (c *grpcIoTexClient) GetBlock(ctx context.Context, height int64) (ret *types.Block, err error) {
	if err = c.connect(); err != nil {
		return
	}
	return c.getBlock(ctx, height)
}

func (c *grpcIoTexClient) GetLatestBlock(ctx context.Context) (*types.Block, error) {
	if err := c.connect(); err != nil {
		return nil, err
	}
	return c.getLatestBlock(ctx)
}

func (c *grpcIoTexClient) getLatestBlock(ctx context.Context) (*types.Block, error) {
	res, err := c.client.GetChainMeta(context.Background(), &iotexapi.GetChainMetaRequest{})
	if err != nil {
		return nil, err
	}
	return c.getBlock(ctx, int64(res.ChainMeta.Height))
}

func (c *grpcIoTexClient) GetGenesisBlock(ctx context.Context) (*types.Block, error) {
	if err := c.connect(); err != nil {
		return nil, err
	}
	return c.getBlock(ctx, 1)
}

func (c *grpcIoTexClient) GetAccount(ctx context.Context, height int64, owner string) (ret *types.AccountBalanceResponse, err error) {
	if err = c.connect(); err != nil {
		return
	}

	if owner == RewardingAddress {
		return c.getRewardingAccount(ctx, height)
	}

	request := &iotexapi.GetAccountRequest{Address: owner}
	resp, err := c.client.GetAccount(ctx, request)
	if err != nil {
		return
	}
	acc := resp.GetAccountMeta()
	blk := resp.GetBlockIdentifier()
	ret = &types.AccountBalanceResponse{
		BlockIdentifier: &types.BlockIdentifier{
			Index: int64(blk.GetHeight()),
			Hash:  blk.GetHash(),
		},
		Balances: []*types.Amount{{
			Value: acc.GetBalance(),
			Currency: &types.Currency{
				Symbol:   c.cfg.Currency.Symbol,
				Decimals: c.cfg.Currency.Decimals,
				Metadata: nil,
			}}},
		Metadata: &map[string]interface{}{NonceKey: acc.GetPendingNonce()},
	}
	return
}

func (c *grpcIoTexClient) getRewardingAccount(ctx context.Context, height int64) (ret *types.AccountBalanceResponse, err error) {
	// call readState
	out, err := c.client.ReadState(context.Background(), &iotexapi.ReadStateRequest{
		ProtocolID: []byte(rewardingProtocolID),
		MethodName: []byte(availableBalanceMethodID),
		Arguments:  nil,
	})
	if err != nil {
		return
	}
	val, ok := big.NewInt(0).SetString(string(out.Data), 10)
	if !ok {
		err = errors.New("balance convert error")
		return
	}
	blk, err := c.getLatestBlock(ctx)
	if err != nil {
		return
	}

	ret = &types.AccountBalanceResponse{
		BlockIdentifier: blk.BlockIdentifier,
		Balances: []*types.Amount{{
			Value: val.String(),
			Currency: &types.Currency{
				Symbol:   c.cfg.Currency.Symbol,
				Decimals: c.cfg.Currency.Decimals,
				Metadata: nil,
			},
		},
		},
		Metadata: &map[string]interface{}{NonceKey: 0},
	}
	return
}

func (c *grpcIoTexClient) GetTransactions(ctx context.Context, height int64) (ret []*types.Transaction, err error) {
	if err = c.connect(); err != nil {
		return
	}
	actionMap, receiptMap, hashSlice, err := c.getRawBlock(ctx, height)
	if err != nil {
		return
	}

	// get systemlog by height and store to map
	hashToSystemLog, err := c.getSystemLog(ctx, height)
	if err != nil {
		return
	}

	ret = make([]*types.Transaction, 0)
	for _, h := range hashSlice {
		act := actionMap[h]
		if act.GetCore().GetGrantReward() != nil {
			// gas fee and amount both 0
			continue
		}
		r, ok := receiptMap[h]
		if !ok {
			err = errors.New(fmt.Sprintf("failed find receipt:%s", hex.EncodeToString(h[:])))
			return
		}
		decode, err := c.decodeAction(ctx, act, hex.EncodeToString(h[:]), r, hashToSystemLog)
		if err != nil {
			return nil, err
		}
		if decode != nil {
			ret = append(ret, decode)
		}
	}
	return
}

func (c *grpcIoTexClient) SubmitTx(ctx context.Context, tx *iotextypes.Action) (txid string, err error) {
	if err = c.connect(); err != nil {
		return
	}
	ret, err := c.client.SendAction(ctx, &iotexapi.SendActionRequest{Action: tx})
	if err != nil {
		return
	}
	txid = ret.ActionHash
	return
}

func (c *grpcIoTexClient) GetStatus(ctx context.Context) (*iotexapi.GetChainMetaResponse, error) {
	err := c.connect()
	if err != nil {
		return nil, err
	}
	return c.client.GetChainMeta(context.Background(), &iotexapi.GetChainMetaRequest{})
}

func (c *grpcIoTexClient) GetVersion(ctx context.Context) (*iotexapi.GetServerMetaResponse, error) {
	err := c.connect()
	if err != nil {
		return nil, err
	}
	return c.client.GetServerMeta(ctx, &iotexapi.GetServerMetaRequest{})
}

func (c *grpcIoTexClient) GetConfig() *config.Config {
	return c.cfg
}

func (c *grpcIoTexClient) connect() (err error) {
	c.Lock()
	defer c.Unlock()
	// Check if the existing connection is good.
	if c.grpcConn != nil && c.grpcConn.GetState() != connectivity.Shutdown {
		return
	}
	opts := []grpc.DialOption{}
	if c.cfg.Server.SecureEndpoint {
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	c.grpcConn, err = grpc.Dial(c.cfg.Server.Endpoint, opts...)
	c.client = iotexapi.NewAPIServiceClient(c.grpcConn)
	return err
}

func (c *grpcIoTexClient) getBlock(ctx context.Context, height int64) (ret *types.Block, err error) {
	var parentHeight uint64
	if height <= 1 {
		parentHeight = 1
	} else {
		parentHeight = uint64(height) - 1
	}
	count := uint64(2)
	if parentHeight == uint64(height) {
		count = 1
	}
	request := &iotexapi.GetBlockMetasRequest{
		Lookup: &iotexapi.GetBlockMetasRequest_ByIndex{
			ByIndex: &iotexapi.GetBlockMetasByIndexRequest{
				Start: parentHeight,
				Count: count,
			},
		},
	}
	resp, err := c.client.GetBlockMetas(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(resp.BlkMetas) == 0 {
		return nil, errors.New("not found")
	}
	var blk, parentBlk *iotextypes.BlockMeta
	if len(resp.BlkMetas) == 2 {
		blk = resp.BlkMetas[1]
		parentBlk = resp.BlkMetas[0]
	} else {
		blk = resp.BlkMetas[0]
		parentBlk = resp.BlkMetas[0]
	}
	ret = &types.Block{
		BlockIdentifier: &types.BlockIdentifier{
			Index: int64(blk.Height),
			Hash:  blk.Hash,
		},
		ParentBlockIdentifier: &types.BlockIdentifier{
			Index: int64(parentHeight),
			Hash:  parentBlk.Hash,
		},
		Timestamp: blk.Timestamp.Seconds * 1e3, // ms,
	}
	return
}

func (c *grpcIoTexClient) getRawBlock(ctx context.Context, height int64) (actionMap map[hash.Hash256]*iotextypes.Action, receiptMap map[hash.Hash256]*iotextypes.Receipt, hashSlice []hash.Hash256, err error) {
	getRawBlocksRes, err := c.client.GetRawBlocks(ctx, &iotexapi.GetRawBlocksRequest{
		StartHeight:  uint64(height),
		Count:        1,
		WithReceipts: true,
	})
	if err != nil || len(getRawBlocksRes.GetBlocks()) != 1 {
		return
	}

	actionMap = make(map[hash.Hash256]*iotextypes.Action)
	receiptMap = make(map[hash.Hash256]*iotextypes.Receipt)
	// hashSlice for fixed sequence,b/c map is unordered
	hashSlice = make([]hash.Hash256, 0)
	blk := getRawBlocksRes.GetBlocks()[0]
	for _, act := range blk.GetBlock().GetBody().GetActions() {
		var pro []byte
		pro, err = proto.Marshal(act)
		if err != nil {
			return
		}
		actionMap[hash.Hash256b(pro)] = act
		hashSlice = append(hashSlice, hash.Hash256b(pro))
	}
	for _, receipt := range blk.GetReceipts() {
		receiptMap[hash.BytesToHash256(receipt.ActHash)] = receipt
	}
	return
}

func (c *grpcIoTexClient) getSystemLog(ctx context.Context, height int64) (hashToSystemLog map[string][]*iotextypes.ActionSystemLog_Transaction, err error) {
	hashToSystemLog = make(map[string][]*iotextypes.ActionSystemLog_Transaction)
	systemlog, err := c.client.GetSystemLogByBlockHeight(ctx, &iotexapi.GetSystemLogByBlockHeightRequest{BlockHeight: uint64(height)})
	if err == nil && len(systemlog.GetBlockSystemLog().GetActionSystemLog()) != 0 {
		for _, a := range systemlog.GetBlockSystemLog().GetActionSystemLog() {
			hashToSystemLog[hex.EncodeToString(a.GetActionHash())] = a.GetTransactions()
		}
	}
	return
}

func (c *grpcIoTexClient) decodeAction(ctx context.Context, act *iotextypes.Action, h string, receipt *iotextypes.Receipt, systemLog map[string][]*iotextypes.ActionSystemLog_Transaction) (ret *types.Transaction, err error) {
	srcPub, err := crypto.BytesToPublicKey(act.GetSenderPubKey())
	if err != nil {
		return
	}
	callerAddr, err := address.FromBytes(srcPub.Hash())
	if err != nil {
		return
	}
	ret, status, err := c.gasFeeAndStatus(callerAddr, act, h, receipt)
	if err != nil {
		return
	}

	// handle execution action
	if act.GetCore().GetExecution() != nil {
		// get contract address generated of this action hash
		err = c.handleExecution(ctx, ret, act, h, callerAddr, status, systemLog)
		return
	}

	amount, senderSign, actionType, dst, err := assertAction(act)
	if err != nil || amount == "" || actionType == "" {
		return
	}

	// handle general action that is not stake withdraw
	if actionType != StakeWithdraw {
		// for general action,if amount is 0 just return
		err = c.handleGeneralAction(ret, callerAddr, amount, senderSign, actionType, dst, status)
		return
	}

	// handle stake withdraw action
	trans, ok := systemLog[h]
	if !ok {
		return nil, errors.New("stake withdraw action but have no systemlog")
	}
	err = c.handleSystemlog(ret, trans, StakeWithdraw, status, 2)
	return
}

func (c *grpcIoTexClient) handleGeneralAction(ret *types.Transaction, callerAddr address.Address, amount, senderSign, actionType, dst, status string) error {
	if amount == "0" {
		return nil
	}
	senderAmountWithSign := amount
	dstAmountWithSign := amount
	if senderSign == "-" {
		senderAmountWithSign = senderSign + amount
	} else {
		dstAmountWithSign = "-" + amount
	}

	aal := addressAmountList{{callerAddr.String(), senderAmountWithSign}}
	if dst != "" {
		aal = append(aal, &addressAmount{dst, dstAmountWithSign})
	}
	return c.addOperation(ret, aal, actionType, status, 2)
}

func (c *grpcIoTexClient) handleExecutionAmount(act *iotextypes.Action,
	contractAddr string, callerAddr address.Address) (aal addressAmountList, err error) {
	amount := act.GetCore().GetExecution().GetAmount()
	if amount == "0" {
		return
	}
	// deal with pure transfer to contract address
	aal = addressAmountList{{callerAddr.String(), "-" + amount}, {contractAddr, amount}}
	return
}

func (c *grpcIoTexClient) handleSystemlog(ret *types.Transaction,
	transaction []*iotextypes.ActionSystemLog_Transaction, actionType, status string, startIndex int64) (err error) {
	var aal addressAmountList
	for _, trans := range transaction {
		amount := trans.GetAmount()
		if amount == "0" {
			continue
		}
		aal = append(aal, addressAmountList{{trans.Sender, "-" + amount}, {trans.Recipient, amount}}...)
	}
	return c.addOperation(ret, aal, actionType, status, startIndex)
}

func (c *grpcIoTexClient) handleExecution(ctx context.Context, ret *types.Transaction, act *iotextypes.Action, h string, callerAddr address.Address, status string,
	systemLog map[string][]*iotextypes.ActionSystemLog_Transaction) (err error) {
	contractAddr := act.GetCore().GetExecution().GetContract()
	if contractAddr == "" {
		contractAddr, err = getContractAddress(ctx, h, c.client)
		if err != nil {
			return
		}
	}
	amount := act.GetCore().GetExecution().GetAmount()
	// deal with pure transfer to contract address
	var aal addressAmountList
	if amount != "0" {
		aal = addressAmountList{{callerAddr.String(), "-" + amount}, {contractAddr, amount}}
	}

	// if there's systemlog,the above aal is included
	if trans, ok := systemLog[h]; ok {
		return c.handleSystemlog(ret, trans, Execution, status, 2)
	}
	return c.addOperation(ret, aal, Execution, status, 2)
}

func (c *grpcIoTexClient) gasFeeAndStatus(callerAddr address.Address, act *iotextypes.Action, h string, receipt *iotextypes.Receipt) (ret *types.Transaction, status string, err error) {
	status = StatusSuccess
	if receipt.GetStatus() != 1 {
		status = StatusFail
	}
	gasConsumed := new(big.Int).SetUint64(receipt.GetGasConsumed())
	gasPrice, ok := new(big.Int).SetString(act.GetCore().GetGasPrice(), 10)
	if !ok {
		err = errors.New("convert gas price error")
		return
	}
	gasFee := gasPrice.Mul(gasPrice, gasConsumed)
	amount := gasFee.String()
	// if gasFee is not 0
	if gasFee.Sign() == 1 {
		amount = "-" + amount
	}
	ret = &types.Transaction{TransactionIdentifier: &types.TransactionIdentifier{h}}
	aal := addressAmountList{{callerAddr.String(), amount}, {RewardingAddress, gasFee.String()}}
	err = c.addOperation(ret, aal, ActionTypeFee, status, 0)
	return
}

func (c *grpcIoTexClient) addOperation(ret *types.Transaction, amountList addressAmountList, actionType, status string, startIndex int64) error {
	sort.Sort(amountList)
	var oper []*types.Operation
	for _, s := range amountList {
		oper = append(oper, &types.Operation{
			OperationIdentifier: &types.OperationIdentifier{
				Index:        startIndex,
				NetworkIndex: nil,
			},
			RelatedOperations: nil,
			Type:              actionType,
			Status:            status,
			Account: &types.AccountIdentifier{
				Address:    s.address,
				SubAccount: nil,
				Metadata:   nil,
			},
			Amount: &types.Amount{
				Value: s.amount,
				Currency: &types.Currency{
					Symbol:   c.cfg.Currency.Symbol,
					Decimals: c.cfg.Currency.Decimals,
					Metadata: nil,
				},
				Metadata: nil,
			},
			Metadata: nil,
		})
		startIndex++
	}
	ret.Operations = append(ret.Operations, oper...)
	return nil
}

func assertAction(act *iotextypes.Action) (amount, senderSign, actionType, dst string, err error) {
	amount = "0"
	senderSign = "-"
	switch {
	case act.GetCore().GetTransfer() != nil:
		actionType = Transfer
		amount = act.GetCore().GetTransfer().GetAmount()
		dst = act.GetCore().GetTransfer().GetRecipient()
	case act.GetCore().GetDepositToRewardingFund() != nil:
		actionType = DepositToRewardingFund
		amount = act.GetCore().GetDepositToRewardingFund().GetAmount()
		dst = RewardingAddress
	case act.GetCore().GetClaimFromRewardingFund() != nil:
		actionType = ClaimFromRewardingFund
		amount = act.GetCore().GetClaimFromRewardingFund().GetAmount()
		senderSign = "+"
		dst = RewardingAddress
	case act.GetCore().GetStakeAddDeposit() != nil:
		actionType = StakeAddDeposit
		amount = act.GetCore().GetStakeAddDeposit().GetAmount()
		dst = StakingAddress
	case act.GetCore().GetStakeCreate() != nil:
		actionType = StakeCreate
		amount = act.GetCore().GetStakeCreate().GetStakedAmount()
		dst = StakingAddress
	case act.GetCore().GetStakeWithdraw() != nil:
		// TODO need to add amount when it's available on iotex-core
		actionType = StakeWithdraw
	case act.GetCore().GetCandidateRegister() != nil:
		actionType = CandidateRegister
		amount = act.GetCore().GetCandidateRegister().GetStakedAmount()
		dst = StakingAddress
	}
	return
}

func getContractAddress(ctx context.Context, h string, client iotexapi.APIServiceClient) (contractAddr string, err error) {
	// need to get contract address generated of this action hash
	responseReceipt, err := client.GetReceiptByAction(ctx, &iotexapi.GetReceiptByActionRequest{ActionHash: h})
	if err != nil {
		return
	}
	contractAddr = responseReceipt.GetReceiptInfo().GetReceipt().GetContractAddress()
	return
}
