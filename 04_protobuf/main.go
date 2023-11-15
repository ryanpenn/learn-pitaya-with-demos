package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/client"
	"github.com/topfreegames/pitaya/v2/cluster"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/route"
	"github.com/topfreegames/pitaya/v2/serialize/protobuf"

	"learn-pitaya-with-demos/protobuf/protos"
	"learn-pitaya-with-demos/protobuf/services"
)

func configureBackend(app pitaya.Pitaya) {
	room := services.NewRoom(app)
	app.Register(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)

	app.RegisterRemote(room,
		component.WithName("room"),
		component.WithNameFunc(strings.ToLower),
	)
}

func configureFrontend(app pitaya.Pitaya, port int) {
	app.Register(services.NewConnector(app),
		component.WithName("connector"),
		component.WithNameFunc(strings.ToLower),
	)

	app.RegisterRemote(services.NewConnectorRemote(app),
		component.WithName("connectorremote"),
		component.WithNameFunc(strings.ToLower),
	)

	err := app.AddRoute("room", func(
		ctx context.Context,
		route *route.Route,
		payload []byte,
		servers map[string]*cluster.Server,
	) (*cluster.Server, error) {
		for k := range servers {
			return servers[k], nil
		}
		return nil, nil
	})

	if err != nil {
		fmt.Printf("error adding route %s\n", err.Error())
	}

	// Dictionary 目的是为了压缩路由调用
	err = app.SetDictionary(map[string]uint16{
		"connector.getsessiondata": 1,
		"connector.setsessiondata": 2,
		"room.room.getsessiondata": 3,
		"onMessage":                4,
		"onMembers":                5,
	})

	if err != nil {
		fmt.Printf("error setting route dictionary %s\n", err.Error())
	}
}

func main() {
	cli := flag.Bool("cli", false, "run client")
	port := flag.Int("port", 3250, "the port to listen")
	svType := flag.String("type", "connector", "the server type")
	isFrontend := flag.Bool("frontend", true, "if server is frontend")

	flag.Parse()

	if *cli {
		// 运行测试客户端
		RunClient()
		return
	}

	builder := pitaya.NewDefaultBuilder(*isFrontend, *svType, pitaya.Cluster, map[string]string{}, *config.NewDefaultBuilderConfig())
	if *isFrontend {
		// 运行前端服务器
		tcp := acceptor.NewTCPAcceptor(fmt.Sprintf(":%d", *port))
		builder.AddAcceptor(tcp)
	}

	builder.Groups = groups.NewMemoryGroupService(*config.NewDefaultMemoryGroupConfig())
	// 配置protobuf序列化
	builder.Serializer = protobuf.NewSerializer()

	app := builder.Build()
	defer app.Shutdown()

	if !*isFrontend {
		configureBackend(app)
	} else {
		configureFrontend(app, *port)
	}

	app.Start()
}

func RunClient() {
	time.Sleep(time.Second)

	// 通过pitaya提供的client测试连接
	c := client.New(logrus.InfoLevel, 100*time.Second)
	err := c.ConnectTo("127.0.0.1:3250")
	if err != nil {
		fmt.Println("conn server err:", err)
		return
	}

	go func(c *client.Client) {
		for {
			select {
			case data := <-c.MsgChannel():
				// 处理服务器返回的数据
				if data.Err {
					fmt.Println("data err:", string(data.Data))
					break
				}
				// fmt.Println("data ----->", string(data.Data))
				var res protos.Response
				protobuf.NewSerializer().Unmarshal(data.Data, res)
				fmt.Printf("res: %#v\n", res)
			}
		}
	}(c)

	// 发送请求
	// 服务器类型.组件名.方法名
	if _, err := c.SendRequest("room.room.entry", []byte("")); err != nil {
		fmt.Println("send request err:", err)
		return
	}
	time.Sleep(time.Second)

	if _, err := c.SendRequest("room.room.join", []byte("")); err != nil {
		fmt.Println("send request err:", err)
		return
	}
	time.Sleep(time.Second)

	select {}
}
