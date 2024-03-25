package shared

type MonitorResponse struct {
	Data []Monitor `json:"data"`
}

type WorkspaceResponse struct {
	Data []Workspace `json:"data"`
}
type WindowResponse struct {
	Data []Window `json:"data"`
}
