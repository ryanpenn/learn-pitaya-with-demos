package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
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
	rooms      map[string]*RoomInfo
}

func NewChatManager(app pitaya.Pitaya, c *config.ChatConfig) *ChatManager {
	return &ChatManager{
		app:   app,
		cfg:   c,
		rooms: make(map[string]*RoomInfo),
	}
}

func GetManager(app pitaya.Pitaya) *ChatManager {
	m, _ := app.GetModule("manager")
	return m.(*ChatManager)
}

func (m *ChatManager) AfterInit() {
	// 加载聊天群组信息
	chatDB := db.GetModule(m.app, "chat")
	_ = chatDB

	m.roomsMutex.Lock()
	defer m.roomsMutex.Unlock()

	// 创建聊天群组
	list := []int64{1, 2}
	var key string
	for _, v := range list {
		key = groupName(v)
		m.rooms[key] = &RoomInfo{
			ID:       v,
			RoomType: 0,
			Name:     groupName(v),
		}
		m.app.GroupCreate(context.Background(), key)
	}
}

func (m *ChatManager) BeforeShutdown() {
	// TODO save data
}

func (m *ChatManager) GetChatHistory(ctx context.Context, msg *ReqChatHistory) ([]*ChatHistory, error) {
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
	group := groupName(to)

	m.roomsMutex.Lock()
	_, exist := m.rooms[group]
	m.roomsMutex.Unlock()

	if exist {
		if err := m.app.GroupBroadcast(ctx, "gate", group, "chat.push.message", &ChatMessage{
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
	key := groupName(groupID)

	m.roomsMutex.Lock()
	_, exist := m.rooms[key]
	m.roomsMutex.Unlock()

	if !exist {
		return fmt.Errorf("group %d not exist", groupID)
	}

	if have, _ := m.app.GroupContainsMember(ctx, key, uid); !have {
		return m.app.GroupAddMember(ctx, key, uid)
	}

	return nil
}

func (m *ChatManager) leaveGroup(ctx context.Context, groupID, id int64) error {
	uid := fmt.Sprintf("%d", id)
	key := groupName(groupID)

	m.roomsMutex.Lock()
	_, exist := m.rooms[key]
	m.roomsMutex.Unlock()

	if !exist {
		return fmt.Errorf("group %d not exist", groupID)
	}

	if have, _ := m.app.GroupContainsMember(ctx, key, uid); have {
		return m.app.GroupRemoveMember(ctx, key, uid)
	}

	return nil
}

func groupName(roomId int64) string {
	return fmt.Sprintf("room_%d", roomId)
}
