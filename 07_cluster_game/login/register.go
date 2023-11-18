package login

import (
	"fmt"
	"learn-pitaya-with-demos/cluster_game/pkg/db"
	"strings"
	"time"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"

	"learn-pitaya-with-demos/cluster_game/login/internal"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

// Register 注册Web服务
func Register(app pitaya.Pitaya, c *config.LoginConfig) {

	// 注册 db 模块
	if err := db.RegisterModule(app, c.DbName, c.DbAddr, time.Duration(c.DbTimeout)*time.Second); err != nil {
		panic(fmt.Errorf("db module register err: %v", err))
	}

	// 注册Web模块
	m := internal.NewWebModule(app, c)
	if err := app.RegisterModule(m, "web"); err != nil {
		panic(fmt.Errorf("web module register err: %v", err))
	}

	// 注册 remote
	app.RegisterRemote(internal.NewRemote(app, c), component.WithName("remote"), component.WithNameFunc(strings.ToLower))
}
