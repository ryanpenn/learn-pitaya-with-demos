package client

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	v2client "github.com/topfreegames/pitaya/v2/client"
	"github.com/vmihailenco/msgpack"

	"learn-pitaya-with-demos/msgp/msg"
)

func Run() {
	// 通过pitaya提供的client连接
	c := v2client.New(logrus.InfoLevel, 100*time.Second)
	err := c.ConnectTo("127.0.0.1:9001")
	if err != nil {
		fmt.Println("conn server err:", err)
		return
	}

	go func(c *v2client.Client) {
		// 处理服务器消息
		for data := range c.MsgChannel() {
			if data.Err {
				fmt.Println("data err:", string(data.Data))
				break
			}

			// Type:
			// 	Request  Type = 0x00  客户端请求
			//	Notify   Type = 0x01  客户端通知
			//	Response Type = 0x02  服务器返回
			//	Push     Type = 0x03  服务器推送

			if data.Type == 2 {
				var m msg.Response
				if err := msgpack.Unmarshal(data.Data, &m); err != nil {
					fmt.Println("err:", err)
					break
				}
				fmt.Println("Response ----->", m)
			} else {
				// Route 推送消息的客户端路由地址
				if data.Route == "message" {
					var m msg.Message
					if err := msgpack.Unmarshal(data.Data, &m); err != nil {
						fmt.Println("err:", err)
						break
					}
					fmt.Println("Message ----->", m)
				}
			}

			// raw message.
			// fmt.Println("data ----->", string(data.Data))
		}
	}(c)

	// 使用 msgpack 序列化
	m, _ := msgpack.Marshal(&msg.NewUser{Name: "client"})

	for {
		// 服务器类型.组件名.方法名
		if _, err := c.SendRequest("chat.room.join", m); err != nil {
			fmt.Println("send request err:", err)
			return
		}
		time.Sleep(time.Second)
	}
}
