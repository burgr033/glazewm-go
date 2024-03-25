package client

import (
	"encoding/json"
	"fmt"
	"github.com/burgr033/glazewm-go/types/shared"
	"github.com/gorilla/websocket"
	"net/http"
)

type WmClient struct {
	defaultPort int
	socket      *websocket.Conn
}

func NewWmClient() *WmClient {
	return &WmClient{
		defaultPort: 6123, // Default port for GlazeWM
	}
}
func (client *WmClient) Connect() error {
	socketURL := fmt.Sprintf("ws://localhost:%d", client.defaultPort)
	conn, _, err := websocket.DefaultDialer.Dial(socketURL, http.Header{})
	if err != nil {
		return err
	}
	client.socket = conn
	return nil
}
func (client *WmClient) Close() {
	if client.socket != nil {
		err := client.socket.Close()
		if err != nil {
			fmt.Println("error closing connection")
		}
	}
}
func (client *WmClient) SendMessage(message string) error {
	if client.socket == nil {
		return fmt.Errorf("socket connection is not established")
	}
	return client.socket.WriteMessage(websocket.TextMessage, []byte(message))
}
func (client *WmClient) ReadMessage() (string, error) {
	if client.socket == nil {
		return "", fmt.Errorf("socket connection is not established")
	}
	_, message, err := client.socket.ReadMessage()
	if err != nil {
		return "", err
	}
	return string(message), nil
}

func (client *WmClient) GetMonitors() ([]shared.Monitor, error) {
	if err := client.SendMessage("monitors"); err != nil {
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
func (client *WmClient) GetWindows() ([]shared.Window, error) {
	if err := client.SendMessage("windows"); err != nil {
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

func (client *WmClient) GetWorkspaces() ([]shared.Workspace, error) {
	if err := client.SendMessage("workspaces"); err != nil {
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
