package main

import (
	"context"

	"github.com/whj1990/mine-grrpc/pbs"

	"github.com/whj1990/mine-grrpc/internal/service"
)

// AppMineImpl implements the last service interface defined in the IDL.
type AppImpl struct {
	reviewService service.ReviewService
}

// ReviewProjectList implements the AppMineImpl interface.
func (s *AppImpl) ReviewProjectList(ctx context.Context, req *pbs.ReviewProjectListParams) (*pbs.ReviewProjectListResponse, error) {
	return s.reviewService.ReviewProjectList(ctx, req)
}
