package build

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/constants"
	"github.com/topfreegames/pitaya/v2/modules"
	"strconv"
)

// BuildWithGRPC 通过grpc实现服务器间的通信，可以作为nats的替代方案
func BuildWithGRPC(builder *pitaya.Builder, port int, meta map[string]string, c *config.Config) pitaya.Pitaya {
	meta[constants.GRPCHostKey] = "127.0.0.1"
	meta[constants.GRPCPortKey] = strconv.Itoa(port)

	grpcServerConfig := config.NewDefaultGRPCServerConfig()
	grpcServerConfig.Port = port

	if gs, err := cluster.NewGRPCServer(*grpcServerConfig, builder.Server, builder.MetricsReporters); err != nil {
		panic(err)
	} else {
		builder.RPCServer = gs
	}

	bs := modules.NewETCDBindingStorage(builder.Server, builder.SessionPool, *config.NewETCDBindingConfig(c))

	if gc, err := cluster.NewGRPCClient(
		*config.NewDefaultGRPCClientConfig(),
		builder.Server,
		builder.MetricsReporters,
		bs,
		cluster.NewInfoRetriever(*config.NewDefaultInfoRetrieverConfig()),
	); err != nil {
		panic(err)
	} else {
		builder.RPCClient = gc
	}

	app := builder.Build()
	if err := app.RegisterModule(bs, "bindingsStorage"); err != nil {
		panic(err)
	}

	return app
}
