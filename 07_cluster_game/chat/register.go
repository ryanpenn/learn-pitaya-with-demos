package chat

import (
	"github.com/topfreegames/pitaya/v2"
	"github.com/topfreegames/pitaya/v2/component"
	"learn-pitaya-with-demos/cluster_game/chat/internal"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
	"strings"
)

// Register 注册服务
func Register(app pitaya.Pitaya, c *config.ChatConfig) {
	// handler
	app.Register(internal.NewChatHandler(app), component.WithName("handler"), component.WithNameFunc(strings.ToLower))
}
