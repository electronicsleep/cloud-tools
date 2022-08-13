package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "echo short cmd",
	Long:  "echo long cmd",
	Run: func(cmd *cobra.Command, args []string) {
		echo(args)
	},
}

func init() {
	RootCmd.AddCommand(echoCmd)
}

func echo(args []string) {
	fmt.Println("running echo", args)
}
