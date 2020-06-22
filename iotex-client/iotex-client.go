package iotex_client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"sort"
	"sync"

	"github.com/coinbase/rosetta-sdk-go/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
)

const (
	// Transfer action type
	Transfer = "transfer"
	// Execution action type
	Execution = "execution"
	// DepositToRewardingFund action type
	DepositToRewardingFund = "depositToRewardingFund"
	// ClaimFromRewardingFund action type
	ClaimFromRewardingFund = "claimFromRewardingFund"
	// StakeCreate action type
	StakeCreate = "stakeCreate"
	// GrantReward action type
	GrantReward = "grantReward"
	// StakeUnstake action type
	StakeUnstake = "stakeUnstake"
	// StakeWithdraw action type
	StakeWithdraw = "stakeWithdraw"
	// StakeAddDeposit action type
	StakeAddDeposit = "stakeAddDeposit"
	// StakeRestake action type
	StakeRestake = "stakeRestake"
	// StakeChangeCandidate action type
	StakeChangeCandidate = "stakeChangeCandidate"
	// StakeTransferOwnership action type
	StakeTransferOwnership = "stakeTransferOwnership"
	// CandidateRegister action type
	CandidateRegister = "candidateRegister"
	// CandidateUpdate action type
	CandidateUpdate = "candidateUpdate"
	// PutPollResult action type
	PutPollResult = "putPollResult"
	statusSuccess = "success"
	statusFail    = "fail"
	actionTypeFee = "fee"
)

var IoTexCurrency = &types.Currency{
	Symbol:   "Iotx",
	Decimals: 18,
}

type (
	Genesis struct {
		Index int64  `json: "index"`
		Hash  string `json: "hash"`
	}
	NetworkIdentifier struct {
		Blockchain string `json: "blockchain"`
		Network    string `json: "network"`
	}
	Config struct {
		Genesis_block_identifier Genesis           `json: "genesis_block_identifier"`
		Network_identifier       NetworkIdentifier `json: "network_identifier"`
	}
	// IoTexClient is the IoTex blockchain client interface.
	IoTexClient interface {
		// GetChainID returns the network chain context, derived from the
		// genesis document.
		GetChainID(ctx context.Context) (string, error)

		// GetBlock returns the IoTex block at given height.
		GetBlock(ctx context.Context, height int64) (*IoTexBlock, error)

		// GetLatestBlock returns latest IoTex block.
		GetLatestBlock(ctx context.Context) (*IoTexBlock, error)

		// GetGenesisBlock returns the IoTex genesis block.
		GetGenesisBlock(ctx context.Context) (*IoTexBlock, error)

		// GetAccount returns the IoTex staking account for given owner address
		// at given height.
		GetAccount(ctx context.Context, height int64, owner string) (*Account, error)

		// SubmitTx submits the given JSON-encoded transaction to the node.
		SubmitTx(ctx context.Context, tx *iotextypes.Action) (txid string, err error)

		// GetStatus returns the status overview of the node.
		GetStatus(ctx context.Context) (*iotexapi.GetChainMetaResponse, error)
		// GetVersion returns the server's version.
		GetVersion(ctx context.Context) (*iotexapi.GetServerMetaResponse, error)
		// GetTransactions returns transactions of the block.
		GetTransactions(ctx context.Context, height int64) ([]*types.Transaction, error)
		// GetConfig returns the config.
		GetConfig() *Config
	}

	// IoTexBlock is the IoTex blockchain's block.
	IoTexBlock struct {
		Height       int64  // Block height.
		Hash         string // Block hash.
		Timestamp    int64  // UNIX time, converted to milliseconds.
		ParentHeight int64  // Height of parent block.
		ParentHash   string // Hash of parent block.
	}

	Account struct {
		Nonce   uint64
		Balance string
	}

	// grpcIoTexClient is an implementation of IoTexClient using gRPC.
	grpcIoTexClient struct {
		sync.RWMutex

		endpoint string
		grpcConn *grpc.ClientConn
		cfg      *Config
	}
)

// NewIoTexClient returns an implementation of IoTexClient
func NewIoTexClient(grpcAddr string, cfgPath string) (cli IoTexClient, err error) {
	file, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		return
	}
	cfg := &Config{}
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return
	}
	grpc, err := grpc.Dial(grpcAddr, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	if err != nil {
		return
	}
	cli = &grpcIoTexClient{endpoint: grpcAddr, grpcConn: grpc, cfg: cfg}
	return
}

func (c *grpcIoTexClient) GetChainID(ctx context.Context) (string, error) {
	return c.cfg.Network_identifier.Network, nil
}

