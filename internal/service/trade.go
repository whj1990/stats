package service

import (
	"context"
	"time"

	"github.com/whj1990/go-core/util"
	"github.com/whj1990/mine-grrpc/internal/repo"
	"github.com/whj1990/mine-grrpc/pbs"
)

type TradeService interface {
	TradeStats(context.Context, *pbs.TradeStatsReq) (*pbs.TradeStatsResp, error)
}

type tradeService struct {
	orderRepo repo.OrderRepo
}

func (s *tradeService) TradeStats(ctx context.Context, req *pbs.TradeStatsReq) (resp *pbs.TradeStatsResp, err error) {
	var sTime, eTime time.Time
	nowTime := time.Now()
	eTime = nowTime
	switch req.Type {
	case 1: //日
		sTime = util.GetZeroTime(nowTime)
	case 2: //周
		sTime = util.GetFirstDateOfWeek(nowTime)

	case 3: //月
		sTime = util.GetFirstDateOfMonth(nowTime)
	}
	//本周期交易数据指标
	resp.Data.Amount, resp.Data.Count, err = s.orderRepo.TradeStats(ctx, req.TenantId, sTime, eTime)
	if err != nil {
		return
	}
	//最近一小时交易数据指标
	resp.Data.AmountHour, resp.Data.CountHour, err = s.orderRepo.TradeStats(ctx, req.TenantId, nowTime.Add(-1*time.Hour), nowTime)
	if err != nil {
		return
	}
	return
}
