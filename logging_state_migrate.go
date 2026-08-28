// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

const (
	// Provider installation messages
	ProviderInstallationStart           LogMessageType = "provider_installation_start"
	StateStoreProviderInstallationStart LogMessageType = "state_store_provider_installation_start"
	ProviderQueryUsePreviousVersion     LogMessageType = "provider_query_use_previous_version"
	ProviderQueryUsePreviousConstraints LogMessageType = "provider_query_use_constraints"
	ProviderQueryUseLatest              LogMessageType = "provider_query_use_latest"
	ProviderVersionAlreadyInstalled     LogMessageType = "provider_version_already_installed"
	ProviderVersionFoundInCacheDir      LogMessageType = "provider_version_found_in_cache_dir"
	ProviderVersionInstallationStart    LogMessageType = "provider_version_installation_start"
	ProviderVersionInstallationComplete LogMessageType = "provider_version_installation_complete"
	BuiltInProviderAvailable            LogMessageType = "built_in_provider_available"
	ThirdPartyProvidersInstalled        LogMessageType = "third_party_providers_installed"

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

type ProviderInstallationStartMessage struct {
	baseLogMessage
}

type StateStoreProviderInstallationStartMessage struct {
	baseLogMessage
}

type ProviderQueryUsePreviousVersionMessage struct {
	baseLogMessage
}

type ProviderQueryUsePreviousConstraintsMessage struct {
	baseLogMessage
}

type ProviderQueryUseLatestMessage struct {
	baseLogMessage
}

type ProviderVersionAlreadyInstalledMessage struct {
	baseLogMessage
}

type ProviderVersionFoundInCacheDirMessage struct {
	baseLogMessage
}

type ProviderVersionInstallationStartMessage struct {
	baseLogMessage
}

type ProviderVersionInstallationCompleteMessage struct {
	baseLogMessage
}

type BuiltInProviderAvailableMessage struct {
	baseLogMessage
}

type ThirdPartyProvidersInstalledMessage struct {
	baseLogMessage
}

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
