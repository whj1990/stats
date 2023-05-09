//go:build wireinject
// +build wireinject

package main

import (
	"google.golang.org/grpc"
	"mine/mine-grrpc/internal/repo"
	"mine/mine-grrpc/internal/service"

	"github.com/google/wire"
)

func initServer() (*grpc.Server, error) {
	panic(wire.Build(repo.ProviderSet, service.ProviderSet, newAppMineImpl, newServer))
}
