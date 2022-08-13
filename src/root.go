package main

import (
	"flag"
	"os"

	"github.com/spf13/cobra"
)

type rootOptions struct {
	env    string
	region string
}

var rootOpts = &rootOptions{}

var RootCmd = &cobra.Command{
	Use:          "ct",
	SilenceUsage: true,
	Short:        "cloud-tools cli",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&rootOpts.env, "env", "qa", "deploy env")
	RootCmd.PersistentFlags().StringVar(&rootOpts.region, "region", "us-west-1", "region")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	RootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}
