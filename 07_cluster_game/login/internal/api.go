package internal

import (
	"github.com/gin-gonic/gin"

	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

type HttpHandler struct {
	cfg *config.LoginConfig
}

// Entry 获取登录地址
func (h *HttpHandler) Entry(c *gin.Context) {

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
