package internal

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/modules"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"sync"
)

type GameManager struct {
	modules.Base
	app pitaya.Pitaya
	cfg *config.GameConfig

	playersLock sync.Mutex
	players     map[int64]*Player
}

func NewGameManager(app pitaya.Pitaya, c *config.GameConfig) *GameManager {
	return &GameManager{
		app:     app,
		cfg:     c,
		players: map[int64]*Player{},
	}
}

func (m *GameManager) AfterInit() {
	// start http server
}

func (m *GameManager) BeforeShutdown() {
	// dump data
}

func (m *GameManager) Shutdown() error {

	return nil
}

func (m *GameManager) GetPlayer(uid int64) (*Player, bool) {
	// TODO
	return &Player{
		UserID: uid,
	}, true
}
