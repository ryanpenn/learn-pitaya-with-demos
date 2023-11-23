package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"github.com/topfreegames/pitaya/v2/errors"
)

type ChatHandler struct {
	component.Base
	app     pitaya.Pitaya
	manager *ChatManager
}

func NewChatHandler(app pitaya.Pitaya) *ChatHandler {
	return &ChatHandler{
		app: app,
	}
}

func (h *ChatHandler) AfterInit() {
	h.manager = GetManager(h.app)
}

func (h *ChatHandler) History(ctx context.Context, msg *ReqChatHistory) ([]*ChatHistory, error) {
	list, err := h.manager.getChatHistory(ctx, msg)
	if err != nil {
		return nil, errors.NewError(err, "500")
	}

	return list, nil
}

func (h *ChatHandler) Send(ctx context.Context, msg *ChatMessage) {
	if msg != nil {
		var err error
		switch msg.ChatType {
		case 0:
			err = h.manager.pushToUser(ctx, msg.From, msg.To, msg.Content)
		case 1:
			err = h.manager.pushToGroup(ctx, msg.From, msg.To, msg.Content)
		default:
			fmt.Println("error chat type")
		}

		if err != nil {
			fmt.Println("push err:", err)
		}
	}
}
