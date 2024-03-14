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

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:   "del",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Short: "Delete a variable from cache",
	Long: `Delete a variable by key. For example:

cr delete key.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := http_client.InitCacheVariableClient(flags.Port); err != nil {
			if flags.ShowErrMsg {
				fmt.Println("error:", err.Error())
			} else {
				fmt.Println(flags.DefaultValue)
			}
			return
		}

		err := client_handlers.CacheVariableDel(context.Background(), args[0])
		if err != nil {
			if flags.ShowErrMsg {
				fmt.Println("error:", err.Error())
			} else {
				fmt.Println(flags.DefaultValue)
			}
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)
	delCmd.Flags().Int32VarP(&flags.Port, "port", "p", 6366, "specific the port of target server")
	delCmd.Flags().BoolVarP(&flags.ShowErrMsg, "showErrMsg", "", false, "show error message when internal error")
	delCmd.Flags().StringVarP(&flags.DefaultValue, "defaultValue", "d", "", "specific the defaultValue when internal error and not show error message")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// delCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// delCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
