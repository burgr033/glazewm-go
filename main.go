package main

import (
	wmclient "github.com/burgr033/glazewm-go/client"
	"github.com/burgr033/glazewm-go/command"
	"log"
)

func main() {
	// Create a new instance of WmClient
	client := wmclient.NewWmClient()

	// Connect to the GlazeWM server
	err := client.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to GlazeWM server: %v", err)
	}
	defer client.Close()
	_, _ = client.GetMonitors()
	_ = client.RunCommand(command.ReloadConfig)
}
