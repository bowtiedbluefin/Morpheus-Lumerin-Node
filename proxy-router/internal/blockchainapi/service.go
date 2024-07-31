package blockchainapi

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/contracts/sessionrouter"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi/structs"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/registries"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type BlockchainService struct {
	ethClient          *ethclient.Client
	providerRegistry   *registries.ProviderRegistry
	modelRegistry      *registries.ModelRegistry
	marketplace        *registries.Marketplace
	sessionRouter      *registries.SessionRouter
	morToken           *registries.MorToken
	explorerClient     *ExplorerClient
	sessionStorage     *storages.SessionStorage
	proxyService       *proxyapi.ProxyServiceSender
	diamonContractAddr common.Address

	legacyTx   bool
	privateKey interfaces.PrKeyProvider
	log        lib.ILogger
}

var (
	ErrPrKey         = errors.New("cannot get private key")
	ErrTxOpts        = errors.New("failed to get transactOpts")
	ErrNonce         = errors.New("failed to get nonce")
	ErrEstimateGas   = errors.New("failed to estimate gas")
	ErrSignTx        = errors.New("failed to sign transaction")
	ErrSendTx        = errors.New("failed to send transaction")
	ErrWaitMined     = errors.New("failed to wait for transaction to be mined")
	ErrSessionStore  = errors.New("failed to store session")
	ErrSessionReport = errors.New("failed to get session report from provider")

	ErrBid           = errors.New("failed to get bid")
	ErrProvider      = errors.New("failed to get provider")
	ErrTokenSupply   = errors.New("failed to parse token supply")
	ErrBudget        = errors.New("failed to parse token budget")
	ErrMyAddress     = errors.New("failed to get my address")
	ErrInitSession   = errors.New("failed to initiate session")
	ErrApprove       = errors.New("failed to approve")
	ErrMarshal       = errors.New("failed to marshal open session payload")
	ErrGetUserStake  = errors.New("failed to get user locked stake")
	ErrWithdrawStake = errors.New("failed to withdraw user locked stake")

	ErrNoBid = errors.New("no bids available")
	ErrModel = errors.New("can't get model")
)

func NewBlockchainService(
	ethClient *ethclient.Client,
	diamonContractAddr common.Address,
	morTokenAddr common.Address,
	explorerApiUrl string,
	privateKey interfaces.PrKeyProvider,
	sessionStorage *storages.SessionStorage,
	proxyService *proxyapi.ProxyServiceSender,
	log lib.ILogger,
	legacyTx bool,
) *BlockchainService {
	providerRegistry := registries.NewProviderRegistry(diamonContractAddr, ethClient, log)
	modelRegistry := registries.NewModelRegistry(diamonContractAddr, ethClient, log)
	marketplace := registries.NewMarketplace(diamonContractAddr, ethClient, log)
	sessionRouter := registries.NewSessionRouter(diamonContractAddr, ethClient, log)
	morToken := registries.NewMorToken(morTokenAddr, ethClient, log)

	explorerClient := NewExplorerClient(explorerApiUrl, morTokenAddr.String())
	return &BlockchainService{
		ethClient:          ethClient,
		providerRegistry:   providerRegistry,
		modelRegistry:      modelRegistry,
		marketplace:        marketplace,
		sessionRouter:      sessionRouter,
		legacyTx:           legacyTx,
		privateKey:         privateKey,
		morToken:           morToken,
		explorerClient:     explorerClient,
		proxyService:       proxyService,
		sessionStorage:     sessionStorage,
		diamonContractAddr: diamonContractAddr,
		log:                log,
	}
}

func (s *BlockchainService) GetLatestBlock(ctx context.Context) (uint64, error) {
	return s.ethClient.BlockNumber(ctx)
}

func (s *BlockchainService) GetAllProviders(ctx context.Context) ([]*structs.Provider, error) {
	addrs, providers, err := s.providerRegistry.GetAllProviders(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*structs.Provider, len(addrs))
	for i, value := range providers {
		result[i] = &structs.Provider{
			Address:   addrs[i],
			Endpoint:  value.Endpoint,
			Stake:     value.Stake,
			IsDeleted: value.IsDeleted,
			CreatedAt: value.CreatedAt,
		}
	}

	return result, nil
}

func (s *BlockchainService) GetAllModels(ctx context.Context) ([]*structs.Model, error) {
	ids, models, err := s.modelRegistry.GetAllModels(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]*structs.Model, len(ids))
	for i, value := range models {
		result[i] = &structs.Model{
			Id:        ids[i],
			IpfsCID:   value.IpfsCID,
			Fee:       value.Fee,
			Stake:     value.Stake,
			Owner:     value.Owner,
			Name:      value.Name,
			Tags:      value.Tags,
			CreatedAt: value.CreatedAt,
			IsDeleted: value.IsDeleted,
		}
	}

	return result, nil
}

