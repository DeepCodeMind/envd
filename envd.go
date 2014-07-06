package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	VERSION = "0.1.0"
)

var options struct {
	Path  string
	Port  int
	Token string
	Auth  bool
}

var services []Service

func main() {
	flag.StringVar(&options.Path, "c", "", "Path to config directory")
	flag.IntVar(&options.Port, "p", 3050, "Port to listen on")
	flag.StringVar(&options.Token, "t", "", "Authentication token")

	flag.Parse()

	if options.Path == "" {
		options.Path = "./config"
	}

	// Do not require authentication if token is not set
	if options.Token == "" {
		options.Auth = false
	} else {
		options.Auth = true
	}

	var err error
	services, err = readServices()

	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("envd v%s\n", VERSION)
	fmt.Println("config path:", options.Path)
	fmt.Println("services detected:", len(services))
	fmt.Println("starting server on port", options.Port)

	startServer()
}
