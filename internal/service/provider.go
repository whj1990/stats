package service

import (
	"github.com/whj1990/stats/internal/repo"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewTradeService)

func NewTradeService(orderRepo repo.OrderRepo) TradeService {
	return &tradeService{orderRepo}
}
