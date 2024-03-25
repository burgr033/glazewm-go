package main

import (
	"fmt"
	wmclient "github.com/burgr033/glazewm-go/client"
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

	// Retrieve information about monitors
	windows, err := client.GetWindows()
	if err != nil {
		log.Fatalf("Failed to get windows: %v", err)
	}
	monitors, err := client.GetMonitors()
	if err != nil {
		log.Fatalf("Failed to get windows: %v", err)
	}
	workspaces, err := client.GetWorkspaces()
	if err != nil {
		log.Fatalf("Failed to get windows: %v", err)
	}
	fmt.Printf("\n\n\nWINDOWS###################################################\n\n\n")
	for _, window := range windows {
		fmt.Printf("%+v\n", window)
		fmt.Printf("\n\n\n")
	}
	fmt.Printf("\n\n\nWORKSPACES###################################################\n\n\n")
	for _, workspace := range workspaces {
		fmt.Printf("%+v\n", workspace)
		fmt.Printf("\n\n\n")
	}
	fmt.Printf("\n\n\nMONITORS###################################################\n\n\n")
	for _, monitor := range monitors {
		fmt.Printf("%+v\n", monitor)
		fmt.Printf("\n\n\n")
	}
}
