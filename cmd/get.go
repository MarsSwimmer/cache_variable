/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"qinglin.org/cache_variable/biz/http_client/client_handlers"
	"qinglin.org/cache_variable/biz/http_client"
	"qinglin.org/cache_variable/cmd/flags"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "Get a variable from ram",
	Long: `Get a variable from ram. For example:

cr get param -p 6366`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := http_client.InitCacheVariableClient(flags.Port); err != nil {
			if flags.ShowErrMsg {
				fmt.Println("error:", err.Error())
			} else {
				fmt.Println(flags.DefaultValue)
			}
			return
		}

		result, err := client_handlers.CacheVariableGet(context.Background(), args[0])
		if err != nil {
			if flags.ShowErrMsg {
				fmt.Println("error:", err.Error())
			} else {
				fmt.Println(flags.DefaultValue)
			}
			return
		}
		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().Int32VarP(&flags.Port, "port", "p", 6366, "specific the port of target server")
	getCmd.Flags().BoolVarP(&flags.ShowErrMsg, "showErrMsg", "", false, "show error message when internal error")
	getCmd.Flags().StringVarP(&flags.DefaultValue, "defaultValue", "d", "", "specific the defaultValue when internal error and not show error message")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
