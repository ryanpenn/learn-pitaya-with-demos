package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"

	"learn-pitaya-with-demos/cluster_chat/services"
)

type srvConfig struct {
	port     int
	frontend bool
}

var scs map[string]srvConfig = map[string]srvConfig{
	"game": {
		port:     3250,
		frontend: true,
	},
	"worker": {
		port:     3251,
		frontend: false,
	},
	"log": {
		port:     3252,
		frontend: false,
	},
}

func main() {

	srvType := flag.String("type", "game", "server type")
	flag.Parse()

	srvConfig, ok := scs[*srvType]
	if !ok {
		fmt.Println("not found type:", *srvType)
		return
	}

	// 启动worker
	// 通过redis实现可靠RPC(ReliableRPC),发生任何错误，pitaya会进行重试
	conf := viper.New()
	conf.SetDefault("pitaya.worker.redis.url", "localhost:6379")
	conf.SetDefault("pitaya.worker.redis.pool", "3")
	config := config.NewConfig(conf)

	// 构造器
	builder := pitaya.NewBuilderWithConfigs(srvConfig.frontend, *srvType, pitaya.Cluster, map[string]string{}, config)

	// 如果是前端服务器
	if srvConfig.frontend {
		builder.AddAcceptor(acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", srvConfig.port)))
	}

	// 构建
	app := builder.Build()
	switch *srvType {
	case "game":
		app.Register(services.NewRoom(app), component.WithName("room"), component.WithNameFunc(strings.ToLower))
		app.Register(services.NewAccount(app), component.WithName("account"), component.WithNameFunc(strings.ToLower))
	case "log":
		app.RegisterRemote(&services.Log{}, component.WithName("log"), component.WithNameFunc(strings.ToLower))
	case "worker":
		worker := services.Worker{}
		worker.Configure(app)
	}

	// 延迟关闭
	defer app.Shutdown()
	// 启动
	app.Start()

}
