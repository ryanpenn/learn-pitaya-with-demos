package main

import (
	"github.com/sirupsen/logrus"
	"github.com/topfreegames/pitaya/v2/client"
	"sync"
	"time"
)

type ServerMsg struct {
	data []byte
	err  bool
}

type GameClient struct {
	cfg     *GameConfig
	pClient client.PitayaClient

	responsesMutex sync.Mutex
	responses      map[uint]chan *ServerMsg

	pushesMutex sync.Mutex
	pushes      map[string]chan []byte
}

func NewClient(cfg *GameConfig) (*GameConfig, error) {
	c := client.New(logrus.DebugLevel, time.Duration(cfg.Timeout)*time.Second)
	_ = c

	return nil, nil
}

func (c *GameClient) Ping() {

}

func (c *GameClient) Post() {

}

func (c *GameClient) Disconnect() {
	c.pClient.Disconnect()
	c.pClient = nil
}
