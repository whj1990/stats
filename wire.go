//go:build wireinject
// +build wireinject

package main

import (
	"mine/mine-grrpc/internal/repo"
	"mine/mine-grrpc/internal/service"

	"google.golang.org/grpc"

	"github.com/google/wire"
)

func initServer() (*grpc.Server, error) {
	panic(wire.Build(repo.ProviderSet, service.ProviderSet, newAppImpl, newServer))
}
