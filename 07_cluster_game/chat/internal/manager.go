package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/logger"
	"github.com/topfreegames/pitaya/v2/modules"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"learn-pitaya-with-demos/cluster_game/pkg/db"
	"sync"
	"time"
)

type ChatManager struct {
	modules.Base
	app pitaya.Pitaya
	cfg *config.ChatConfig

	roomsMutex sync.Mutex
	rooms      map[int64]string // id:name
}

func NewChatManager(app pitaya.Pitaya, c *config.ChatConfig) *ChatManager {
	return &ChatManager{
		app:   app,
		cfg:   c,
		rooms: make(map[int64]string),
	}
}

func GetManager(app pitaya.Pitaya) *ChatManager {
	m, _ := app.GetModule("manager")
	return m.(*ChatManager)
}

func (m *ChatManager) AfterInit() {
	// TODO load data from db
	chatDB := db.GetModule(m.app, "chat")
	_ = chatDB

	err := m.createGroup(context.TODO(), 0)
	if err != nil {
		logger.Log.Errorf("GroupCreate err %v", err)
	}
}

func (m *ChatManager) BeforeShutdown() {
	// TODO save data
}

func (m *ChatManager) getChatHistory(ctx context.Context, msg *ReqChatHistory) ([]*ChatHistory, error) {
	// 获取聊天记录
	var list []*ChatHistory

	// TODO load from db

	return list, nil
}

func (m *ChatManager) pushToUser(ctx context.Context, from, to int64, content string) error {
	_, err := m.app.SendPushToUsers("chat.push.message", &ChatMessage{
		ChatType: 0,
		From:     from,
		To:       to,
		Content:  content,
		Time:     time.Now().Unix(),
	}, []string{fmt.Sprintf("%d", to)}, "gate")
	return err
}

func (m *ChatManager) pushToGroup(ctx context.Context, from, to int64, content string) error {
	m.roomsMutex.Lock()
	_, exist := m.rooms[to]
	m.roomsMutex.Unlock()

	if exist {
		if err := m.app.GroupBroadcast(ctx, "gate", groupName(to), "chat.push.message", &ChatMessage{
			ChatType: 1,
			From:     from,
			To:       to,
			Content:  content,
			Time:     time.Now().Unix(),
		}); err != nil {
			return fmt.Errorf("pushToRoom GroupBroadcast err: %v", err)
		}

		return nil
	}

	return fmt.Errorf("group %d not exist", to)
}

func (m *ChatManager) joinToGroup(ctx context.Context, groupID, id int64) error {
	uid := fmt.Sprintf("%d", id)
	name := groupName(groupID)

	m.roomsMutex.Lock()
	_, exist := m.rooms[groupID]
	m.roomsMutex.Unlock()

	if !exist {
		return fmt.Errorf("group %s not exist", name)
	}

	if have, _ := m.app.GroupContainsMember(ctx, name, uid); !have {
		return m.app.GroupAddMember(ctx, name, uid)
	}

	return nil
}

func (m *ChatManager) leaveGroup(ctx context.Context, groupID, id int64) error {
	uid := fmt.Sprintf("%d", id)
	name := groupName(groupID)

	m.roomsMutex.Lock()
	_, exist := m.rooms[groupID]
	m.roomsMutex.Unlock()

	if !exist {
		return fmt.Errorf("group %s not exist", name)
	}

	if have, _ := m.app.GroupContainsMember(ctx, name, uid); have {
		return m.app.GroupRemoveMember(ctx, name, uid)
	}

	return nil
}

func (m *ChatManager) createGroup(ctx context.Context, groupID int64) error {
	name := groupName(groupID)

	m.roomsMutex.Lock()
	_, exist := m.rooms[groupID]
	if !exist {
		m.rooms[groupID] = name
	}
	m.roomsMutex.Unlock()

	if exist {
		return fmt.Errorf("group %s is exist", name)
	}

	return m.app.GroupCreate(ctx, name)
}

func (m *ChatManager) removeGroup(ctx context.Context, groupID int64) error {
	name := groupName(groupID)

	m.roomsMutex.Lock()
	_, exist := m.rooms[groupID]
	if exist {
		delete(m.rooms, groupID)
	}
	m.roomsMutex.Unlock()

	if !exist {
		return fmt.Errorf("group %s not exist", name)
	}

	return m.app.GroupDelete(ctx, name)
}

func groupName(roomId int64) string {
	return fmt.Sprintf("room_%d", roomId)
}
