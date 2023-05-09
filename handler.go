package main

import (
	"context"
	"mine/mine-grrpc/pbs"

	"mine/mine-grrpc/internal/service"
)

// AppMineImpl implements the last service interface defined in the IDL.
type AppMineImpl struct {
	reviewService service.ReviewService
}

// ReviewProjectList implements the AppMineImpl interface.
func (s *AppMineImpl) ReviewProjectList(ctx context.Context, req *pbs.ReviewProjectListParams) (*pbs.ReviewProjectListResponse, error) {
	return s.reviewService.ReviewProjectList(ctx, req)
}
