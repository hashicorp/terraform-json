// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

// WorkspaceListOutput represents JSON output from terraform workspace list
// (available from 1.16 onwards)
type WorkspaceListOutput struct {
	FormatVersion string `json:"format_version"`

	Workspaces  []WorkspaceListEntry `json:"workspaces"`
	Diagnostics []Diagnostic         `json:"diagnostics"`
}

// WorkspaceListEntry represents a single workspace entry in the list of workspaces
type WorkspaceListEntry struct {
	Name      string `json:"name"`
	IsCurrent bool   `json:"is_current"`
}
