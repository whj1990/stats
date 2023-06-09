package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/dc7ex/stats/pbs"

	"github.com/dc7ex/stats/internal/service"
)

type AppImpl struct {
	tradeService service.TradeService
}

func (s *AppImpl) StreamClientServer(srv pbs.HandleServer_StreamClientServerServer) error {
	for true {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			//发送结果，并关闭
			return srv.SendAndClose(&pbs.ParamResp{Date: time.Now().Format(time.DateTime)})
		}
		if err != nil {
			return err
		}
		log.Println(res.Num)

	}
	return nil
}

func (s *AppImpl) TradeStats(ctx context.Context, req *pbs.TradeStatsReq) (*pbs.TradeStatsResp, error) {
	return s.tradeService.TradeStats(ctx, req)
}
