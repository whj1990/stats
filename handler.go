package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/whj1990/mine-grrpc/pbs"

	"github.com/whj1990/mine-grrpc/internal/service"
)

type AppImpl struct {
	reviewService service.ReviewService
}

func (s *AppImpl) ReviewProjectList(ctx context.Context, req *pbs.ReviewProjectListParams) (*pbs.ReviewProjectListResponse, error) {
	return s.reviewService.ReviewProjectList(ctx, req)
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
