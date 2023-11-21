package game

import (
	"github.com/topfreegames/pitaya/v2"
	pConfig "github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/serialize"
	"learn-pitaya-with-demos/cluster_game/game/internal"
	"strconv"

	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

func Start(c *config.GameConfig, serializer serialize.Serializer, conf *pConfig.Config) {
	meta := map[string]string{
		"game_server_id": strconv.Itoa(c.ServerID),
	}

	builder := pitaya.NewBuilderWithConfigs(false, c.ServerType, pitaya.Cluster, meta, conf)
	builder.Serializer = serializer
	builder.Groups = groups.NewMemoryGroupService(*pConfig.NewDefaultMemoryGroupConfig())

	pip := internal.Pipeline{}
	builder.HandlerHooks.BeforeHandler.PushBack(pip.BeforeRequest)
	builder.HandlerHooks.AfterHandler.PushBack(pip.AfterRequest)

	app := builder.Build()
	defer app.Shutdown()

	Register(app, c)
	pip.Register(app)

	app.Start()
}
