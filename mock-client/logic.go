package main

import (
	"fmt"
	"sync"
	"time"
)

func RunBot(num int, cfg *GameConfig) {
	wg := sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(i)
		go func(index int) {
			defer wg.Done()
			name := fmt.Sprintf("player_%d", index)
			bot := NewBot(name, name, cfg)
			if bot.Init() {
				if bot.LoginOrReg() {
					list := bot.ServerList()
					if len(list) > 0 {
						if bot.ConnectToGame(list[0].ServerID) {
							defer func() {
								<-time.After(time.Duration(10) * time.Second)
								bot.Shutdown()
							}()

							// game
							bot.PlayerInfo(list[0].Role.RoleID)
							bot.DoTask()

							// chat
							bot.Chat(num)
							bot.WorldChat(list[0].ServerID)
							bot.CrossChat()
						}
					}
				}
			}
		}(i)
	}

	wg.Wait()
}
