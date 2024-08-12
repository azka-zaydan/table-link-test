package listener

import (
	"fmt"
	"net"

	"github.com/azka-zaydan/table-link-test/configs"
	"github.com/rs/zerolog/log"
)

func Init(configs *configs.Config) (lis net.Listener, err error) {
	lis, err = net.Listen(configs.Server.GRPC.Network, fmt.Sprintf("%s:%s", configs.Server.GRPC.Host, configs.Server.GRPC.Port))
	if err != nil {
		log.Fatal().Msg("Failed to listen")
		return
	}

	return
}
