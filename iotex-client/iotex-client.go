package iotex_client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"sync"
	"time"

	"github.com/coinbase/rosetta-sdk-go/types"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials"

	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
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
	// GrantReward action type
	GrantReward = "grantReward"
	// StakeCreate action type
	StakeCreate = "stakeCreate"
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
)

var IoTexCurrency = &types.Currency{
	Symbol:   "IoTex",
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

		// GetBlock returns the Oasis block at given height.
		GetBlock(ctx context.Context, height int64) (*IoTexBlock, error)

		// GetLatestBlock returns latest Oasis block.
		GetLatestBlock(ctx context.Context) (*IoTexBlock, error)

		// GetGenesisBlock returns the Oasis genesis block.
		GetGenesisBlock(ctx context.Context) (*IoTexBlock, error)

		// GetAccount returns the Oasis staking account for given owner address
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
	grpc, err := newDefaultGRPCConn(grpcAddr)
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
	fmt.Println(parentHeight, height)
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
			continue
		}
		ret = append(ret, decode)
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
	} else {
		// Connection needs to be re-established.
		c.grpcConn = nil
	}
	c.grpcConn, err = iotex.NewDefaultGRPCConn(c.endpoint)
	return err
}

// newDefaultGRPCConn creates a default grpc connection. With tls and retry.
func newDefaultGRPCConn(endpoint string) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(10 * time.Second)),
		grpc_retry.WithMax(3),
	}
	return grpc.Dial(endpoint,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
}

func decodeAction(act *iotexapi.ActionInfo, client iotexapi.APIServiceClient) (ret *types.Transaction, err error) {
	ctx := context.Background()
	requestGetReceipt := &iotexapi.GetReceiptByActionRequest{ActionHash: act.GetActHash()}
	responseReceipt, err := client.GetReceiptByAction(ctx, requestGetReceipt)
	if err != nil {
		return
	}
	status := "succeed"
	if responseReceipt.GetReceiptInfo().GetReceipt().GetStatus() != 1 {
		status = "fail"
	}
	gasFee, ok := new(big.Int).SetString(act.GetGasFee(), 10)
	if !ok {
		return nil, errors.New("convert gasFee error")
	}
	var actionType, dst string
	amount := "0"
	senderSign := "-"
	switch {
	case act.GetAction().GetCore().GetTransfer() != nil:
		actionType = Transfer
		amount = act.GetAction().GetCore().GetTransfer().GetAmount()
		dst = act.GetAction().GetCore().GetTransfer().GetRecipient()
	case act.GetAction().GetCore().GetExecution() != nil:
		actionType = Execution
		amount = act.GetAction().GetCore().GetExecution().GetAmount()
		dst = act.GetAction().GetCore().GetExecution().GetContract()
	case act.GetAction().GetCore().GetDepositToRewardingFund() != nil:
		actionType = DepositToRewardingFund
		amount = act.GetAction().GetCore().GetDepositToRewardingFund().GetAmount()
		//dst=act.GetAction().GetCore().GetDepositToRewardingFund().get
	case act.GetAction().GetCore().GetClaimFromRewardingFund() != nil:
		actionType = ClaimFromRewardingFund
		amount = act.GetAction().GetCore().GetClaimFromRewardingFund().GetAmount()
		senderSign = "+"
	case act.GetAction().GetCore().GetStakeAddDeposit() != nil:
		// TODO need to check this
		actionType = StakeAddDeposit
		amount = act.GetAction().GetCore().GetClaimFromRewardingFund().GetAmount()
	//des = act.GetAction().GetCore().GetClaimFromRewardingFund()
	case act.GetAction().GetCore().GetStakeCreate() != nil:
		// TODO need to check this
		actionType = StakeCreate
		amount = act.GetAction().GetCore().GetStakeCreate().GetStakedAmount()
		//des = act.GetAction().GetCore().GetClaimFromRewardingFund()
		//case act.GetAction().GetCore().GetStakeWithdraw() != nil:
		//	// TODO need to check this
		//	actionType = StakeAddDeposit
		//	amount = act.GetAction().GetCore().GetStakeWithdraw()()
		//des = act.GetAction().GetCore().GetClaimFromRewardingFund()
		//senderSign = "+"
	}

	if amount == "0" {
		return nil, nil
	}
	amountInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return nil, errors.New("convert amount error")
	}
	amountInt = amountInt.Add(amountInt, gasFee)
	var senderAmountWithSign, dstAmountWithSign string
	if senderSign == "-" {
		senderAmountWithSign = senderSign + amountInt.String()
		dstAmountWithSign = amount
	} else {
		senderAmountWithSign = amountInt.String()
		dstAmountWithSign = "-" + amount
	}
	oper := []*types.Operation{
		&types.Operation{
			OperationIdentifier: &types.OperationIdentifier{
				Index:        0,
				NetworkIndex: nil,
			},
			RelatedOperations: nil,
			Type:              actionType,
			Status:            status,
			Account: &types.AccountIdentifier{
				Address:    act.Sender,
				SubAccount: nil,
				Metadata:   nil,
			},
			Amount: &types.Amount{
				Value:    senderAmountWithSign,
				Currency: IoTexCurrency,
				Metadata: nil,
			},
			Metadata: nil,
		},
		&types.Operation{
			OperationIdentifier: &types.OperationIdentifier{
				Index:        1,
				NetworkIndex: nil,
			},
			RelatedOperations: nil,
			Type:              actionType,
			Status:            status,
			Account: &types.AccountIdentifier{
				Address:    dst,
				SubAccount: nil,
				Metadata:   nil,
			},
			Amount: &types.Amount{
				Value:    dstAmountWithSign,
				Currency: IoTexCurrency,
				Metadata: nil,
			},
			Metadata: nil,
		},
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
