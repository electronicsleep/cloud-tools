package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test command",
	Long:  `simple test command for testing`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args:", args)
		fmt.Println("env:", rootOpts.env)
		fmt.Println("region:", rootOpts.region)
		fmt.Println("verbose:", rootOpts.verbose)
		test(args)
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
}

func test(args []string) {
	fmt.Println("run test", args)
	var config configStruct
	config.getConfig()
	sites := config.Sites
	fmt.Println("sites:", sites)
	sendMessage("cloudtools test", config)
}
