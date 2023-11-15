package logic

import (
	"context"
	"fmt"
	"strconv"

	// "github.com/lonng/nano/component"
	// "github.com/lonng/nano/examples/demo/tadpole/logic/protocol"
	// "github.com/lonng/nano/session"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"

	"pitaya_tadpole/logic/protocol"
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

	return &protocol.LoginResponse{
		Status: protocol.LoginStatusSucc,
		ID:     fakeUID,
	}, nil
}