func (s *BlockchainService) GetBidsByProvider(ctx context.Context, providerAddr common.Address, offset *big.Int, limit uint8) ([]*structs.Bid, error) {
	ids, bids, err := s.marketplace.GetBidsByProvider(ctx, providerAddr, offset, limit)
	if err != nil {
		return nil, err
	}

	return mapBids(ids, bids), nil
}

func (s *BlockchainService) GetBidsByModelAgent(ctx context.Context, modelId [32]byte, offset *big.Int, limit uint8) ([]*structs.Bid, error) {
	ids, bids, err := s.marketplace.GetBidsByModelAgent(ctx, modelId, offset, limit)
	if err != nil {
		return nil, err
	}

	return mapBids(ids, bids), nil
}

func (s *BlockchainService) GetActiveBidsByModel(ctx context.Context, modelId common.Hash) ([]*structs.Bid, error) {
	ids, bids, err := s.marketplace.GetActiveBidsByModel(ctx, modelId)
	if err != nil {
		return nil, err
	}

	return mapBids(ids, bids), nil
}

func (s *BlockchainService) GetActiveBidsByProvider(ctx context.Context, provider common.Address) ([]*structs.Bid, error) {
	ids, bids, err := s.marketplace.GetActiveBidsByProvider(ctx, provider)
	if err != nil {
		return nil, err
	}

	return mapBids(ids, bids), nil
}

func (s *BlockchainService) GetBidByID(ctx context.Context, ID common.Hash) (*structs.Bid, error) {
	bid, err := s.marketplace.GetBidById(ctx, ID)
	if err != nil {
		return nil, err
	}

	return mapBid(ID, *bid), nil
}

func (s *BlockchainService) GetRatedBids(ctx context.Context, modelID common.Hash) ([]ScoredBid, error) {
	modelStats, err := s.marketplace.GetModelStats(ctx, modelID)
	if err != nil {
		return nil, err
	}

	bidIDs, bids, providerModelStats, err := s.marketplace.GetAllBidsWithRating(ctx, modelID)
	if err != nil {
		return nil, err
	}

	ratedBids := rateBids(bidIDs, bids, providerModelStats, modelStats, s.log)

	return ratedBids, nil
}

func (s *BlockchainService) OpenSession(ctx context.Context, approval, approvalSig []byte, stake *big.Int) (common.Hash, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrPrKey, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrTxOpts, err)
	}

	sessionID, providerID, userID, err := s.sessionRouter.OpenSession(transactOpt, approval, approvalSig, stake, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrSendTx, err)
	}

	err = s.sessionStorage.AddSession(&storages.Session{
		Id:           sessionID.Hex(),
		UserAddr:     userID.Hex(),
		ProviderAddr: providerID.Hex(),
	})
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrSessionStore, err)
	}

	return sessionID, err
}

func (s *BlockchainService) CloseSession(ctx context.Context, sessionID common.Hash) (common.Hash, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrPrKey, err)
	}

	report, err := s.proxyService.GetSessionReport(ctx, sessionID)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrSessionReport, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrTxOpts, err)
	}

	tx, err := s.sessionRouter.CloseSession(transactOpt, sessionID, report.Message, report.SignedReport, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrSendTx, err)
	}

	return tx, nil
}

func (s *BlockchainService) GetSession(ctx *gin.Context, sessionID common.Hash) (*sessionrouter.Session, error) {
	return s.sessionRouter.GetSession(ctx, sessionID)
}

func (s *BlockchainService) GetProviderClaimableBalance(ctx *gin.Context, sessionID common.Hash) (*big.Int, error) {
	return s.sessionRouter.GetProviderClaimableBalance(ctx, sessionID)
}

func (s *BlockchainService) GetBalance(ctx *gin.Context) (*big.Int, *big.Int, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return nil, nil, lib.WrapError(ErrPrKey, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return nil, nil, lib.WrapError(ErrTxOpts, err)
	}

	ethBalance, err := s.ethClient.BalanceAt(ctx, transactOpt.From, nil)
	if err != nil {
		return nil, nil, err
	}

	morBalance, err := s.morToken.GetBalance(ctx, transactOpt.From)
	if err != nil {
		return nil, nil, err
	}

	return ethBalance, morBalance, nil
}

