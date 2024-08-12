package router

import (
	"google.golang.org/grpc"
)

func Init(opts []grpc.ServerOption) *grpc.Server {

	gServer := grpc.NewServer(opts...)

	// put routes here

	return gServer
}
