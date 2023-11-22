package chat

import (
	"github.com/topfreegames/pitaya/v2"
	pConfig "github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/serialize"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

func Start(c *config.ChatConfig, serializer serialize.Serializer, conf *pConfig.Config) {
	meta := map[string]string{}

	builder := pitaya.NewBuilderWithConfigs(false, c.ServerType, pitaya.Cluster, meta, conf)
	builder.Serializer = serializer
	builder.Groups = groups.NewMemoryGroupService(*pConfig.NewDefaultMemoryGroupConfig())

	app := builder.Build()
	defer app.Shutdown()

	Register(app, c)
	app.Start()
}
