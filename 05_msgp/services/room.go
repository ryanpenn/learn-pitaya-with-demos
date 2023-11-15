package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"

	"learn-pitaya-with-demos/msgp/msg"
)

type RoomComponent struct {
	component.Base
	app pitaya.Pitaya
}

func NewRoomComponent(app pitaya.Pitaya) *RoomComponent {
	return &RoomComponent{
		app: app,
	}
}

// Join 加入房间
func (c *RoomComponent) Join(ctx context.Context, m *msg.NewUser) (*msg.Response, error) {
	s := c.app.GetSessionFromCtx(ctx)
	fakeUID := s.ID()
	s.Bind(ctx, strconv.FormatInt(fakeUID, 10))

	s.Push("message", &msg.Message{
		Content: fmt.Sprintf("Welcome %s-%s", m.Name, s.UID()),
	})

	resp := &msg.Response{
		Code: 0,
		Msg:  "OK",
	}

	return resp, nil
}
