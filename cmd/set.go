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
	"qinglin.org/cache_variable/util/convert"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Args:  cobra.MatchAll(cobra.RangeArgs(2, 3), cobra.OnlyValidArgs),
	Short: "Add a variable to ram",
	Long: `Add a variable to ram. For example:

cr set param hello 20, it means add a variable called param value hello to ram, it will expire after 20 seconds.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := http_client.InitCacheVariableClient(flags.Port); err != nil {
			fmt.Println("error:", err.Error())
			return
		}

		expire := int32(20) // default expire
		if len(args) == 3 {
			if v, err := convert.ToAnyE[int32](args[2]); err != nil {
				fmt.Println("error:", "param expire is not number")
				return
			} else {
				expire = v
			}
		}

		if err := client_handlers.CacheVariableSet(context.Background(), args[0], args[1], expire); err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().Int32VarP(&flags.Port, "port", "p", 6366, "specific the port of target server")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
