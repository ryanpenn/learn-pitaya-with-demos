package internal

import (
	"github.com/gin-contrib/cors"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

func HttpEngine(c *config.LoginConfig) http.Handler {
	r := gin.Default()
	r.Use(cors.Default())

	// check run mode
	if strings.EqualFold(c.Mode, "debug") {
		r.GET("/ping", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "pong")
		})
	} else {
		r.Use(Timeout(time.Duration(c.HttpTimeout) * time.Second))
	}

	// routers
	g := r.Group(c.ContextPath)
	{
		h := &HttpHandler{cfg: c}
		g.POST("entry", h.Entry)
		g.POST("login", h.Login)
		g.POST("reg", h.Reg)

		// need auth
		g.POST("list", Auth(), h.ServerList)
	}

	return r
}
