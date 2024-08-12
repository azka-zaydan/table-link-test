package listener

import (
	"fmt"
	"net"

	"github.com/azka-zaydan/table-link-test/configs"
	"github.com/azka-zaydan/table-link-test/transport/grpc/router"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func Listen(configs *configs.Config) {
	lis, err := net.Listen(configs.Server.GRPC.Network, fmt.Sprintf("%s:%s", configs.Server.GRPC.Host, configs.Server.GRPC.Port))
	if err != nil {
		log.Fatal().Msg("Failed to listen")
	}

	gServer := router.Init([]grpc.ServerOption{})

	gServer.Serve(lis)
}
