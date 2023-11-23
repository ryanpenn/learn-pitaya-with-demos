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
	fmt.Println("PlayerInfo", in)

	err := h.app.GetSessionFromCtx(ctx).Bind(ctx, fmt.Sprintf("%s", in.ID))
	if err != nil {
		fmt.Println("Bind uid err", err)
		return nil, err
	}

	// 加入聊天组
	sid := h.app.GetServer().Metadata["game_server_id"]
	groupID, _ := strconv.ParseInt(sid, 10, 64)
	join := &ChatJoin{
		UID:     in.ID,
		GroupID: groupID,
	}
	arg, _ := json.Marshal(join)
	err = h.app.RPC(ctx, "chat.remote.join", &protos.RPCEmpty{}, &protos.RPCMsg{
		Code:    0,
		Content: string(arg),
	})
	if err != nil {
		fmt.Println("chat.remote.join err", err)
	}

	return &PlayerInfoResp{
		ID:    in.ID,
		Name:  "player",
		Level: 1,
		Exp:   0,
	}, nil
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