func (s *BlockchainService) SendETH(ctx *gin.Context, to common.Address, amount *big.Int) (common.Hash, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrPrKey, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrTxOpts, err)
	}

	nonce, err := s.ethClient.PendingNonceAt(ctx, transactOpt.From)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrNonce, err)
	}

	estimatedGas, err := s.ethClient.EstimateGas(ctx, ethereum.CallMsg{
		From:  transactOpt.From,
		To:    &to,
		Value: amount,
	})
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrEstimateGas, err)
	}

	//TODO: check if this is the right way to calculate gas
	gas := float64(estimatedGas) * 1.5
	tx := types.NewTransaction(nonce, to, amount, uint64(gas), transactOpt.GasPrice, nil)
	signedTx, err := s.signTx(ctx, tx, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrSignTx, err)
	}

	err = s.ethClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrSendTx, err)
	}

	// Wait for the transaction receipt
	_, err = bind.WaitMined(ctx, s.ethClient, signedTx)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrWaitMined, err)
	}

	return signedTx.Hash(), nil
}

func (s *BlockchainService) SendMOR(ctx context.Context, to common.Address, amount *big.Int) (common.Hash, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrPrKey, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrTxOpts, err)
	}

	tx, err := s.morToken.Transfer(transactOpt, to, amount)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrSendTx, err)
	}

	return tx.Hash(), nil
}

func (s *BlockchainService) GetAllowance(ctx context.Context, spender common.Address) (*big.Int, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return nil, lib.WrapError(ErrPrKey, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return nil, lib.WrapError(ErrTxOpts, err)
	}

	allowance, err := s.morToken.GetAllowance(ctx, transactOpt.From, spender)
	if err != nil {
		return nil, err
	}

	return allowance, nil
}

func (s *BlockchainService) Approve(ctx context.Context, spender common.Address, amount *big.Int) (common.Hash, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrPrKey, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrTxOpts, err)
	}

	tx, err := s.morToken.Approve(transactOpt, spender, amount)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrSendTx, err)
	}

	return tx.Hash(), nil
}

func (s *BlockchainService) GetTodaysBudget(ctx context.Context) (*big.Int, error) {
	return s.sessionRouter.GetTodaysBudget(ctx)
}

func (s *BlockchainService) ClaimProviderBalance(ctx context.Context, sessionID [32]byte, amount *big.Int) (common.Hash, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrPrKey, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrTxOpts, err)
	}

	txHash, err := s.sessionRouter.ClaimProviderBalance(transactOpt, sessionID, amount)
	if err != nil {
		return common.Hash{}, err
	}

	return txHash, nil
}

func (s *BlockchainService) GetTokenSupply(ctx context.Context) (*big.Int, error) {
	return s.morToken.GetTotalSupply(ctx)
}

func (s *BlockchainService) GetSessions(ctx *gin.Context, user, provider common.Address, offset *big.Int, limit uint8) ([]*structs.Session, error) {
	var (
		sessions []sessionrouter.Session
		err      error
	)
	if (user != common.Address{}) {
		sessions, err = s.sessionRouter.GetSessionsByUser(ctx, common.HexToAddress(ctx.Query("user")), offset, limit)
	} else {
		// hasProvider
		sessions, err = s.sessionRouter.GetSessionsByProvider(ctx, common.HexToAddress(ctx.Query("provider")), offset, limit)
	}
	if err != nil {
		return nil, err
	}
	return mapSessions(sessions), nil
}

func (s *BlockchainService) GetTransactions(ctx context.Context, page uint64, limit uint8) ([]structs.RawTransaction, error) {

	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return nil, lib.WrapError(ErrPrKey, err)
	}

	transactOpt, err := s.getTransactOpts(ctx, prKey)
	if err != nil {
		return nil, lib.WrapError(ErrTxOpts, err)
	}
	address := transactOpt.From

	ethTrxs, err := s.explorerClient.GetEthTransactions(address, page, limit)
	if err != nil {
		return nil, err
	}
	morTrxs, err := s.explorerClient.GetTokenTransactions(address, page, limit)
	if err != nil {
		return nil, err
	}

	allTrxs := append(ethTrxs, morTrxs...)
	sort.Slice(allTrxs, func(i, j int) bool {
		blockNumber1, err := strconv.ParseInt(allTrxs[i].BlockNumber, 10, 0)
		if err != nil {
			return false
		}
		blockNumber2, err := strconv.ParseInt(allTrxs[j].BlockNumber, 10, 0)
		if err != nil {
			return false
		}

		return blockNumber1 > blockNumber2
	})

	return allTrxs, nil
}

func (s *BlockchainService) openSessionByBid(ctx context.Context, bidID common.Hash, duration *big.Int) (common.Hash, error) {
	supply, err := s.GetTokenSupply(ctx)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrTokenSupply, err)
	}

	budget, err := s.GetTodaysBudget(ctx)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrBudget, err)
	}

	bid, err := s.marketplace.GetBidById(ctx, bidID)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrBid, err)
	}

	totalCost := duration.Mul(bid.PricePerSecond, duration)
	stake := totalCost.Div(totalCost.Mul(supply, totalCost), budget)

	userAddr, err := s.GetMyAddress(ctx)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrMyAddress, err)
	}

	provider, err := s.providerRegistry.GetProviderById(ctx, bid.Provider)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrProvider, err)
	}

	initRes, err := s.proxyService.InitiateSession(ctx, userAddr, bid.Provider, stake, bidID, provider.Endpoint)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrInitSession, err)
	}

	_, err = s.Approve(ctx, s.diamonContractAddr, stake)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrApprove, err)
	}

	return s.OpenSession(ctx, initRes.Approval, initRes.ApprovalSig, stake)
}

