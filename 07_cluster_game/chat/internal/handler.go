package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"time"
)

type ChatHandler struct {
	component.Base
	app pitaya.Pitaya
}

func NewChatHandler(app pitaya.Pitaya) *ChatHandler {
	return &ChatHandler{
		app: app,
	}
}

func createGroups(app pitaya.Pitaya) {
	svrs, err := app.GetServersByType("game")
	if err != nil {
		fmt.Println("AfterInit GetServersByType err", err)
		return
	}

	// 每个game服一个room
	for _, s := range svrs {
		roomName := fmt.Sprintf("room_%s", s.Metadata["game_server_id"])
		err := app.GroupCreate(context.Background(), roomName)
		if err != nil {
			fmt.Println("AfterInit GroupCreate err", err)
		} else {
			fmt.Println("GroupCreate...", roomName)
		}
	}

	// 跨服聊天room
	err = app.GroupCreate(context.Background(), groupName(0))
	if err != nil {
		fmt.Println("AfterInit GroupCreate 0 err", err)
	}
}

func (h *ChatHandler) AfterInit() {
	// TODO (待改进)直接调用获取不到game服，延迟3秒再执行
	time.AfterFunc(time.Second*3, func() {
		createGroups(h.app)
	})
}

func (h *ChatHandler) Send(ctx context.Context, msg *ChatMessage) {
	if msg != nil {
		s := h.app.GetSessionFromCtx(ctx)
		uid := s.UID()

		switch msg.ChatType {
		case 0:
			_, err := h.app.SendPushToUsers("chat.push.message", msg, []string{fmt.Sprintf("%d", msg.To)}, "gate")
			if err != nil {
				fmt.Println("SendPushToUsers err", err)
			}
		case 1:
			pushToRoom(h.app, ctx, msg, uid)
		default:
			fmt.Println("error chat type")
		}
	}
}

func pushToRoom(app pitaya.Pitaya, ctx context.Context, msg *ChatMessage, uid string) {
	// TODO (待改进) 玩家首次发言才会加到群组中,没有加入的收不到消息
	if have, _ := app.GroupContainsMember(ctx, groupName(msg.To), uid); !have {
		err := app.GroupAddMember(ctx, groupName(msg.To), uid)
		if err != nil {
			fmt.Println("pushToRoom GroupAddMember err", err)
			return
		}
	}

	err := app.GroupBroadcast(ctx, "gate", groupName(msg.To), "chat.push.message", msg)
	if err != nil {
		fmt.Println("pushToRoom GroupBroadcast err", err)
	}
}

func groupName(roomId int64) string {
	return fmt.Sprintf("room_%d", roomId)
}
