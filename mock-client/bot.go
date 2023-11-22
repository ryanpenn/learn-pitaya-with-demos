package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Bot struct {
	Name     string
	Password string
	uid      int64
	client   *GameClient
	stopCh   chan struct{}
}

func NewBot(name, pass string, cfg *GameConfig) *Bot {
	return &Bot{
		Name:     name,
		Password: pass,
		client:   NewClient(cfg),
		stopCh:   make(chan struct{}),
	}
}

func (b *Bot) Init() bool {
	if err := b.client.Ping(); err != nil {
		fmt.Println("Ping err", err)
		return false
	}
	return true
}

func (b *Bot) LoginOrReg() bool {
	req, _ := json.Marshal(&LoginReq{
		Name:     b.Name,
		Password: b.Password,
	})

	resp, err := b.client.Post("api/login", req, false)
	if err != nil {
		fmt.Println("Login err", err)
		return false
	}
	if resp == nil {
		resp, err = b.client.Post("api/reg", req, false)
		if err != nil {
			fmt.Println("Reg err", err)
			return false
		}
	}

	var l LoginResp
	err = json.Unmarshal(resp, &l)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return false
	}

	fmt.Println("LoginOrReg success", l.Code)
	b.client.SetAddr(l.Addr)
	b.client.SetToken(l.Token)
	return true
}

func (b *Bot) ServerList() []*ServerInfo {
	var list []*ServerInfo

	resp, err := b.client.Post("api/serverlist", nil, true)
	if err != nil {
		fmt.Println("Server list err", err)
		return nil
	}

	err = json.Unmarshal(resp, &list)
	if err != nil {
		fmt.Println("Unmarshal err", err)
		return nil
	}

	return list
}

func (b *Bot) ConnectToGame(info *ServerInfo) bool {
	err := b.client.Connect(map[string]interface{}{
		"token": b.client.GetToken(),
		"id":    info.ServerID,
		"key":   info.ServerKey,
	})
	if err != nil {
		fmt.Println("Connect err", err)
		return false
	}

	// listen server msg
	go b.client.StartListening()
	go func() {
		for {
			select {
			case <-b.stopCh:
				return
			default:
				b.OnPlayerUpdate(b.client.ReceivePush("game.push.playerupdate", 1))
				b.OnChatMessage(b.client.ReceivePush("chat.push.message", 1))
			}
		}
	}()

	return true
}

func (b *Bot) Shutdown() {
	b.stopCh <- struct{}{}
	b.client.Disconnect()
	close(b.stopCh)
}

func (b *Bot) OnPlayerUpdate(data []byte, err error) {
	if err != nil {
		return
	}

	info := PlayerUpdateInfo{}
	err = json.Unmarshal(data, &info)
	if err != nil {
		fmt.Println("OnPlayerUpdate Unmarshal err", err)
		return
	}

	fmt.Println("PlayerUpdate...")
}

func (b *Bot) OnChatMessage(data []byte, err error) {
	if err != nil {
		return
	}

	msg := ChatMessage{}
	err = json.Unmarshal(data, &msg)
	if err != nil {
		fmt.Println("OnChatMessage Unmarshal err", err)
		return
	}

	switch msg.ChatType {
	case 0:
		fmt.Printf("收到用户 %d 的消息：%s \n", msg.From, msg.Content)
	case 1:
		fmt.Printf("收到用户 %d 在世界频道的消息：%s \n", msg.From, msg.Content)
	case 2:
		fmt.Printf("收到用户 %d 在跨服频道的消息：%s \n", msg.From, msg.Content)
	}

	fmt.Println("OnChatMessage", msg)
}

func (b *Bot) PlayerInfo(roleID int64) {
	route := "game.handler.playerinfo"
	req := PlayerInfoReq{
		ID: roleID,
	}
	data, _ := json.Marshal(req)
	resp, err := b.client.Request(route, data)
	if err != nil {
		fmt.Println("PlayerInfo Request err", err)
		return
	}

	var info PlayerInfo
	err = json.Unmarshal(resp.data, &info)
	if err != nil {
		fmt.Println("PlayerInfo Unmarshal err", err)
		return
	}

	b.uid = info.ID
	fmt.Println("PlayerInfo", info)
}

func (b *Bot) DoTask() {
	route := "game.handler.dotask"
	taskID := rand.Int63n(5) + 1
	data, _ := json.Marshal(&TaskReq{
		ID: taskID,
	})
	err := b.client.Notify(route, data)
	if err != nil {
		fmt.Println("DoTask Notify err", err)
		return
	}
}

func (b *Bot) Chat(num int) {
	route := "chat.handler.send"

	// 随机发给一位玩家(包括自己）
	targetID := int64(rand.Intn(num) + 1)
	msg := ChatMessage{
		ChatType: 0, // 单聊
		From:     b.uid,
		To:       targetID,
		Content:  fmt.Sprintf("单聊消息"),
		Time:     time.Now().Unix(),
	}
	data, _ := json.Marshal(&msg)
	err := b.client.Notify(route, data)
	if err != nil {
		fmt.Println("Chat Notify err", err)
		return
	}
}

func (b *Bot) WorldChat(serverID int) {
	route := "chat.handler.send"
	msg := ChatMessage{
		ChatType: 1, // 群聊
		From:     b.uid,
		To:       int64(serverID),
		Content:  fmt.Sprintf("世界消息"),
		Time:     time.Now().Unix(),
	}
	data, _ := json.Marshal(&msg)
	err := b.client.Notify(route, data)
	if err != nil {
		fmt.Println("WorldChat Notify err", err)
		return
	}
}

func (b *Bot) CrossChat() {
	route := "chat.handler.send"
	msg := ChatMessage{
		ChatType: 1, // 群聊
		From:     b.uid,
		To:       0, // 跨服
		Content:  fmt.Sprintf("跨服消息"),
		Time:     time.Now().Unix(),
	}
	data, _ := json.Marshal(&msg)
	err := b.client.Notify(route, data)
	if err != nil {
		fmt.Println("CrossChat Notify err", err)
		return
	}
}
