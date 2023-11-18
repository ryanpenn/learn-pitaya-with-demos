package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2/client"
	"github.com/topfreegames/pitaya/v2/conn/message"
	"github.com/topfreegames/pitaya/v2/session"
	"io"
	"net/http"
	"sync"
	"time"
)

type ServerMsg struct {
	data []byte
	err  bool
}

type GameClient struct {
	cfg   *GameConfig
	token string

	serverAddr string
	pClient    client.PitayaClient

	responsesMutex sync.Mutex
	responses      map[uint]chan *ServerMsg

	pushesMutex sync.Mutex
	pushes      map[string]chan []byte
}

func NewClient(cfg *GameConfig) *GameClient {
	return &GameClient{
		cfg:       cfg,
		pClient:   client.New(logrus.DebugLevel, time.Duration(cfg.Timeout)*time.Second),
		responses: map[uint]chan *ServerMsg{},
		pushes:    map[string]chan []byte{},
	}
}

func (c *GameClient) Ping() error {
	if resp, err := http.Get(fmt.Sprintf("http://%s/ping", c.cfg.Host)); err != nil {
		return err
	} else {
		defer resp.Body.Close()
		if data, err := io.ReadAll(resp.Body); err != nil {
			return err
		} else {
			fmt.Printf("%d - %s\n", resp.StatusCode, string(data))
			return nil
		}
	}
}

func (c *GameClient) Post(url string, data []byte, auth bool) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%s/%s", c.cfg.Host, url), bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	if auth {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}

	cli := &http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("%d - size of body: %d\n", resp.StatusCode, len(body))
	return body, nil
}

func (c *GameClient) SetToken(token string) {
	c.token = token
}

func (c *GameClient) GetToken() string {
	return c.token
}

func (c *GameClient) SetAddr(addr string) {
	c.serverAddr = addr
}

func (c *GameClient) Connect(data map[string]interface{}) error {
	handshake := &session.HandshakeData{
		Sys: session.HandshakeClientData{
			Platform:    "windows",
			LibVersion:  "v0.1",
			BuildNumber: "100",
			Version:     "v0.1",
		},
		User: data,
	}
	c.pClient.SetClientHandshakeData(handshake)
	return c.pClient.ConnectToWS(c.serverAddr, "")
}

func (c *GameClient) Disconnect() {
	c.pClient.Disconnect()
	c.pClient = nil
}

func (c *GameClient) Connected() bool {
	return c.pClient.ConnectedStatus()
}

func (c *GameClient) getResponseChannelForID(id uint) chan *ServerMsg {
	c.responsesMutex.Lock()
	defer c.responsesMutex.Unlock()

	if _, ok := c.responses[id]; !ok {
		c.responses[id] = make(chan *ServerMsg)
	}
	return c.responses[id]
}

func (c *GameClient) removeResponseChannelForID(id uint) {
	c.responsesMutex.Lock()
	defer c.responsesMutex.Unlock()

	delete(c.responses, id)
}

func (c *GameClient) getPushChannelForRoute(route string) chan []byte {
	c.pushesMutex.Lock()
	defer c.pushesMutex.Unlock()

	if _, ok := c.pushes[route]; !ok {
		c.pushes[route] = make(chan []byte)
	}
	return c.pushes[route]
}

func (c *GameClient) Request(route string, data []byte) (*ServerMsg, error) {
	messageID, err := c.pClient.SendRequest(route, data)
	if err != nil {
		return nil, err
	}

	ch := c.getResponseChannelForID(messageID)

	// wait for response
	select {
	case responseData := <-ch:
		return responseData, nil
	case <-time.After(time.Duration(c.cfg.Timeout) * time.Second):
		return nil, fmt.Errorf("timeout waiting for response on route %s", route)
	}
}

func (c *GameClient) Notify(route string, data []byte) error {
	err := c.pClient.SendNotify(route, data)
	return err
}

func (c *GameClient) ReceivePush(route string, timeout int) ([]byte, error) {
	ch := c.getPushChannelForRoute(route)

	select {
	case data := <-ch:
		return data, nil
	case <-time.After(time.Duration(timeout) * time.Second):
		return nil, fmt.Errorf("timeout waiting for push on route %s", route)
	}
}

func (c *GameClient) StartListening() {
	channel := c.pClient.MsgChannel()
	go func() {
		for m := range channel {
			switch m.Type {
			case message.Response:
				ch := c.getResponseChannelForID(m.ID)
				ch <- &ServerMsg{
					data: m.Data,
					err:  m.Err,
				}
				c.removeResponseChannelForID(m.ID)
			case message.Push:
				ch := c.getPushChannelForRoute(m.Route)
				ch <- m.Data
			default:
				fmt.Println("Unknown message type")
			}
		}
	}()
}
