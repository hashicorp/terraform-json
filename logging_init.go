// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package tfjson

const (
	InitOutput LogMessageType = "init_output"
)

// InitOutputMessage represents messages of type "init_output"
type InitOutputMessage struct {
	baseLogMessage
	MessageCode string `json:"message_code"`
}
