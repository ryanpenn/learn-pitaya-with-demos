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

type GameHandler struct {
	component.Base
	app pitaya.Pitaya
}

func NewGameHandler(app pitaya.Pitaya) *GameHandler {
	return &GameHandler{
		app: app,
	}
}

func (h *GameHandler) PlayerInfo(ctx context.Context, in *PlayerInfoReq) (*PlayerInfoResp, error) {
	// TODO load data from db or cache by in.ID

	pl := &PlayerInfoResp{
		ID:    in.ID,
		Name:  fmt.Sprintf("player_%d", in.ID),
		Level: 1,
		Exp:   0,
	}

	err := h.app.GetSessionFromCtx(ctx).Bind(ctx, fmt.Sprintf("%d", in.ID))
	if err != nil {
		fmt.Println("Bind UID err", err)
		return nil, err
	}

	h.joinChatRoom(ctx, pl.ID)
	return pl, nil
}

// DoTask notify handler
func (h *GameHandler) DoTask(ctx context.Context, in *TaskReq) {
	fmt.Println("DoTask")

	if ids, err := h.app.SendPushToUsers("game.push.playerupdate", &PlayerUpdateInfo{
		Player: &PlayerInfoResp{
			ID:    1,
			Name:  "player",
			Level: 2,
			Exp:   100,
		},
		Info: "update",
	}, []string{h.app.GetSessionFromCtx(ctx).UID()}, "gate"); err != nil {
		fmt.Println("DoTask push err", err)
	} else {
		fmt.Println("DoTask push:", len(ids))
	}
}

func (h *GameHandler) joinChatRoom(ctx context.Context, uid int64) {
	// 加入聊天组
	sid := h.app.GetServer().Metadata["game_server_id"]
	groupID, _ := strconv.ParseInt(sid, 10, 64)

	join := &ChatJoin{
		UID:     uid,
		GroupID: groupID,
	}
	arg, _ := json.Marshal(join)
	err := h.app.RPC(ctx, "chat.remote.join", &protos.RPCEmpty{}, &protos.RPCMsg{
		Code:    0,
		Content: string(arg),
	})
	if err != nil {
		fmt.Println("chat.remote.join err", err)
	}
}
