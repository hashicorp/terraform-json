// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0
package tfjson

import (
	"encoding/json"
)

type LogMessageType string

const (
	MessageTypeVersion    LogMessageType = "version"
	MessageTypeLog        LogMessageType = "log"
	MessageTypeDiagnostic LogMessageType = "diagnostic"
)

// allLogMessageTypes is a slice containing all recognised message types
// to be passed into cmp.AllowUnexported
var allLogMessageTypes = []any{
	VersionLogMessage{},
	LogMessage{},
	DiagnosticLogMessage{},
	UnknownLogMessage{},
}

func unmarshalByType(t LogMessageType, b []byte) (LogMsg, error) {
	switch t {
	case MessageTypeVersion:
		v := VersionLogMessage{}
		return v, json.Unmarshal(b, &v)
	case MessageTypeLog:
		v := LogMessage{}
		return v, json.Unmarshal(b, &v)
	case MessageTypeDiagnostic:
		v := DiagnosticLogMessage{}
		return v, json.Unmarshal(b, &v)
	}

	v := UnknownLogMessage{}
	return v, json.Unmarshal(b, &v)
}
