package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/topfreegames/pitaya/v2"

	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

type HttpHandler struct {
	app pitaya.Pitaya
	cfg *config.LoginConfig
}

// Login 登录
func (h *HttpHandler) Login(c *gin.Context) {

}

// Reg 注册
func (h *HttpHandler) Reg(c *gin.Context) {

}

// ServerList 服务器列表
func (h *HttpHandler) ServerList(c *gin.Context) {

}
