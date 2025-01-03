package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

func run_lib(args []string) {
	fmt.Println("run_lib", args)
}

// Struct for config.yaml
type configStruct struct {
	SlackURL string   `yaml:"slack_url"`
	SlackMsg string   `yaml:"slack_msg"`
	Email    string   `yaml:"email"`
	Servers  []string `yaml:"servers"`
}

// Get config from config.yaml
func (config *configStruct) getConfig() *configStruct {
	yamlFile, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Printf("ERROR: YAML file not found #%v ", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
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

func sendMessage(send_text string, config configStruct) {
	if config.SlackURL != "" {
		fmt.Println("SlackURL is set: sending message")
		postSlack(send_text, config)
	} else {
		fmt.Println("INFO: SlackURL is not set: no messages will be sent")
	}
}

func postSlack(message string, config configStruct) {
	t := time.Now()
	tf := t.Format("2006/01/02 15:04:05")
	fmt.Println(tf + " INFO: postSlack:" + message)
	send_text := tf + " " + message + ": " + config.SlackMsg

	var jsonData = []byte(`{
		"text": "` + send_text + `",
        }`)

	if is_connected() {
		request, error := http.NewRequest("POST", config.SlackURL, bytes.NewBuffer(jsonData))
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")

		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			panic(error)
		}
		defer response.Body.Close()

		fmt.Println("INFO: response Status:", response.Status)
		fmt.Println("INFO: response Headers:", response.Header)
		body, _ := io.ReadAll(response.Body)
		fmt.Println("INFO: response Body:", string(body))
	} else {
		fmt.Println("ERROR: Not connected to the net")
	}
}

func is_connected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
