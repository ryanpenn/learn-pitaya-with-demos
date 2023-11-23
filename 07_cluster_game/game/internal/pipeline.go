package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/errors"
	"strconv"
)

type Pipeline struct {
	app pitaya.Pitaya
	mgr *GameManager
}

func (p *Pipeline) Register(app pitaya.Pitaya) {
	p.app = app
}

func (p *Pipeline) AfterInit() {
	if m, err := p.app.GetModule("game_manager"); err != nil {
		panic("pipeline error: game manager not found")
	} else {
		p.mgr = m.(*GameManager)
	}
}

func (p *Pipeline) BeforeRequest(ctx context.Context, in interface{}) (context.Context, interface{}, error) {
	if uid := p.app.GetSessionFromCtx(ctx).UID(); uid == "" {
		// 首次进入game服（还没有绑定UID），必须是请求指定的接口（PlayerInfoReq）
		if _, ok := in.(*PlayerInfoReq); ok {
			return ctx, in, nil // pass
		}

		return ctx, in, errors.NewError(fmt.Errorf("uid not found"), "404") // 非法的UID
	} else {
		if userID, err := strconv.ParseInt(uid, 10, 64); err != nil {
			return ctx, in, errors.NewError(fmt.Errorf("invalid uid"), "404") // 非法的UID
		} else {
			// 获取当前玩家角色，并写入context
			if player, ok := p.mgr.GetPlayer(userID); !ok {
				return ctx, in, errors.NewError(fmt.Errorf("uid not found"), "404") // 非法的UID
			} else {
				ctx = context.WithValue(ctx, "player", player)
				return ctx, in, nil // pass
			}
		}
	}
}

func (p *Pipeline) AfterRequest(ctx context.Context, out interface{}, err error) (interface{}, error) {
	if err != nil {
		fmt.Println("After Request: ", err)
		if _, ok := err.(*errors.Error); !ok {
			err = errors.NewError(err, "500")
		}
	}
	return out, err
}
