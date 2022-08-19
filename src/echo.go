package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "echo command",
	Long: `Simple echo command for testing.
Backticks are for multi line strings
to explain in more detail what the command does`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args:", args)
		fmt.Println("env:", rootOpts.env)
		fmt.Println("region:", rootOpts.region)
		fmt.Println("verbose:", rootOpts.verbose)
		echo(args)
	},
}

func init() {
	RootCmd.AddCommand(echoCmd)
}

func echo(args []string) {
	fmt.Println("running echo", args)
}
