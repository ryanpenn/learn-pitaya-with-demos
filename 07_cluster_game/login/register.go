package login

import (
	"github.com/topfreegames/pitaya/v2"

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
	app.RegisterModule(m, "web")

	// 注册 handler
	//app.Register(handlers.NewHandler(&handlers.HConfig{
	//	App:            app,
	//	WebConfig:      c,
	//	AccountService: m.accService,
	//	TokenService:   m.tokenService,
	//}), component.WithName("handler"), component.WithNameFunc(strings.ToLower))

	// 注册 remote
	//app.RegisterRemote(handlers.NewRemote(&handlers.RConfig{
	//	App:               app,
	//	GamePlayerService: m.playerService,
	//}), component.WithName("remote"), component.WithNameFunc(strings.ToLower))

}
