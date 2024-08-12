package router

import (
	"github.com/azka-zaydan/table-link-test/configs"
	"github.com/azka-zaydan/table-link-test/shared/logger"
	"github.com/azka-zaydan/table-link-test/transport/grpc/listener"
	"google.golang.org/grpc"
)

var opts []grpc.ServerOption

func Start(configs *configs.Config) {

	gServer := grpc.NewServer(opts...)

	// put routes here
	lis, err := listener.Init(configs)
	if err != nil {
		logger.ErrorWithStack(err)

	}

	gServer.Serve(lis)
}
