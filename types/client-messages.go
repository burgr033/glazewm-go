package types

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

//
// const SubscribeMessage SubscribeMessageType
// const UnsubscribeMessage UnsubscribeMessageType
// const InvokeCommandMessage InvokeCommandMessageType
// const GetMonitorsMessage GetMonitorsMessageType
// const GetWorkspacesMessage GetWorkspacesMessageType
// const GetWindowsMessage GetWindowsMessageType
// const GetFocusedContainerMessage GetFocusedContainerMessageType
// const GetBindingModeMessage GetBindingModeMessageType
//
type ServerMessageType string

type BaseServerMessage struct {
	Success     bool
	MessageType ServerMessageType
	Data        interface{}
	Error       string
}

type ClientResponseMessage struct {
	BaseServerMessage
	ClientMessage string
}

type EventSubscriptionMessage struct {
	BaseServerMessage
	SubscriptionID string
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
	FocusInDirectionCommand      WmCommand = "focus"
	FocusWorkspaceCommand        WmCommand = "focus workspace"
	MoveWindowCommand            WmCommand = "move"
	MoveWindowToWorkspaceCommand WmCommand = "move to workspace"
	SetWindowSizeCommand         WmCommand = "set"
	ResizeWindowCommand          WmCommand = "resize"
	ResizeWindowBordersCommand   WmCommand = "resize borders"
	SetWindowStateCommand        WmCommand = "set"
	ToggleWindowStateCommand     WmCommand = "toggle"
	ChangeFocusModeCommand       WmCommand = "focus mode toggle"
	ChangeTilingDirectionCommand WmCommand = "tiling direction"
	ExitWmCommand                WmCommand = "exit wm"
	ReloadConfigCommand          WmCommand = "reload config"
	CloseWindowCommand           WmCommand = "close"
	ExecProcessCommand           WmCommand = "exec"
	IgnoreCommand                WmCommand = "ignore"
)
