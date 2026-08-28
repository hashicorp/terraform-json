// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0
package tfjson

import (
	"bytes"
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

	// init
	InitOutputMessage{},

	// query
	ListStartMessage{},
	ListResourceFoundMessage{},
	ListCompleteMessage{},

	// state migrate
	// Provider trust-related message
	ProviderInteractiveApprovalMessage{},
	ProviderInteractiveRejectionMessage{},
	ProviderAutomaticApprovalMessage{},
	// state migration
	MigrationStartMessage{},
	MigrationCompleteMessage{},
	MigrationErroredMessage{},
	MigrationFinalizedMessage{},
	MigrationSourceInitializationStartMessage{},
	MigrationSourceInitializationCompleteMessage{},
	MigrationDestinationInitializationStartMessage{},
	MigrationDestinationInitializationCompleteMessage{},
}

func unmarshalByType(t LogMessageType, b []byte) (LogMsg, error) {
	d := json.NewDecoder(bytes.NewReader(b))

	// decode numbers as json.Number to avoid losing precision
	d.UseNumber()

	switch t {

	// generic
	case MessageTypeVersion:
		v := VersionLogMessage{}
		return v, d.Decode(&v)
	case MessageTypeLog:
		v := LogMessage{}
		return v, d.Decode(&v)
	case MessageTypeDiagnostic:
		v := DiagnosticLogMessage{}
		return v, d.Decode(&v)

	// init
	case InitOutput:
		v := InitOutputMessage{}
		return v, json.Unmarshal(b, &v)

	// query
	case MessageListStart:
		v := ListStartMessage{}
		return v, d.Decode(&v)
	case MessageListResourceFound:
		v := ListResourceFoundMessage{}
		return v, d.Decode(&v)
	case MessageListComplete:
		v := ListCompleteMessage{}
		return v, d.Decode(&v)

	// state migrate
	// provider trust
	case ProviderInteractiveApproval:
		v := ProviderInteractiveApprovalMessage{}
		return v, d.Decode(&v)
	case ProviderInteractiveRejection:
		v := ProviderInteractiveRejectionMessage{}
		return v, d.Decode(&v)
	case ProviderAutomaticApproval:
		v := ProviderAutomaticApprovalMessage{}
		return v, d.Decode(&v)
	// state migration
	case LogMigrationStart:
		v := MigrationStartMessage{}
		return v, d.Decode(&v)
	case LogMigrationComplete:
		v := MigrationCompleteMessage{}
		return v, d.Decode(&v)
	case LogMigrationErrored:
		v := MigrationErroredMessage{}
		return v, d.Decode(&v)
	case LogMigrationFinalized:
		v := MigrationFinalizedMessage{}
		return v, d.Decode(&v)
	case LogMigrationSourceInitializationStart:
		v := MigrationSourceInitializationStartMessage{}
		return v, d.Decode(&v)
	case LogMigrationSourceInitializationComplete:
		v := MigrationSourceInitializationCompleteMessage{}
		return v, d.Decode(&v)
	case LogMigrationDestinationInitializationStart:
		v := MigrationDestinationInitializationStartMessage{}
		return v, d.Decode(&v)
	case LogMigrationDestinationInitializationComplete:
		v := MigrationDestinationInitializationCompleteMessage{}
		return v, d.Decode(&v)
	}

	v := UnknownLogMessage{}
	return v, d.Decode(&v)
}
