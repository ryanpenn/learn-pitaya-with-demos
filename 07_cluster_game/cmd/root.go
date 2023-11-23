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
		fmt.Println("choose a server to start")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	// start
	startCmd.AddCommand(servers.LoginServer) // start login
	startCmd.AddCommand(servers.GameServer)  // start game
	startCmd.AddCommand(servers.ChatServer)  // start chat
	startCmd.AddCommand(servers.GateServer)  // start gate
}

func Execute() {
	//设置pitaya并发数
	viper.BindPFlag("pitaya.concurrency.handler.dispatch", startCmd.PersistentFlags().Lookup("core"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
