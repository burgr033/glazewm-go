package events

import (
	shared "github.com/burgr033/glazewm-go/types/shared"
)

// wmEvent Define an interface for all possible GlazeWM event interfaces
type wmEvent interface{}

type applicationExitingEvent struct {
}

type bindingModeChangedEvent struct {
	bindingMode string
}

type focusChangedEvent struct {
	focusedContainer shared.Container
}

type focusedContainerMovedEvent struct {
	focusedContainer shared.Container
}

type monitorAddedEvent struct {
	addedMonitor shared.Monitor
}

type monitorRemovedEvent struct {
	removedID         string
	removedDeviceName string
}

type tilingDirectionChangedEvent struct {
	newTilingDirection shared.TilingDirection
}

type userConfigReloadedEvent struct {
}

type windowManagedEvent struct {
	managedWindow shared.Window
}

type windowUnmanagedEvent struct {
	unmanagedID     string
	unmanagedHandle int
}

type workspaceActivatedEvent struct {
	activatedWorkspace shared.Workspace
}

type workspaceDeactivatedEvent struct {
	removedId   string
	removedName string
}

type workingAreaResizedEvent struct {
	affectedMonitor shared.Monitor
}

var eventMap = map[string]interface{}{
	"binding_mode_changed":     bindingModeChangedEvent{},
	"focus_changed":            focusChangedEvent{},
	"focused_container_moved":  focusedContainerMovedEvent{},
	"monitor_added":            monitorAddedEvent{},
	"monitor_removed":          monitorRemovedEvent{},
	"tiling_direction_changed": tilingDirectionChangedEvent{},
	"user_config_reloaded":     userConfigReloadedEvent{},
	"window_managed":           windowManagedEvent{},
	"window_unmanaged":         windowUnmanagedEvent{},
	"workspace_activated":      workspaceActivatedEvent{},
	"workspace_deactivated":    workspaceDeactivatedEvent{},
	"working_area_resized":     workingAreaResizedEvent{},
	"application_exiting":      applicationExitingEvent{},
}

// getEventData Utility function to get event interface for a given event type
func getEventData(eventType string) interface{} {
	return eventMap[eventType]
}