func (c *grpcIoTexClient) GetBlock(ctx context.Context, height int64) (ret *IoTexBlock, err error) {
	err = c.reconnect()
	if err != nil {
		return
	}
	var parentHeight uint64
	if height <= 1 {
		parentHeight = 1
	} else {
		parentHeight = uint64(height) - 1
	}
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
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
	resp, err := client.GetBlockMetas(ctx, request)
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
	ret = &IoTexBlock{
		Height:       int64(blk.Height),
		Hash:         blk.Hash,
		Timestamp:    blk.Timestamp.Seconds * 1e3, // ms
		ParentHeight: int64(parentHeight),
		ParentHash:   parentBlk.Hash,
	}
	return
}

func (c *grpcIoTexClient) GetLatestBlock(ctx context.Context) (*IoTexBlock, error) {
	err := c.reconnect()
	if err != nil {
		return nil, err
	}
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	res, err := client.GetChainMeta(context.Background(), &iotexapi.GetChainMetaRequest{})
	if err != nil {
		return nil, err
	}
	return c.GetBlock(ctx, int64(res.ChainMeta.Height))
}

func (c *grpcIoTexClient) GetGenesisBlock(ctx context.Context) (*IoTexBlock, error) {
	return c.GetBlock(ctx, c.cfg.Genesis_block_identifier.Index)
}

func (c *grpcIoTexClient) GetAccount(ctx context.Context, height int64, owner string) (ret *Account, err error) {
	err = c.reconnect()
	if err != nil {
		return
	}
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	request := &iotexapi.GetAccountRequest{Address: owner}
	resp, err := client.GetAccount(ctx, request)
	if err != nil {
		return nil, err
	}
	ret = &Account{
		Nonce:   resp.AccountMeta.Nonce,
		Balance: resp.AccountMeta.Balance,
	}
	return
}

func (c *grpcIoTexClient) GetTransactions(ctx context.Context, height int64) (ret []*types.Transaction, err error) {
	blk, err := c.GetBlock(ctx, height)
	if err != nil {
		return
	}
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	// this limit from iotex-core's default value
	limit := uint64(1000)
	var actionInfo []*iotexapi.ActionInfo
	for i := uint64(0); ; i++ {
		request := &iotexapi.GetActionsRequest{
			Lookup: &iotexapi.GetActionsRequest_ByBlk{
				ByBlk: &iotexapi.GetActionsByBlockRequest{
					BlkHash: blk.Hash,
					Start:   i * limit,
					Count:   limit,
				},
			},
		}
		res, err := client.GetActions(context.Background(), request)
		if err != nil {
			break
		}
		actionInfo = append(actionInfo, res.ActionInfo...)
	}

	ret = make([]*types.Transaction, 0)
	for _, act := range actionInfo {
		decode, err := decodeAction(act, client)
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
	err = c.reconnect()
	if err != nil {
		return
	}
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	ret, err := client.SendAction(ctx, &iotexapi.SendActionRequest{Action: tx})
	if err != nil {
		return
	}
	txid = ret.ActionHash
	return
}

func (c *grpcIoTexClient) GetStatus(ctx context.Context) (*iotexapi.GetChainMetaResponse, error) {
	err := c.reconnect()
	if err != nil {
		return nil, err
	}
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	return client.GetChainMeta(context.Background(), &iotexapi.GetChainMetaRequest{})
}

func (c *grpcIoTexClient) GetVersion(ctx context.Context) (*iotexapi.GetServerMetaResponse, error) {
	err := c.reconnect()
	if err != nil {
		return nil, err
	}
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	return client.GetServerMeta(context.Background(), &iotexapi.GetServerMetaRequest{})
}

func (c *grpcIoTexClient) GetConfig() *Config {
	return c.cfg
}

func (c *grpcIoTexClient) reconnect() (err error) {
	c.Lock()
	defer c.Unlock()

	// Check if the existing connection is good.
	if c.grpcConn != nil && c.grpcConn.GetState() != connectivity.Shutdown {
		return
	}
	c.grpcConn, err = grpc.Dial(c.endpoint, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	return err
}

func decodeAction(act *iotexapi.ActionInfo, client iotexapi.APIServiceClient) (ret *types.Transaction, err error) {
	ret, status, err := gasFeeAndStatus(act, client)
	if err != nil {
		return
	}
	if act.GetAction().GetCore().GetExecution() != nil {
		// this one need special handler,TODO test when testnet enable systemlog
		err = handleExecution(ret, status, act.ActHash, client)
		return
	}
	amount, senderSign, actionType, dst, err := assertAction(act)
	if err != nil {
		return nil, err
	}
	if amount == "0" || actionType == "" {
		return nil, nil
	}
	var senderAmountWithSign, dstAmountWithSign string
	if senderSign == "-" {
		senderAmountWithSign = senderSign + amount
		dstAmountWithSign = amount
	} else {
		senderAmountWithSign = amount
		dstAmountWithSign = "-" + amount
	}
	src := []*addressAmount{{address: act.Sender, amount: senderAmountWithSign}}
	var dstAll []*addressAmount
	if dst != "" {
		dstAll = []*addressAmount{{address: dst, amount: dstAmountWithSign}}
	}
	err = packTransaction(ret, src, dstAll, actionType, status, act.ActHash)
	return
}

func handleExecution(ret *types.Transaction, status, hash string, client iotexapi.APIServiceClient) (err error) {
	request := &iotexapi.GetEvmTransfersByActionHashRequest{
		ActionHash: hash,
	}
	resp, err := client.GetEvmTransfersByActionHash(context.Background(), request)
	if err != nil {
		return
	}
	var src, dst addressAmountList
	for _, transfer := range resp.GetActionEvmTransfers().GetEvmTransfers() {
		src = append(src, &addressAmount{
			address: transfer.From,
			amount:  "-" + new(big.Int).SetBytes(transfer.Amount).String(),
		})
		dst = append(dst, &addressAmount{
			address: transfer.To,
			amount:  new(big.Int).SetBytes(transfer.Amount).String(),
		})
	}
	return packTransaction(ret, src, dst, Execution, status, hash)
}

func gasFeeAndStatus(act *iotexapi.ActionInfo, client iotexapi.APIServiceClient) (ret *types.Transaction, status string, err error) {
	ctx := context.Background()
	requestGetReceipt := &iotexapi.GetReceiptByActionRequest{ActionHash: act.GetActHash()}
	responseReceipt, err := client.GetReceiptByAction(ctx, requestGetReceipt)
	if err != nil {
		return
	}
	status = statusSuccess
	if responseReceipt.GetReceiptInfo().GetReceipt().GetStatus() != 1 {
		status = statusFail
	}

	gasConsumed := new(big.Int).SetUint64(responseReceipt.GetReceiptInfo().GetReceipt().GetGasConsumed())
	gasPrice, ok := new(big.Int).SetString(act.GetAction().GetCore().GetGasPrice(), 10)
	if !ok {
		err = errors.New("convert gas price error")
		return
	}
	gasFee := gasPrice.Mul(gasPrice, gasConsumed)
	// if gasFee is 0
	if gasFee.Sign() != 1 {
		return nil, "", nil
	}
	sender := addressAmountList{{address: act.Sender, amount: "-" + gasFee.String()}}
	var oper []*types.Operation
	_, oper, err = addOperation(sender, actionTypeFee, status, 0, oper)
	if err != nil {
		return nil, "", err
	}
	ret = &types.Transaction{
		TransactionIdentifier: &types.TransactionIdentifier{
			act.ActHash,
		},
		Operations: oper,
		Metadata:   nil,
	}
	return
}

func packTransaction(ret *types.Transaction, src, dst addressAmountList, actionType, status, hash string) (err error) {
	sort.Sort(src)
	sort.Sort(dst)
	var oper []*types.Operation
	endIndex, oper, err := addOperation(src, actionType, status, 1, oper)
	if err != nil {
		return
	}
	_, oper, err = addOperation(dst, actionType, status, endIndex, oper)
	if err != nil {
		return
	}
	ret.Operations = append(ret.Operations, oper...)
	return
}

func addOperation(l addressAmountList, actionType, status string, startIndex int64, oper []*types.Operation) (int64, []*types.Operation, error) {
	for _, s := range l {
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
				Value:    s.amount,
				Currency: IoTexCurrency,
				Metadata: nil,
			},
			Metadata: nil,
		})
		startIndex++
	}
	return startIndex, oper, nil
}

