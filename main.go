package main

import (
	"fmt"
	args "github.com/JorgenEvens/ddns-client/arguments"
	"os"
)

func usage() {
	fmt.Println("Usage: " + os.Args[0] + " (configure [--username=<username>] [--endpoint=<username>] | run) [--config=<file>]")
}

func main() {
	provided, action := args.Get(0)

	if !provided {
		usage()
		return
	}

	if action == "configure" {
		configure()
	} else if action == "run" {
		run()
	} else {
		usage()
	}
}
