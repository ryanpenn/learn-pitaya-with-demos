package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/topfreegames/pitaya/v2"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

func HttpEngine(app pitaya.Pitaya, c *config.LoginConfig) http.Handler {
	r := gin.Default()
	r.Use(cors.Default())

	// check run mode
	if strings.EqualFold(c.Mode, "debug") {
		gin.SetMode(gin.DebugMode)
		r.GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "pong")
		})
	} else {
		gin.SetMode(gin.ReleaseMode)
		r.Use(Timeout(time.Duration(c.HttpTimeout) * time.Second))
	}

	// routers
	g := r.Group(c.ContextPath)
	{
		h := &HttpHandler{app: app, cfg: c}
		g.POST("login", h.Login)
		g.POST("reg", h.Reg)

		// need auth
		g.POST("serverlist", Auth(c.TokenSecure), h.ServerList)
	}

	return r
}
