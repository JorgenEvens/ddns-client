package main

import (
	"bufio"
	"fmt"
	"github.com/JorgenEvens/ddns-client/arguments"
	"os"
	"strings"
)

func prompt(prompt string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(prompt + ": ")
	username, _ := reader.ReadString('\n')
	return strings.TrimRight(username, "\n")
}

func configure() {
	read, endpoint := arguments.Value("endpoint")
	if !read {
		endpoint = prompt("API Location")
	}

	read, username := arguments.Value("username")
	if !read {
		username = prompt("Username")
	}

	read, config := arguments.Value("config")
	if !read {
		config = "./config.json"
	}

	password := prompt("Password")

	var api Api
	api.Endpoint = endpoint

	acquireToken(&api, username, password)

	err := api.SaveConfiguration(config)
	if err != nil {
		fmt.Println(err)
	}

}

func acquireToken(api *Api, username string, password string) {
	var user User
	user.Username = username
	user.Password = password

	if user.Client == nil {
		user.Client = new(Client)
	}

	user.Client = NewClient()

	err := user.Login(api)

	if err != nil {
		fmt.Println(err)
		return
	}
}
