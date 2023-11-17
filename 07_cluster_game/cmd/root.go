package cmd

import (
	"fmt"
	"learn-pitaya-with-demos/cluster_game/cmd/servers"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a cmd",
	Long:  `Start a cmd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose a cmd to start")
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  `start`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("choose a serve to start")
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		//internal.InitConfig()
		//internal.InitLogger()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	// start
	startCmd.AddCommand(servers.LoginServer) // start login
}

func Execute() {
	viper.BindPFlag("gameconfig.area", startCmd.PersistentFlags().Lookup("area"))
	viper.BindPFlag("gateconfig.port", startCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("pitaya.concurrency.handler.dispatch", startCmd.PersistentFlags().Lookup("core"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
