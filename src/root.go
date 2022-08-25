package main

import (
	"flag"
	"os"

	"github.com/spf13/cobra"
)

type rootOptions struct {
	env     string
	region  string
	verbose bool
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
	RootCmd.PersistentFlags().StringVarP(&rootOpts.env, "env", "e", "qa", "env")
	RootCmd.PersistentFlags().StringVarP(&rootOpts.region, "region", "r", "us-west-1", "region")
	RootCmd.PersistentFlags().BoolVarP(&rootOpts.verbose, "verbose", "v", false, "verbose")

	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	RootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
}
