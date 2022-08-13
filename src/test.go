package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test short command",
	Long:  "testing long command",
	Run: func(cmd *cobra.Command, args []string) {
		run_test(args)
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
}

func test(args []string) {
	fmt.Println("running test", args)
	ask_continue()
	fmt.Println("lets go!")
}
