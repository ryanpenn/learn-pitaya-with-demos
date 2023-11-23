package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"learn-pitaya-with-demos/cluster_game/protos"
	"strconv"
)

type GameRemote struct {
	component.Base
	app pitaya.Pitaya
}

func NewGameRemote(app pitaya.Pitaya) *GameRemote {
	return &GameRemote{
		app: app,
	}
}

func (r *GameRemote) Offline(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCEmpty, error) {
	// 处理用户掉线
	fmt.Println("Offline", arg.Content)

	r.leaveChatRoom(ctx)

	return &protos.RPCEmpty{}, nil
}

func (r *GameRemote) leaveChatRoom(ctx context.Context) {
	uid := r.app.GetSessionFromCtx(ctx).UID()
	id, _ := strconv.ParseInt(uid, 10, 64)

	sid := r.app.GetServer().Metadata["game_server_id"]
	groupID, _ := strconv.ParseInt(sid, 10, 64)

	join := &ChatJoin{
		UID:     id,
		GroupID: groupID,
	}
	data, _ := json.Marshal(join)
	err := r.app.RPC(ctx, "chat.remote.leave", &protos.RPCEmpty{}, &protos.RPCMsg{
		Code:    0,
		Content: string(data),
	})
	if err != nil {
		fmt.Println("chat.remote.join leave", err)
	}
}