func (s *BlockchainService) OpenSessionByModelId(ctx context.Context, modelID common.Hash, duration *big.Int) (common.Hash, error) {
	supply, err := s.GetTokenSupply(ctx)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrTokenSupply, err)
	}

	budget, err := s.GetTodaysBudget(ctx)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrBudget, err)
	}

	modelStats, err := s.marketplace.GetModelStats(ctx, modelID)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrModel, err)
	}

	bidIDs, bids, providerStats, err := s.marketplace.GetAllBidsWithRating(ctx, modelID)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrBid, err)
	}

	if len(bids) == 0 {
		return common.Hash{}, ErrNoBid
	}

	userAddr, err := s.GetMyAddress(ctx)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrMyAddress, err)
	}

	scoredBids := rateBids(bidIDs, bids, providerStats, modelStats, s.log)
	for i, bid := range scoredBids {
		s.log.Infof("trying to open session with provider #%d %s", i, bid.Bid.Provider.String())
		hash, err := s.tryOpenSession(ctx, bid, duration, supply, budget, userAddr)
		if err == nil {
			return hash, nil
		}
		s.log.Errorf("failed to open session with provider %s: %s", bid.Bid.Provider.String(), err.Error())
	}

	return common.Hash{}, fmt.Errorf("no provider accepting session")
}

func (s *BlockchainService) tryOpenSession(ctx context.Context, bid ScoredBid, duration *big.Int, supply *big.Int, budget *big.Int, userAddr common.Address) (common.Hash, error) {
	provider, err := s.providerRegistry.GetProviderById(ctx, bid.Bid.Provider)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrProvider, err)
	}

	totalCost := duration.Mul(bid.Bid.PricePerSecond, duration)
	stake := totalCost.Div(totalCost.Mul(supply, totalCost), budget)

	initRes, err := s.proxyService.InitiateSession(ctx, userAddr, bid.Bid.Provider, stake, bid.Bid.Id, provider.Endpoint)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrInitSession, err)
	}

	_, err = s.Approve(ctx, s.diamonContractAddr, stake)
	if err != nil {
		return common.Hash{}, lib.WrapError(ErrApprove, err)
	}

	return s.OpenSession(ctx, initRes.Approval, initRes.ApprovalSig, stake)
}

func (s *BlockchainService) GetMyAddress(ctx context.Context) (common.Address, error) {
	prKey, err := s.privateKey.GetPrivateKey()
	if err != nil {
		return common.Address{}, lib.WrapError(ErrPrKey, err)
	}

	return lib.PrivKeyBytesToAddr(prKey)
}

func (s *BlockchainService) getTransactOpts(ctx context.Context, privKey lib.HexString) (*bind.TransactOpts, error) {
	privateKey, err := crypto.ToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	chainId, err := s.ethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	// TODO: deal with likely gasPrice issue so our transaction processes before another pending nonce.
	if s.legacyTx {
		gasPrice, err := s.ethClient.SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}
		transactOpts.GasPrice = gasPrice
	}

	transactOpts.Value = big.NewInt(0)
	transactOpts.Context = ctx

	return transactOpts, nil
}

func (s *BlockchainService) signTx(ctx context.Context, tx *types.Transaction, privKey lib.HexString) (*types.Transaction, error) {
	privateKey, err := crypto.ToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	chainId, err := s.ethClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	return types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
}

func (s *BlockchainService) GetWithdrawableUserStake(ctx context.Context, limit uint8) (*struct {
	Available *big.Int
	Hold      *big.Int
}, error) {
	userAddr, err := s.GetMyAddress(ctx)
	if err != nil {
		return nil, lib.WrapError(ErrMyAddress, err)
	}

	stake, err := s.sessionRouter.GetWithdrawableUserStake(ctx, userAddr, limit)
	if err != nil {
		return nil, lib.WrapError(ErrGetUserStake, err)
	}

	return stake, nil
}

func (s *BlockchainService) WithdrawUserStake(ctx context.Context, amount *big.Int, iterations uint8) error {
	err := s.sessionRouter.WithdrawUserStake(ctx, amount, iterations)
	if err != nil {
		return lib.WrapError(ErrWithdrawStake, err)
	}

	return nil
}
