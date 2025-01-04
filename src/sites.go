package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var SitesCmd = &cobra.Command{
	Use:     "sites",
	Aliases: []string{"check", "check sites"},
	Short:   "check sites",
	Long:    `check sites from inventory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args:", args)
		fmt.Println("verbose:", rootOpts.verbose)
		sites(args, rootOpts.verbose)
	},
}

func init() {
	RootCmd.AddCommand(SitesCmd)
}

func sites(args []string, verbose bool) {
	var config configStruct
	config.getConfig()
	sites := config.Sites
	if verbose == true {
		fmt.Println("verbose on")
		fmt.Println("sites:", sites)
		fmt.Println("check-sites:", args)
	}

	errorNum := 0
	for _, s := range sites {
		fmt.Println("site:", s)
		requestURL := fmt.Sprintf("%s", s)
		res, err := http.Get(requestURL)
		if err != nil {
			fmt.Printf("ERROR: http request: %s\n", err)
			sendMessage("error with site: "+s, config)
			errorNum += 1
			continue
		} else {
			fmt.Printf("status: %d\n", res.StatusCode)
		}
	}
	fmt.Println("errorNum:", errorNum)
}
