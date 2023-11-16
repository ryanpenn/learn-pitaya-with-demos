package internal

import (
	"context"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"

	"learn-pitaya-with-demos/cluster_game/protos"
)

// Remote Web服RPC接口
type Remote struct {
	component.Base
	app pitaya.Pitaya
}

// 获取角色信息
func (r *Remote) PlayerInfo(ctx context.Context, arg *protos.Arg) (*protos.Response, error) {

	return &protos.Response{Code: 200, Msg: "ok"}, nil
}

// 更新角色信息
func (r *Remote) PlayerInfoUpdate(ctx context.Context, arg *protos.Arg) (*protos.Response, error) {

	return nil, nil
}
