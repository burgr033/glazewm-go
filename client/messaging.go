package client

import (
	"encoding/json"
	"fmt"
	"github.com/burgr033/glazewm-go/command"
	"github.com/burgr033/glazewm-go/types/shared"
	"strings"
)

// GetMonitors lists all monitors and puts them into the structs
func (client *WmClient) GetMonitors() ([]shared.Monitor, error) {
	if err := client.SendMessage(command.GetMonitorsMessageType); err != nil {
		return nil, err
	}
	jsonDataResponse, err := client.ReadMessage()
	if err != nil {
		return nil, err
	}
	var responseData shared.MonitorResponse
	if err := json.Unmarshal([]byte(jsonDataResponse), &responseData); err != nil {
		return nil, err
	}
	return responseData.Data, nil
}

// GetWindows lists all windows and puts them into the structs
func (client *WmClient) GetWindows() ([]shared.Window, error) {
	if err := client.SendMessage(command.GetWindowsMessageType); err != nil {
		return nil, err
	}
	jsonDataResponse, err := client.ReadMessage()
	if err != nil {
		return nil, err
	}
	var responseData shared.WindowResponse
	if err := json.Unmarshal([]byte(jsonDataResponse), &responseData); err != nil {
		return nil, err
	}
	return responseData.Data, nil
}

// GetWorkspaces lists all workspaces and puts them into the structs
func (client *WmClient) GetWorkspaces() ([]shared.Workspace, error) {
	if err := client.SendMessage(command.GetWorkspacesMessageType); err != nil {
		return nil, err
	}
	jsonDataResponse, err := client.ReadMessage()
	if err != nil {
		return nil, err
	}
	var responseData shared.WorkspaceResponse
	if err := json.Unmarshal([]byte(jsonDataResponse), &responseData); err != nil {
		return nil, err
	}
	return responseData.Data, nil
}

// RunCommand runs the command
func (client *WmClient) RunCommand(command command.WmCommand) error {
	var commandString strings.Builder
	commandString.WriteString("command \"")
	commandString.WriteString(string(command))
	commandString.WriteString("\"")
	fmt.Println(commandString.String())
	if err := client.SendMessage(commandString.String()); err != nil {
		return err
	}
	jsonDataResponse, err := client.ReadMessage()
	if err != nil {
		return err
	}
	fmt.Println(jsonDataResponse)
	return nil
}
