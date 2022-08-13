package main

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var checkSitesCmd = &cobra.Command{
	Use:   "check_sites",
	Short: "check sites short",
	Long:  "check sites from inventory",
	Run: func(cmd *cobra.Command, args []string) {
		checkSites(args)
	},
}

func init() {
	RootCmd.AddCommand(checkSitesCmd)
}

func checkSites(args []string) {
	urls := return_url()
	fmt.Println("urls:", urls)
	fmt.Println("check_sites", args)

	errorNum := 0
	for i, v := range urls {

		fmt.Println("i:", i)
		fmt.Println("v:", v)
		requestURL := fmt.Sprintf("%s", v)
		res, err := http.Get(requestURL)
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			errorNum += 1
			continue
		}

		fmt.Printf("client: got response!\n")
		fmt.Printf("client: status code: %d\n", res.StatusCode)
	}
	fmt.Println("errorNum:", errorNum)
}
