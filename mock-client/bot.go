package main

import (
	"encoding/json"
	"fmt"
)

type Bot struct {
	Name     string
	Password string
	uid      string
	client   *GameClient
	stopCh   chan struct{}
}

func NewBot(name, pass string, cfg *GameConfig) *Bot {
	return &Bot{
		Name:     name,
		Password: pass,
		client:   NewClient(cfg),
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

func (b *Bot) ConnectToGame(serverID int) bool {
	err := b.client.Connect(map[string]interface{}{
		"token": b.client.GetToken(),
		"id":    serverID,
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
				b.OnPlayerUpdate(b.client.ReceivePush("game.playerupdate", 1))
			}
		}
	}()

	return true
}

func (b *Bot) Shutdown() {
	b.stopCh <- struct{}{}
	close(b.stopCh)
	b.client.Disconnect()
	fmt.Println("bot shutdown...bye")
}

func (b *Bot) OnPlayerUpdate(data []byte, err error) {
	if err != nil {
		return
	}

	fmt.Println("OnPlayerUpdate", data)
}
