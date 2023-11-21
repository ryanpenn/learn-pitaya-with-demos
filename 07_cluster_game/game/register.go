package game

import (
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"learn-pitaya-with-demos/cluster_game/game/internal"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"learn-pitaya-with-demos/cluster_game/pkg/db"
	"strings"
	"time"
)

// Register 注册Web服务
func Register(app pitaya.Pitaya, c *config.GameConfig) {
	// 注册 db 模块
	if err := db.RegisterModule(app, c.DbName, c.DbAddr, time.Duration(c.DbTimeout)*time.Second); err != nil {
		panic(fmt.Errorf("db module register err: %v", err))
	}

	// 注册Manager
	if err := app.RegisterModule(internal.NewGameManager(app, c), "game_manager"); err != nil {
		panic(fmt.Errorf("game manager module register err: %v", err))
	}

	// handler
	app.Register(internal.NewGameHandler(app), component.WithName("handler"), component.WithNameFunc(strings.ToLower))
	// remote
	app.RegisterRemote(internal.NewGameRemote(app), component.WithName("remote"), component.WithNameFunc(strings.ToLower))
}
