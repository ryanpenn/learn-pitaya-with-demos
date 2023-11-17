package servers

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	pConfig "github.com/topfreegames/pitaya/v2/config"
	"github.com/topfreegames/pitaya/v2/serialize"
	"github.com/topfreegames/pitaya/v2/serialize/json"
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
