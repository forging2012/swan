package types

import (
	"encoding/json"
	"fmt"
)

const (
	EventTypeTaskHealthy      = "task_healthy"
	EventTypeTaskWeightChange = "task_weight_change"
	EventTypeTaskUnhealthy    = "task_unhealthy"
)

type TaskEvent struct {
	Type           string  `json:"type"`
	AppID          string  `json:"app_id"`
	AppAlias       string  `json:"app_alias"`
	VersionID      string  `json:"version_id"`
	AppVersion     string  `json:"app_version"`
	TaskID         string  `json:"task_id"`
	IP             string  `json:"task_ip"`
	Port           uint64  `json:"task_port"`
	Weight         float64 `json:"weihgt"`
	GatewayEnabled bool    `json:"gateway"`
}

// Format format task events to SSE text
func (e *TaskEvent) Format() []byte {
	bs, _ := json.Marshal(e)
	return []byte(fmt.Sprintf("event: %s\ndata: %s\n\n", e.Type, string(bs)))
}