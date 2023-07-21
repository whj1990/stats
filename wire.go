//go:build wireinject
// +build wireinject

package main

import (
	"github.com/whj1990/stats/internal/repo"
	"github.com/whj1990/stats/internal/service"

	"google.golang.org/grpc"

	"github.com/google/wire"
)

func initServer() (*grpc.Server, error) {
	panic(wire.Build(repo.ProviderSet, service.ProviderSet, newAppImpl, newServer))
}
