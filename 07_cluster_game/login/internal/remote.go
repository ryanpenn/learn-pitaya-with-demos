package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/errors"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"learn-pitaya-with-demos/cluster_game/protos"
)

// Remote Web服RPC接口
type Remote struct {
	component.Base
	app pitaya.Pitaya
	cfg *config.LoginConfig
}

func NewRemote(app pitaya.Pitaya, c *config.LoginConfig) *Remote {
	return &Remote{
		app: app,
		cfg: c,
	}
}

// Auth 校验Token
func (r *Remote) Auth(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCEmpty, error) {
	if arg != nil {
		fmt.Println("arg", arg.Content)
		if arg.Content == "token" {
			return &protos.RPCEmpty{}, nil
		}
	}

	return nil, errors.NewError(fmt.Errorf("no auth"), "401")
}

// Renewal 续签Token
func (r *Remote) Renewal(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCMsg, error) {

	return nil, nil
}

// 获取角色信息
func (r *Remote) PlayerInfo(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCMsg, error) {

	return &protos.RPCMsg{Code: 200, Content: "ok"}, nil
}

// 更新角色信息
func (r *Remote) PlayerInfoUpdate(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCEmpty, error) {

	return nil, nil
}
