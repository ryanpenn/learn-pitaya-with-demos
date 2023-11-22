package servers

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	pConfig "github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/serialize"
	"github.com/topfreegames/pitaya/v2/serialize/json"
	"learn-pitaya-with-demos/cluster_game/chat"
	"learn-pitaya-with-demos/cluster_game/game"
	"learn-pitaya-with-demos/cluster_game/gate"
	"learn-pitaya-with-demos/cluster_game/login"
	"learn-pitaya-with-demos/cluster_game/pkg/config"
)

var LoginServer = &cobra.Command{
	Use:   "login",
	Short: "start login server",
	Run: func(cmd *cobra.Command, args []string) {
		var serializer serialize.Serializer
		serializer = json.NewSerializer()
		//serializer = protobuf.NewSerializer()

		if c, err := config.Load[config.LoginConfig]("assets/config/", "login", "yaml"); err != nil {
			fmt.Println("config load err", err)
			return
		} else {
			login.Start(c, serializer, pConfig.NewConfig(viper.GetViper()))
		}
	},
}

var GameServer = &cobra.Command{
	Use:   "game",
	Short: "start game server",
	Run: func(cmd *cobra.Command, args []string) {
		var serializer serialize.Serializer
		serializer = json.NewSerializer()

		if c, err := config.Load[config.GameConfig]("assets/config/", "game", "yaml"); err != nil {
			fmt.Println("config load err", err)
			return
		} else {
			game.Start(c, serializer, pConfig.NewConfig(viper.GetViper()))
		}
	},
}

var ChatServer = &cobra.Command{
	Use:   "chat",
	Short: "start chat server",
	Run: func(cmd *cobra.Command, args []string) {
		var serializer serialize.Serializer
		serializer = json.NewSerializer()

		if c, err := config.Load[config.ChatConfig]("assets/config/", "chat", "yaml"); err != nil {
			fmt.Println("config load err", err)
			return
		} else {
			chat.Start(c, serializer, pConfig.NewConfig(viper.GetViper()))
		}
	},
}

var GateServer = &cobra.Command{
	Use:   "gate",
	Short: "start gate server",
	Run: func(cmd *cobra.Command, args []string) {
		var serializer serialize.Serializer
		serializer = json.NewSerializer()

		if c, err := config.Load[config.GateConfig]("assets/config/", "gate", "yaml"); err != nil {
			fmt.Println("config load err", err)
			return
		} else {
			gate.Start(c, serializer, pConfig.NewConfig(viper.GetViper()))
		}
	},
}
