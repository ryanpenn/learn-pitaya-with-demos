package logic

import (
	"context"
	"fmt"
	"log"
	"strconv"

	// "github.com/lonng/nano"
	// "github.com/lonng/nano/component"
	// "github.com/lonng/nano/examples/demo/tadpole/logic/protocol"
	// "github.com/lonng/nano/session"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"

	"pitaya_tadpole/logic/protocol"
)

// World contains all tadpoles
type World struct {
	component.Base
	// *nano.Group
	group string
	app   pitaya.Pitaya
}

// NewWorld returns a world instance
func NewWorld(app pitaya.Pitaya) *World {
	groupName := "group"

	err := app.GroupCreate(context.Background(), groupName)
	if err != nil {
		panic(err)
	}

	return &World{
		// Group: nano.NewGroup(uuid.New().String()),
		group: groupName,
		app:   app,
	}
}

// Init initialize world component
func (w *World) Init() {
	// session.Lifetime.OnClosed(func(s *session.Session) {
	// 	w.Leave(s)
	// 	w.Broadcast("leave", &protocol.LeaveWorldResponse{ID: s.ID()})
	// 	log.Println(fmt.Sprintf("session count: %d", w.Count()))
	// })
}

// Enter was called when new guest enter
/*
func (w *World) Enter(s *session.Session, msg []byte) error {
	w.Add(s)
	log.Println(fmt.Sprintf("session count: %d", w.Count()))
	return s.Response(&protocol.EnterWorldResponse{ID: s.ID()})
}
*/
func (w *World) Enter(ctx context.Context, msg []byte) (*protocol.EnterWorldResponse, error) {
	s := w.app.GetSessionFromCtx(ctx)
	uid := strconv.FormatInt(s.ID(), 10)
	s.Bind(ctx, uid)

	fmt.Println("World.Enter ------------> ", uid)

	w.app.GroupAddMember(ctx, w.group, uid)
	mbs, _ := w.app.GroupMembers(ctx, w.group)

	s.OnClose(func() {
		fmt.Println("OnClose ------------> ", uid)
		w.app.GroupRemoveMember(ctx, w.group, uid)

		// 下线通知
		w.app.GroupBroadcast(ctx, "game", w.group, "leave", &protocol.LeaveWorldResponse{ID: s.ID()})
	})

	log.Printf("session count: %d", len(mbs))
	return &protocol.EnterWorldResponse{ID: s.ID()}, nil
}

// Update refresh tadpole's position
/*
func (w *World) Update(s *session.Session, msg []byte) error {
	return w.Broadcast("update", msg)
}
*/
func (w *World) Update(ctx context.Context, msg *protocol.UpdateMessage) {
	fmt.Println("World.Update ------------> ", w.app.GetSessionFromCtx(ctx).ID(), msg)
	// 状态变更通知
	if err := w.app.GroupBroadcast(ctx, "game", w.group, "update", msg); err != nil {
		fmt.Println("World.Update err:", err)
	}
}

// Message handler was used to communicate with each other
/*
func (w *World) Message(s *session.Session, msg *protocol.WorldMessage) error {
	msg.ID = s.ID()
	return w.Broadcast("message", msg)
}
*/
func (w *World) Message(ctx context.Context, msg *protocol.WorldMessage) {
	s := w.app.GetSessionFromCtx(ctx)
	msg.ID = s.ID()

	fmt.Println("World.Message ------------> ", msg)
	if err := w.app.GroupBroadcast(ctx, "game", w.group, "message", msg); err != nil {
		fmt.Println("World.Message err:", err)
	}
}
