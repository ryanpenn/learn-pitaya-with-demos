package internal

import (
	"context"
	"encoding/json"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/errors"
	"learn-pitaya-with-demos/cluster_game/protos"
)

type ChatRemote struct {
	component.Base
	app     pitaya.Pitaya
	manager *ChatManager
}

func NewChatRemote(app pitaya.Pitaya) *ChatRemote {
	return &ChatRemote{
		app: app,
	}
}

func (r *ChatRemote) AfterInit() {
	r.manager = GetManager(r.app)
}

func (r *ChatRemote) Join(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCEmpty, error) {
	var req ChatJoin
	err := json.Unmarshal([]byte(arg.Content), &req)
	if err != nil {
		return &protos.RPCEmpty{}, err
	}

	// 用户加入群组
	err = r.manager.joinToGroup(ctx, req.GroupID, req.UID)
	if err != nil {
		return &protos.RPCEmpty{}, errors.NewError(err, "500")
	}

	return &protos.RPCEmpty{}, nil
}

func (r *ChatRemote) Leave(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCEmpty, error) {
	var req ChatJoin
	err := json.Unmarshal([]byte(arg.Content), &req)
	if err != nil {
		return &protos.RPCEmpty{}, err
	}

	// 用户离开群组
	err = r.manager.leaveGroup(ctx, req.UID, req.GroupID)
	if err != nil {
		return &protos.RPCEmpty{}, err
	}

	return &protos.RPCEmpty{}, nil
}

func (r *ChatRemote) Create(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCEmpty, error) {
	err := r.manager.createGroup(ctx, arg.Code)
	if err != nil {
		return &protos.RPCEmpty{}, errors.NewError(err, "500")
	}
	return &protos.RPCEmpty{}, nil
}

func (r *ChatRemote) Remove(ctx context.Context, arg *protos.RPCMsg) (*protos.RPCEmpty, error) {
	err := r.manager.removeGroup(ctx, arg.Code)
	if err != nil {
		return &protos.RPCEmpty{}, errors.NewError(err, "500")
	}
	return &protos.RPCEmpty{}, nil
}
