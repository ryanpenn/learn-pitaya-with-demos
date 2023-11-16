package internal

import (
	"context"
	"fmt"
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/modules"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"net/http"
	"time"
)

type WebModule struct {
	modules.Base
	app pitaya.Pitaya
	svr *http.Server
}

func NewWebModule(app pitaya.Pitaya, c *config.LoginConfig) *WebModule {
	m := &WebModule{
		app: app,
		svr: &http.Server{
			Addr:    fmt.Sprintf(":%d", c.HttpPort),
			Handler: HttpEngine(c),
		},
	}

	return m
}

func (m *WebModule) AfterInit() {
	// start http server
	//go start(m.svr)
}

func (m *WebModule) BeforeShutdown() {
	// dump data
	//m.autoIdService.Dump()
}

func (m *WebModule) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown http server
	_ = ctx
	//shutdown(ctx, m.svr)
	return nil
}
