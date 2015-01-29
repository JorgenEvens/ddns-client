package main

import (
	"fmt"
	args "github.com/JorgenEvens/ddns-client/arguments"
	"strings"
	"time"
)

func runDaemon(api *Api) {
	read, recordType := args.Value("type")
	if !read {
		recordType = "A"
	}
	recordType = strings.ToUpper(recordType)

	read, domain := args.Value("domain")
	if !read {
		fmt.Println("No domain specified")
		return
	}

	ticker := time.NewTicker(1 * time.Minute)

	for {
		api.Request("PUT", "/"+domain+"/"+recordType, nil, nil)
		<-ticker.C
	}
}

func run() {
	var api Api

	read, config := args.Value("config")
	if !read {
		config = "./config.json"
	}

	api.ReadConfiguration(config)
	runDaemon(&api)

}
