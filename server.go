package main

import (
	"github.com/cutedogspark/api-server-echo/service/Echo"
	"github.com/cutedogspark/api-server-echo/config"
)

const (
	echoDebug = true
	echoPort = 8080
)

func main() {
	var WebService *Echo.Service
	config.Init(WebService)
	WebService = Echo.NewServer(echoPort, echoDebug)
	WebService.Init()
	WebService.add()
	WebService.Run()
	for {
	}
}