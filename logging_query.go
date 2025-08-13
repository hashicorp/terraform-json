package tfjson

import "encoding/json"

const (
	MessageListStart         LogMessageType = "list_start"
	MessageListResourceFound LogMessageType = "list_resource_found"
	MessageListComplete      LogMessageType = "list_complete"
)

// ListStartMessage represents "query" result message of type "list_start"
type ListStartMessage struct {
	baseLogMessage
	Address      string                     `json:"address"`
	ResourceType string                     `json:"resource_type"`
	InputConfig  map[string]json.RawMessage `json:"input_config,omitempty"`
}

// ListResourceFoundMessage represents "query" result message of type "list_resource_found"
type ListResourceFoundMessage struct {
	baseLogMessage
	Address        string                     `json:"address"`
	DisplayName    string                     `json:"display_name"`
	Identity       map[string]json.RawMessage `json:"identity"`
	ResourceType   string                     `json:"resource_type"`
	ResourceObject map[string]json.RawMessage `json:"resource_object,omitempty"`
	Config         string                     `json:"config,omitempty"`
	ImportConfig   string                     `json:"import_config,omitempty"`
}

// ListCompleteMessage represents "query" result message of type "list_complete"
type ListCompleteMessage struct {
	baseLogMessage
	Address      string `json:"address"`
	ResourceType string `json:"resource_type"`
	Total        int    `json:"total"`
}
