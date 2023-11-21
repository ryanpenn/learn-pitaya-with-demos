package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
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

	err := h.app.GetSessionFromCtx(ctx).Bind(ctx, "1")
	if err != nil {
		fmt.Println("Bind uid err", err)
		return nil, err
	}

	return &PlayerInfoResp{
		ID:    1,
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
