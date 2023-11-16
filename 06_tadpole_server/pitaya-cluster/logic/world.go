package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"

	"pitaya_tadpole/logic/protocol"
	"pitaya_tadpole/protos"
)

// World contains all tadpoles
type World struct {
	component.Base
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
		group: groupName,
		app:   app,
	}
}

// Enter was called when new guest enter
func (w *World) Enter(ctx context.Context, msg []byte) (*protocol.EnterWorldResponse, error) {
	s := w.app.GetSessionFromCtx(ctx)
	uid := strconv.FormatInt(s.ID(), 10)
	s.Bind(ctx, uid)

	fmt.Println("World.Enter ------------> ", uid)

	w.app.GroupAddMember(ctx, w.group, uid)
	mbs, _ := w.app.GroupMembers(ctx, w.group)

	//TODO 仅用于示例，应该提供批量获取的接口，不应在循环中调用RPC
	// 获取其他member的坐标位置
	for _, v := range mbs {
		resp := protos.Response{}
		args := protos.Arg{Msg: v}
		if err := w.app.RPC(ctx, "store.store.getpos", &resp, &args); err != nil {
			fmt.Println("rpc err", err)
		}

		if resp.Code == 200 {
			posMsg := &protocol.UpdateMessage{}
			json.Unmarshal([]byte(resp.Msg), posMsg)

			// send to user
			w.app.SendPushToUsers("update", posMsg, []string{uid}, "game")
		}
	}

	s.OnClose(func() {
		fmt.Println("OnClose ------------> ", uid)
		w.app.GroupRemoveMember(ctx, w.group, uid)

		// 下线通知
		w.app.GroupBroadcast(ctx, "game", w.group, "leave", &protocol.LeaveWorldResponse{ID: s.ID()})

		// 移除存储的位置数据
		args := protos.Arg{Msg: uid}
		if err := w.app.RPC(ctx, "store.store.removepos", &protos.Response{}, &args); err != nil {
			fmt.Println("rpc err", err)
		}
	})

	log.Printf("session count: %d", len(mbs))
	return &protocol.EnterWorldResponse{ID: s.ID()}, nil
}

// Update refresh tadpole's position
func (w *World) Update(ctx context.Context, msg *protocol.UpdateMessage) {
	fmt.Println("World.Update ------------> ", w.app.GetSessionFromCtx(ctx).ID(), msg)
	// 状态变更通知
	if err := w.app.GroupBroadcast(ctx, "game", w.group, "update", msg); err != nil {
		fmt.Println("World.Update err:", err)
	}

	// 存储新的位置数据
	rawData, _ := json.Marshal(msg)
	args := protos.Arg{Msg: string(rawData)}
	if err := w.app.RPC(ctx, "store.store.updatepos", &protos.Response{}, &args); err != nil {
		fmt.Println("rpc err", err)
	}
}

// Message handler was used to communicate with each other
func (w *World) Message(ctx context.Context, msg *protocol.WorldMessage) {
	s := w.app.GetSessionFromCtx(ctx)
	msg.ID = s.ID()

	fmt.Println("World.Message ------------> ", msg)
	if err := w.app.GroupBroadcast(ctx, "game", w.group, "message", msg); err != nil {
		fmt.Println("World.Message err:", err)
	}
}
