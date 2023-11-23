package chat

import (
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"learn-pitaya-with-demos/cluster_game/chat/internal"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"learn-pitaya-with-demos/cluster_game/pkg/db"
	"strings"
	"time"
)

// Register 注册服务
func Register(app pitaya.Pitaya, c *config.ChatConfig) {
	// 注册 db 模块
	if err := db.RegisterModule(app, c.DbName, c.DbAddr, time.Duration(c.DbTimeout)*time.Second); err != nil {
		panic(fmt.Errorf("db module register err: %v", err))
	}

	// module
	if err := app.RegisterModule(internal.NewChatManager(app, c), "manager"); err != nil {
		panic(fmt.Errorf("chat manager module register err: %v", err))
	}

	// remote
	app.RegisterRemote(internal.NewChatRemote(app), component.WithName("remote"), component.WithNameFunc(strings.ToLower))

	// handler
	app.Register(internal.NewChatHandler(app), component.WithName("handler"), component.WithNameFunc(strings.ToLower))
}
