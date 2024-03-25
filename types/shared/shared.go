package shared

type BorderDelta struct {
	Left   int `json:"left,omitempty"`
	Top    int `json:"top,omitempty"`
	Right  int `json:"right,omitempty"`
	Bottom int `json:"bottom,omitempty"`
}

type FloatingPlacement struct {
	Left   int `json:"left,omitempty"`
	Top    int `json:"top,omitempty"`
	Right  int `json:"right,omitempty"`
	Bottom int `json:"bottom,omitempty"`
	X      int `json:"x,omitempty"`
	Y      int `json:"y,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type WindowState string

const (
	Tiling     WindowState = "tiling"
	Floating   WindowState = "floating"
	Minimized  WindowState = "minimized"
	Maximized  WindowState = "maximized"
	Fullscreen WindowState = "fullscreen"
)

type MinimizedWindow struct {
	PreviousState WindowState `json:"previousState,omitempty"`
}

type MaximizedWindow struct {
	Window
}

type FullscreenWindow struct {
	Window
}

type FloatingWindow struct {
	Window
}

type Window struct {
	Container
	FloatingPlacement FloatingPlacement `json:"floatingPlacement,omitempty"`
	BorderDelta       BorderDelta       `json:"borderDelta,omitempty"`
	Handle            int64             `json:"handle,omitempty"`
	SizePercentage    float64           `json:"sizePercentage,omitempty"`
	Children          []interface{}     `json:"children,omitempty"`
}

type ContainerType string

const (
	RootContainerType    ContainerType = "root_container"
	MonitorType          ContainerType = "monitor"
	WorkspaceType        ContainerType = "workspace"
	SplitContainerType   ContainerType = "split_container"
	FloatingWindowType   ContainerType = "floating_window"
	TilingWindowType     ContainerType = "tiling_window"
	MinimizedWindowType  ContainerType = "minimized_window"
	MaximizedWindowType  ContainerType = "maximized_window"
	FullscreenWindowType ContainerType = "fullscreen_window"
)

type RootContainer struct {
	Container
	Children []Monitor `json:"children,omitempty"`
}

type EventSubscription struct {
	SubscriptionId string `json:"subscriptionId,omitempty"`
}

type ResizeDimensions string

const (
	Height ResizeDimensions = "height"
	Width  ResizeDimensions = "width"
)

type Direction string

const (
	Left  Direction = "left"
	Right Direction = "right"
	Up    Direction = "up"
	Down  Direction = "down"
)

type TilingDirection string

const (
	Vertical   TilingDirection = "vertical"
	Horizontal TilingDirection = "horizontal"
)

type SplitContainer struct {
	Layout         TilingDirection `json:"layout,omitempty"`
	SizePercentage float64         `json:"sizePercentage,omitempty"`
	Children       []interface{}   `json:"children,omitempty"`
}

type Workspace struct {
	Container
	Layout         TilingDirection `json:"tilingDirection,omitempty"`
	SizePercentage float64         `json:"sizePercentage,omitempty"`
	Name           string          `json:"name,omitempty"`
	Children       []interface{}   `json:"children,omitempty"`
}

type Container struct {
	Id         string        `json:"id,omitempty"`
	X          int           `json:"x,omitempty"`
	Y          int           `json:"y,omitempty"`
	Width      int           `json:"width,omitempty"`
	Height     int           `json:"height,omitempty"`
	TypeO      ContainerType `json:"type,omitempty"`
	FocusIndex int           `json:"focusIndex,omitempty"`
}

type Monitor struct {
	Container
	DeviceName string      `json:"deviceName,omitempty"`
	Children   []Workspace `json:"children,omitempty"`
}
