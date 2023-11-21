package gate

import (
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	pConfig "github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/serialize"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

func Start(c *config.GateConfig, serializer serialize.Serializer, conf *pConfig.Config) {
	meta := map[string]string{}
	builder := pitaya.NewBuilderWithConfigs(true, c.ServerType, pitaya.Cluster, meta, conf)

	builder.AddAcceptor(acceptor.NewWSAcceptor(fmt.Sprintf(":%d", c.Port)))

	builder.Serializer = serializer
	builder.Groups = groups.NewMemoryGroupService(*pConfig.NewDefaultMemoryGroupConfig())

	// 创建监听器 （TODO：可以不用了）
	//l := newGameListener()
	//builder.ServiceDiscovery.AddListener(l)

	// 构建
	app := builder.Build()
	defer app.Shutdown()

	// 添加game服路由规则
	addGameRouter(app, "game")
	// 注册SessionHandler
	sessionHandler(c, app, builder.SessionPool)

	app.Start()
}
