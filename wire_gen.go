// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/whj1990/go-core/store"
	"github.com/whj1990/go-core/trace"
	"github.com/whj1990/stats/internal/repo"
	"github.com/whj1990/stats/internal/service"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func initServer() (*grpc.Server, error) {
	loggerInterface := trace.NewGormLogger()
	db, err := store.NewDB(loggerInterface)
	if err != nil {
		return nil, err
	}
	orderRepo := repo.NewOrderRepo(db)
	tradeService := service.NewTradeService(orderRepo)
	handleServerServer := newAppImpl(tradeService)
	server := newServer(handleServerServer)
	return server, nil
}
