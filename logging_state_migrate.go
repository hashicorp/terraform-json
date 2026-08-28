// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

const (
	// Dependency lock file messages
	ProviderLockfileCreated LogMessageType = "provider_lockfile_created"
	ProviderLockfileUpdated LogMessageType = "provider_lockfile_updated"

	// Provider trust-related messages
	ProviderInteractiveApproval  LogMessageType = "provider_interactive_approval"
	ProviderInteractiveRejection LogMessageType = "provider_interactive_rejection"
	ProviderAutomaticApproval    LogMessageType = "provider_automatic_approval"

	// State migration-related messages
	LogMigrationStart                             LogMessageType = "migration_start"
	LogMigrationComplete                          LogMessageType = "migration_complete"
	LogMigrationErrored                           LogMessageType = "migration_errored"
	LogMigrationFinalized                         LogMessageType = "migration_finalized"
	LogMigrationSourceInitializationStart         LogMessageType = "migration_source_initialization_start"
	LogMigrationSourceInitializationComplete      LogMessageType = "migration_source_initialization_complete"
	LogMigrationDestinationInitializationStart    LogMessageType = "migration_destination_initialization_start"
	LogMigrationDestinationInitializationComplete LogMessageType = "migration_destination_initialization_complete"
)

type ProviderLockfileCreatedMessage struct {
	baseLogMessage
}

type ProviderLockfileUpdatedMessage struct {
	baseLogMessage
}

type ProviderInteractiveApprovalMessage struct {
	baseLogMessage
}

type ProviderInteractiveRejectionMessage struct {
	baseLogMessage
}

type ProviderAutomaticApprovalMessage struct {
	baseLogMessage
}

type MigrationStartMessage struct {
	baseLogMessage
}

type MigrationCompleteMessage struct {
	baseLogMessage
}

type MigrationErroredMessage struct {
	baseLogMessage
}

type MigrationFinalizedMessage struct {
	baseLogMessage
}

type MigrationSourceInitializationStartMessage struct {
	baseLogMessage
}

type MigrationSourceInitializationCompleteMessage struct {
	baseLogMessage
}

type MigrationDestinationInitializationStartMessage struct {
	baseLogMessage
}

type MigrationDestinationInitializationCompleteMessage struct {
	baseLogMessage
}
