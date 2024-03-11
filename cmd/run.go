/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/cobra"
	"qinglin.org/cache_variable/biz/component"
	"qinglin.org/cache_variable/biz/http_server"
	"qinglin.org/cache_variable/cmd/flags"
	"qinglin.org/cache_variable/util/port_check"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Args:  cobra.MatchAll(cobra.RangeArgs(0, 1), cobra.OnlyValidArgs),
	Short: "Run a cache_variable instance",
	Long: `Run a cache_variable instance. For example:

cr run -p 6366. It will run a instance at 6366 port.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !port_check.CheckPortAvaiable(flags.Port) {
			fmt.Println("error: port ", flags.Port, "is in used")
			return
		}

		component.InitComponent(flags.Size * 1024 * 1024)

		h := server.New(server.WithHostPorts(fmt.Sprintf("127.0.0.1:%d", flags.Port)))
		http_server.Register(h)
		h.Spin()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().Int32VarP(&flags.Port, "port", "p", 6366, "specific the port of target server")
	runCmd.Flags().Int32VarP(&flags.Size, "size", "s", 100, "specific the cache size, unit MB")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
