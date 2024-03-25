package client

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

type WmClient struct {
	defaultPort int
	socket      *websocket.Conn
}

// NewWmClient initializes the new client
func NewWmClient() *WmClient {
	return &WmClient{
		defaultPort: 6123, // Default port for GlazeWM
	}
}

// Connect connects to the websocket at the default port
func (client *WmClient) Connect() error {
	socketURL := fmt.Sprintf("ws://localhost:%d", client.defaultPort)
	conn, _, err := websocket.DefaultDialer.Dial(socketURL, http.Header{})
	if err != nil {
		return err
	}
	client.socket = conn
	return nil
}

// Close closes the connection
func (client *WmClient) Close() {
	if client.socket != nil {
		err := client.socket.Close()
		if err != nil {
			fmt.Println("error closing connection")
		}
	}
}

// SendMessage takes a string and sends it to the websocket
func (client *WmClient) SendMessage(message string) error {
	if client.socket == nil {
		return fmt.Errorf("socket connection is not established")
	}
	return client.socket.WriteMessage(websocket.TextMessage, []byte(message))
}

// ReadMessage interprets the message from socket and returns it as string
func (client *WmClient) ReadMessage() (string, error) {
	if client.socket == nil {
		return "", fmt.Errorf("socket connection is not established")
	}
	_, message, err := client.socket.ReadMessage()
	if err != nil {
		return "", err
	}
	fmt.Println(string(message))
	return string(message), nil
}
