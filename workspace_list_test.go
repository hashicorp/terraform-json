// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestWorkspaceListOutput_basic(t *testing.T) {
	output := `{
  "workspaces": [
    {
      "name": "default",
      "is_current": true
    },
    {
      "name": "other",
      "is_current": false
    }
  ],
  "diagnostics": []
}`
	var parsed WorkspaceListOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &WorkspaceListOutput{
		Workspaces: []WorkspaceListEntry{
			{
				Name:      "default",
				IsCurrent: true,
			},
			{
				Name:      "other",
				IsCurrent: false,
			},
		},
		Diagnostics: []Diagnostic{},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestWorkspaceListOutput_warning(t *testing.T) {
	output := `{
  "workspaces": [
    {
      "name": "default",
      "is_current": true
    },
    {
      "name": "other",
      "is_current": false
    }
  ],
  "diagnostics": [
    {
      "severity": "warning",
      "summary": "Warning: Consider yourself warned"
    }
  ]
}`
	var parsed WorkspaceListOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &WorkspaceListOutput{
		Workspaces: []WorkspaceListEntry{
			{
				Name:      "default",
				IsCurrent: true,
			},
			{
				Name:      "other",
				IsCurrent: false,
			},
		},
		Diagnostics: []Diagnostic{
			{
				Severity: "warning",
				Summary:  "Warning: Consider yourself warned",
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestWorkspaceListOutput_error(t *testing.T) {
	output := `{
  "workspaces": [],
  "diagnostics": [
    {
      "severity": "error",
      "summary": "Error: Something went wrong"
    }
  ]
}`
	var parsed WorkspaceListOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &WorkspaceListOutput{
		Workspaces: []WorkspaceListEntry{},
		Diagnostics: []Diagnostic{
			{
				Severity: "error",
				Summary:  "Error: Something went wrong",
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}
