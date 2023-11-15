package main

import (
	"flag"
	"strings"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"

	"learn-pitaya-with-demos/msgp/client"
	"learn-pitaya-with-demos/msgp/serializer"
	"learn-pitaya-with-demos/msgp/services"
)

func main() {
	// 客户端
	cli := flag.Bool("cli", false, "-cli true")
	flag.Parse()

	if *cli {
		client.Run()
		return
	}

	// 配置
	conf := config.NewDefaultBuilderConfig()

	// 构造器
	// 服务器类型: chat
	builder := pitaya.NewDefaultBuilder(true, "chat", pitaya.Cluster, map[string]string{}, *conf)
	builder.AddAcceptor(acceptor.NewTCPAcceptor(":9001"))

	builder.Serializer = serializer.NewMsgPackSerializer() // 使用 msgpack 序列化

	// 构造
	app := builder.Build()
	defer app.Shutdown()

	// 注册组件
	tc := services.NewRoomComponent(app)
	app.Register(tc,
		component.WithName("room"), // 组件名
		component.WithNameFunc(strings.ToLower),
	)

	// 启动服务器
	app.Start()
}
