package main

import (
	"fmt"
	"sync"
	"time"
)

func RunBot(num int, cfg *GameConfig) {
	wg := &sync.WaitGroup{}
	for i := 0; i < num; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()

			name := fmt.Sprintf("player_%d", index)
			bot := NewBot(name, name, cfg)
			fmt.Println("run bot:", name)

			if bot.Init() == false {
				fmt.Println(name, "Init failed.")
				return
			}

			if bot.LoginOrReg() == false {
				fmt.Println(name, "LoginOrReg failed")
				return
			}

			list := bot.ServerList()
			if len(list) == 0 {
				fmt.Println(name, "no server available")
				return
			}

			if bot.ConnectToGame(list[0].ServerKey) {
				defer func() {
					<-time.After(time.Duration(5) * time.Second)
					fmt.Println("bot Shutdown...")
					bot.Shutdown()
				}()

				// game
				roleID := int64(0)
				if list[0] != nil && list[0].Role != nil {
					roleID = list[0].Role.RoleID
				}
				bot.PlayerInfo(roleID)
				bot.DoTask()

				// chat
				//bot.Chat(num)
				//sid, _ := strconv.Atoi(list[0].ServerID)
				//bot.WorldChat(sid)
				//bot.CrossChat()
			}

			fmt.Println(name, "end")
		}(i)
	}

	fmt.Println("wait...")
	wg.Wait()
}
