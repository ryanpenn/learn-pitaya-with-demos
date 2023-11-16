package internal

import (
	"context"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"learn-pitaya-with-demos/cluster_game/pkg/msg"
)

type Handler struct {
	component.Base
	app pitaya.Pitaya
	cfg *config.LoginConfig
}

// Auth 校验Token
func (h *Handler) Auth(ctx context.Context, req *msg.ReqAuth) (*msg.Empty, error) {

	return nil, nil
}
