package internal

import "github.com/topfreegames/pitaya/v2"

type RoomManager struct {
	app   pitaya.Pitaya
	rooms map[string]*RoomInfo
}
