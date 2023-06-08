package main

import (
	"github.com/whj1990/go-core/launch"
	"github.com/whj1990/mine-grrpc/internal/service"
	"github.com/whj1990/mine-grrpc/pbs"
	"google.golang.org/grpc"
)

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	logger, closer := launch.InitPremise(true)
	defer logger.Sync()
	defer closer.Close()
	server, err := initServer()
	if err != nil {
		panic(err)
	}
	launch.RunGrpcServer(server)
	//launch.InitHttpServer()
}
func newAppImpl(tradeService service.TradeService) pbs.HandleServerServer {
	return &AppImpl{tradeService}
}

func newServer(handler pbs.HandleServerServer) *grpc.Server {

	server := grpc.NewServer(launch.GrpcServerOptions()...)
	pbs.RegisterHandleServerServer(server, handler)
	return server
}
