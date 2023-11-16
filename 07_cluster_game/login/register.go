package login

import (
	"fmt"
	"strings"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"

	"learn-pitaya-with-demos/cluster_game/login/internal"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

// Register 注册Web服务
func Register(app pitaya.Pitaya, c *config.LoginConfig) {
	// 注册 id_maker 模块
	//id_maker.RegisterModule(app, &id_maker.MConfig{
	//	NodeIdForShort: 1,
	//	NodeIdForSnow:  1,
	//})

	// 注册 mongo 模块
	//mongo.RegisterModule(app, &mongo.MConfig{
	//	MasterAddr:    config.DefaultCfg.Mongo.MasterAddr, // 主库地址
	//	SlaveAddr:     config.DefaultCfg.Mongo.MasterAddr, // 从库地址
	//	MasterTimeout: config.DefaultCfg.Mongo.MasterTimeout * time.Second,
	//	SlaveTimeout:  config.DefaultCfg.Mongo.SlaveTimeout * time.Second,
	//})

	// 注册Web模块
	m := internal.NewWebModule(app, c)
	err := app.RegisterModule(m, "web")
	if err != nil {
		panic(fmt.Errorf("web module register err: %v", err))
	}

	// 注册 handler
	app.Register(internal.NewHandler(app, c), component.WithName("handler"), component.WithNameFunc(strings.ToLower))

	// 注册 remote
	app.RegisterRemote(internal.NewRemote(app, c), component.WithName("remote"), component.WithNameFunc(strings.ToLower))
}
