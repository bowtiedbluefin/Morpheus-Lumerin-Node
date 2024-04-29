package apibus

import (
	"context"
	"math/big"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/proxyapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/rpcproxy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

type ApiBus struct {
	rpcProxy       *rpcproxy.RpcProxy
	aiEngine       *aiengine.AiEngine
	proxyRouterApi *proxyapi.ProxyRouterApi
}

func NewApiBus(rpcProxy *rpcproxy.RpcProxy, aiEngine *aiengine.AiEngine, proxyRouterApi *proxyapi.ProxyRouterApi) *ApiBus {
	return &ApiBus{
		rpcProxy:       rpcProxy,
		aiEngine:       aiEngine,
		proxyRouterApi: proxyRouterApi,
	}
}

// Proxy Router Api
func (apiBus *ApiBus) GetConfig(ctx context.Context) interface{} {
	return apiBus.proxyRouterApi.GetConfig(ctx)
}

func (apiBus *ApiBus) GetFiles(ctx *gin.Context) (int, interface{}) {
	return apiBus.proxyRouterApi.GetFiles(ctx)
}

func (apiBus *ApiBus) HealthCheck(ctx context.Context) interface{} {
	return apiBus.proxyRouterApi.HealthCheck(ctx)
}

func (apiBus *ApiBus) InitiateSession(ctx *gin.Context) (int, interface{}) {
	return apiBus.proxyRouterApi.InitiateSession(ctx)
}

// AiEngine
func (apiBus *ApiBus) Prompt(ctx context.Context) (string, error) {
	return apiBus.aiEngine.Prompt(ctx)
}

// RpcProxy
func (apiBus *ApiBus) GetLatestBlock(ctx context.Context) (uint64, error) {
	return apiBus.rpcProxy.GetLatestBlock(ctx)
}

func (apiBus *ApiBus) GetAllProviders(ctx context.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetAllProviders(ctx)
}

func (apiBus *ApiBus) GetAllModels(ctx context.Context) (int, gin.H) {
	return apiBus.rpcProxy.GetAllModels(ctx)
}

func (apiBus *ApiBus) GetBidsByProdiver(ctx context.Context, providerAddr string, offset *big.Int, limit uint8) (int, gin.H) {
	addr := common.HexToAddress(providerAddr)
	return apiBus.rpcProxy.GetBidsByProdiver(ctx, addr, offset, limit)
}

func (apiBus *ApiBus) GetBidsByModelAgent(ctx context.Context, modelAgentId [32]byte, offset *big.Int, limit uint8) (int, gin.H) {
	return apiBus.rpcProxy.GetBidsByModelAgent(ctx, modelAgentId, offset, limit)
}
