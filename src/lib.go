package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func run_lib(args []string) {
	fmt.Println("run_lib", args)
}

func return_url() []string {
	urls := []string{
		"https://www.imgidea.com",
		"https://www.chrisgr.com",
		"http://www.electronicsleep.com",
		"http://www.memoryecho.com",
	}
	return urls
}

func ask_continue() bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter y or n to continue")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	s = strings.ToLower(s)
	if strings.Compare(s, "y") == 0 {
		return true
	} else {
		fmt.Println("exit per user request")
		os.Exit(1)
		return false
	}
}
