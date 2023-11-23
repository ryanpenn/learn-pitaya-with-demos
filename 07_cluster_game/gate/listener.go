package gate

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/cluster"
	"learn-pitaya-with-demos/cluster_game/protos"
	"strconv"
)

type Listener struct {
	app pitaya.Pitaya
}

var _ cluster.SDListener = (*Listener)(nil)

func (l *Listener) BindApp(app pitaya.Pitaya) {
	l.app = app
}

func (l *Listener) AddServer(svr *cluster.Server) {
	fmt.Println("AddServer:", svr.ID, svr.Type, svr.Hostname, svr.Frontend)

	if svr.Type == "game" {
		// 服务器启动，创建聊天组
		if svrID, ok := svr.Metadata["game_server_id"]; ok {
			l.rpcToChat(svrID, "chat.remote.create")
		}
	}
}

func (l *Listener) RemoveServer(svr *cluster.Server) {
	fmt.Println("RemoveServer:", svr.ID, svr.Type, svr.Hostname, svr.Frontend)

	if svr.Type == "game" {
		if svrID, ok := svr.Metadata["game_server_id"]; ok {
			_ = svrID
			// 服务器关闭，移除聊天组
			//l.rpcToChat(svrID, "chat.remote.remove")
		}
	}
}

func (l *Listener) rpcToChat(svrID string, route string) {
	sid, _ := strconv.ParseInt(svrID, 10, 64)
	err := l.app.RPC(context.TODO(), route, &protos.RPCEmpty{}, &protos.RPCMsg{
		Code: sid,
	})
	if err != nil {
		fmt.Println("rpc err", err)
	}
}
