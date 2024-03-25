package command

type WmCommand string

const (
	SubscribeMessageType           string = "subscribe"
	UnsubscribeMessageType         string = "unsubscribe"
	InvokeCommandMessageType       string = "command"
	GetMonitorsMessageType         string = "monitors"
	GetWorkspacesMessageType       string = "workspaces"
	GetWindowsMessageType          string = "windows"
	GetFocusedContainerMessageType string = "focused_container"
	GetBindingModeMessageType      string = "binding_mode"
)

type ClientMessage string

type ServerMessageType string

type BaseServerMessage struct {
	Success     bool              `json:"success,omitempty"`
	MessageType ServerMessageType `json:"messageType,omitempty"`
	Data        interface{}       `json:"data,omitempty"`
	Error       string
}

type ClientResponseMessage struct {
	BaseServerMessage
	ClientMessage string `json:"clientMessage,omitempty"`
}

type EventSubscriptionMessage struct {
	BaseServerMessage
	SubscriptionID string `json:"subscriptionID,omitempty"`
}

type ServerMessage interface{}

const (
	ClientResponseType    ServerMessageType = "client_response"
	EventSubscriptionType ServerMessageType = "event_subscription"
)

const (
	FocusModeToggle = "toggle"
)

const (
	FocusInDirection      WmCommand = "focus"
	FocusWorkspace        WmCommand = "focus workspace"
	MoveWindow            WmCommand = "move"
	MoveWindowToWorkspace WmCommand = "move to workspace"
	SetWindowSize         WmCommand = "set"
	ResizeWindow          WmCommand = "resize"
	ResizeWindowBorders   WmCommand = "resize borders"
	SetWindowState        WmCommand = "set"
	ToggleWindowState     WmCommand = "toggle"
	ChangeFocusMode       WmCommand = "focus mode toggle"
	ChangeTilingDirection WmCommand = "tiling direction"
	ExitWm                WmCommand = "exit wm"
	ReloadConfig          WmCommand = "reload config"
	CloseWindow           WmCommand = "close"
	ExecProcess           WmCommand = "exec"
	Ignore                WmCommand = "ignore"
)
