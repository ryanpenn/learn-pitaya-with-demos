package logic

import (
	"context"
	"fmt"
	"strconv"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"

	"pitaya_tadpole/logic/protocol"
	"pitaya_tadpole/protos"
)

// Manager component
type Manager struct {
	component.Base
	app pitaya.Pitaya
}

// NewManager returns  a new manager instance
func NewManager(app pitaya.Pitaya) *Manager {
	return &Manager{
		app: app,
	}
}

// Login handler was used to guest login
func (m *Manager) Login(ctx context.Context, msg *protocol.JoyLoginRequest) (*protocol.LoginResponse, error) {
	s := m.app.GetSessionFromCtx(ctx)
	fakeUID := s.ID()
	s.Bind(ctx, strconv.FormatInt(fakeUID, 10))

	fmt.Println("Manager.Login ------------> User Login", msg.Username, msg.Cipher)

	reply := protos.Response{}
	args := protos.Arg{Msg: msg.Username}
	if err := m.app.RPC(ctx, "store.store.save", &reply, &args); err != nil {
		fmt.Println("rpc err", err)
	}
	fmt.Println("store ------------> reply", reply.Code, reply.Msg)

	return &protocol.LoginResponse{
		Status: protocol.LoginStatusSucc,
		ID:     fakeUID,
	}, nil
}
