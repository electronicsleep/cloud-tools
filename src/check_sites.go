package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var checkSitesCmd = &cobra.Command{
	Use:     "cs",
	Aliases: []string{"check", "check-sites"},
	Short:   "check-sites",
	Long:    `check sites from inventory`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("args:", args)
		fmt.Println("verbose:", rootOpts.verbose)
		checkSites(args, rootOpts.verbose)
	},
}

func init() {
	RootCmd.AddCommand(checkSitesCmd)
}

func checkSites(args []string, verbose bool) {
	urls := return_url()
	if verbose == true {
		fmt.Println("verbose on")
		fmt.Println("urls:", urls)
		fmt.Println("check-sites:", args)
	}

	errorNum := 0
	for i, v := range urls {

		fmt.Println("i:", i)
		fmt.Println("v:", v)
		requestURL := fmt.Sprintf("%s", v)
		res, err := http.Get(requestURL)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			// postSlack("error with site" + v)
			errorNum += 1
			continue
		}

		fmt.Printf("client: status code: %d\n", res.StatusCode)
	}
	fmt.Println("errorNum:", errorNum)
}
