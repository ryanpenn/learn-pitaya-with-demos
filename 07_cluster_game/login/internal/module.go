package internal

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/modules"

	"learn-pitaya-with-demos/cluster_game/pkg/config"
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
	go start(m.svr)
}

func (m *WebModule) BeforeShutdown() {
	// dump data
}

func (m *WebModule) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown http server
	shutdown(ctx, m.svr)
	return nil
}

// 启动Http服务
func start(svr *http.Server) {
	if svr == nil {
		return
	}

	// Graceful server shutdown - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", svr.Addr)

	// Wait for kill signal of channel
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// This blocks until a signal is passed into the quit channel
	<-quit

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// shutdown the server
	shutdown(ctx, svr)
}

// 关闭Http服务
func shutdown(ctx context.Context, svr *http.Server) {
	// Shutdown server
	log.Println("Shutting down server...")
	if err := svr.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
