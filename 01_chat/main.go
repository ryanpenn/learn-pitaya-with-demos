package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/acceptor"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/groups"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/timer"
)

type (
	// Room represents a component that contains a bundle of room related handler
	// like Join/Message
	Room struct {
		component.Base
		timer *timer.Timer
		app   pitaya.Pitaya
	}

	// UserMessage represents a message that user sent
	UserMessage struct {
		Name    string `json:"name"`
		Content string `json:"content"`
	}

	// NewUser message will be received when new user join room
	NewUser struct {
		Content string `json:"content"`
	}

	// AllMembers contains all members uid
	AllMembers struct {
		Members []string `json:"members"`
	}

	// JoinResponse represents the result of joining room
	JoinResponse struct {
		Code   int    `json:"code"`
		Result string `json:"result"`
	}
)

// NewRoom returns a Handler Base implementation
func NewRoom(app pitaya.Pitaya) *Room {
	return &Room{
		app: app,
	}
}

// AfterInit component lifetime callback
func (r *Room) AfterInit() {
	// 每分钟输出一条日志
	r.timer = pitaya.NewTimer(time.Minute, func() {
		count, err := r.app.GroupCountMembers(context.Background(), "room")
		logger.Log.Debugf("UserCount: Time=> %s, Count=> %d, Error=> %q", time.Now().String(), count, err)
	})
}

// 客户端通过 [request] room.join 调用该方法
// - ctx 请求的上下文,其中包含客户端的会话
// Join room
func (r *Room) Join(ctx context.Context, msg []byte) (*JoinResponse, error) {
	// 获得session信息
	s := r.app.GetSessionFromCtx(ctx)
	fakeUID := s.ID() // just use s.ID as uid !!!
	// 绑定用户ID
	err := s.Bind(ctx, strconv.Itoa(int(fakeUID))) // binding session uid

	if err != nil {
		return nil, pitaya.Error(err, "RH-000", map[string]string{"failed": "bind"})
	}

	uids, err := r.app.GroupMembers(ctx, "room")
	if err != nil {
		return nil, err
	}

	s.Push("onMembers", &AllMembers{Members: uids}) // 推送 onMembers 消息给当前连接的客户端

	// 广播消息 chat.room.onNewUser 给所有客户端 | 客户端响应方法为 onNewUser
	// notify others
	r.app.GroupBroadcast(ctx, "chat", "room", "onNewUser", &NewUser{Content: fmt.Sprintf("New user: %s", s.UID())})

	// new user join group
	r.app.GroupAddMember(ctx, "room", s.UID()) // add session to group

	// on session close, remove it from group
	s.OnClose(func() {
		r.app.GroupRemoveMember(ctx, "room", s.UID()) // remove session from group
	})

	return &JoinResponse{Result: "success"}, nil
}

// 客户端通过 [notify] room.message 调用该方法
// Message sync last message to all members
func (r *Room) Message(ctx context.Context, msg *UserMessage) {
	err := r.app.GroupBroadcast(ctx, "chat", "room", "onMessage", msg)
	if err != nil {
		fmt.Println("error broadcasting message", err)
	}
}

var app pitaya.Pitaya

func main() {
	// 首先声明了一个config的生成器
	conf := configApp()

	// @true 表示是面向客户端的前端服务器
	// "chat" 服务器类型
	// pitaya.Cluster 集群模式,也有也有单机模式(Standalone)
	// map 服务器元数据
	// conf 配置生成器的指针
	builder := pitaya.NewDefaultBuilder(true, "chat", pitaya.Cluster, map[string]string{}, *conf)

	// 添加一个接收器(必须是前端服务)
	builder.AddAcceptor(acceptor.NewWSAcceptor(":3250"))

	// 初始化一个内存组驱动
	builder.Groups = groups.NewMemoryGroupService(*config.NewDefaultMemoryGroupConfig()) // config 每 30s 自动检查是否应删除组

	// 构建
	app = builder.Build()
	defer app.Shutdown()

	// 创建一个组 "room"
	err := app.GroupCreate(context.Background(), "room")
	if err != nil {
		panic(err)
	}

	// 创建一个room组件,这个组件就是聊天室逻辑
	// rewrite component and handler name
	room := NewRoom(app)
	app.Register(room,
		component.WithName("room"),              // 组件名称
		component.WithNameFunc(strings.ToLower), // 组件名称方法(将组件的方法名转为小写，前端请求:starx.request("room.join",{},join))
	)

	log.SetFlags(log.LstdFlags | log.Llongfile)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("client"))))
	go http.ListenAndServe(":3251", nil)

	app.Start()
}

func configApp() *config.BuilderConfig {
	// refs: https://pitaya.readthedocs.io/en/latest/configuration.html
	// 默认生成器配置,也可以用 NewBuilderConfig() 创建
	conf := config.NewDefaultBuilderConfig()
	conf.Pitaya.Buffer.Handler.LocalProcess = 15                     // 本地消息处理器缓冲区的大小,默认值20
	conf.Pitaya.Heartbeat.Interval = time.Duration(15 * time.Second) // 心跳间隔,默认值30s
	conf.Pitaya.Buffer.Agent.Messages = 32                           // 每个代理缓冲区的大小,默认值100
	conf.Pitaya.Handler.Messages.Compression = false                 // 是否压缩消息,默认值true
	return conf
}
