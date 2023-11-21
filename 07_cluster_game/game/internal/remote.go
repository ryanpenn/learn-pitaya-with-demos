package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"learn-pitaya-with-demos/cluster_game/protos"
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

	return &protos.RPCEmpty{}, nil
}