func assertAction(act *iotexapi.ActionInfo) (amount, senderSign, actionType, dst string, err error) {
	amount = "0"
	senderSign = "-"
	switch {
	case act.GetAction().GetCore().GetTransfer() != nil:
		actionType = Transfer
		amount = act.GetAction().GetCore().GetTransfer().GetAmount()
		dst = act.GetAction().GetCore().GetTransfer().GetRecipient()
	case act.GetAction().GetCore().GetDepositToRewardingFund() != nil:
		actionType = DepositToRewardingFund
		amount = act.GetAction().GetCore().GetDepositToRewardingFund().GetAmount()
	case act.GetAction().GetCore().GetClaimFromRewardingFund() != nil:
		actionType = ClaimFromRewardingFund
		amount = act.GetAction().GetCore().GetClaimFromRewardingFund().GetAmount()
		senderSign = "+"
	case act.GetAction().GetCore().GetStakeAddDeposit() != nil:
		actionType = StakeAddDeposit
		amount = act.GetAction().GetCore().GetClaimFromRewardingFund().GetAmount()
	case act.GetAction().GetCore().GetStakeCreate() != nil:
		actionType = StakeCreate
		amount = act.GetAction().GetCore().GetStakeCreate().GetStakedAmount()
	case act.GetAction().GetCore().GetStakeWithdraw() != nil:
		// TODO need to add amount when it's available on iotex-core
		actionType = StakeCreate
	case act.GetAction().GetCore().GetCandidateRegister() != nil:
		actionType = CandidateRegister
		amount = act.GetAction().GetCore().GetCandidateRegister().GetStakedAmount()
	}
	return
}
