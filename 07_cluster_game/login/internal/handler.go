package internal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/topfreegames/pitaya/v2"
	"net/http"
	"strconv"

	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

type HttpHandler struct {
	app pitaya.Pitaya
	cfg *config.LoginConfig
}

// Login 登录
func (h *HttpHandler) Login(c *gin.Context) {
	var req LoginReq
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusUnauthorized, "no auth")
		return
	}

	addr := getGateAddr(h.app)
	c.Writer.Header().Set(ContentType, ContentTypeJson)
	c.JSON(http.StatusOK, &LoginResp{
		Code:  0,
		Token: "token",
		Addr:  addr,
	})
}

// Reg 注册
func (h *HttpHandler) Reg(c *gin.Context) {
	panic("not implement")
}

// ServerList 服务器列表
func (h *HttpHandler) ServerList(c *gin.Context) {
	var list []*ServerInfo
	if servers, err := h.app.GetServersByType("game"); err != nil {
		// handle error
		fmt.Println("no game server.")
	} else {
		for k, s := range servers {
			sid := s.Metadata["game_server_id"]
			// k, server uuid in cluster
			// sid, server id in database
			fmt.Println(sid, k)
			//
			list = append(list, &ServerInfo{
				ServerID:   sid,
				ServerName: s.Metadata["server_name"],
				ServerKey:  k,
			})
		}
	}

	c.Writer.Header().Set(ContentType, ContentTypeJson)
	c.JSON(http.StatusOK, &list)
}

func getGateAddr(app pitaya.Pitaya) string {
	servers, err := app.GetServersByType("gate")
	if err != nil {
		fmt.Println("no gate server.")
		return ""
	}

	// 获取负载最低的gate服
	var key string
	m := 0
	for k, v := range servers {
		if s := v.Metadata["conn"]; s != "" {
			if i, err := strconv.Atoi(s); err == nil {
				if i <= m {
					m = i
					key = k
				}
			}
		}
	}

	if key != "" {
		return fmt.Sprintf("%s:%s", servers[key].Hostname, servers[key].Metadata["port"])
	}

	return ""
}
