package main

import (
	"mine/mine-grrpc/internal/service"
	"mine/mine-grrpc/pbs"

	"github.com/whj1990/go-core/launch"
	"google.golang.org/grpc"
)

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	logger, closer := launch.InitPremise()
	defer logger.Sync()
	defer closer.Close()
	server, err := initServer()
	if err != nil {
		panic(err)
	}
	launch.RunGrpcServer(server)
	//launch.InitHttpServer()
}
func newAppImpl(reviewService service.ReviewService) pbs.HandleServerServer {
	return &AppImpl{reviewService}
}

func newServer(handler pbs.HandleServerServer) *grpc.Server {
	server := grpc.NewServer(launch.GrpcServerOptions()...)
	pbs.RegisterHandleServerServer(server, handler)
	return server
}