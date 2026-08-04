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

func TestWorkspaceListOutput_CurrentWorkspace(t *testing.T) {
	t.Run("returns current workspace", func(t *testing.T) {
		wlo := &WorkspaceListOutput{
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

		current := wlo.CurrentWorkspace()
		if current == nil {
			t.Fatal("unexpected nil result")
		}
		if current.Name != "default" {
			t.Fatalf("wanted %q, got: %q", "default", current.Name)
		}
	})
	t.Run("returns nil when there isn't a current workspace", func(t *testing.T) {
		// This scenario could happen if custom workspaces exist and the
		// default workspace is selected, but there hasn't been an apply
		// yet so no default workspace actually exists yet...

		wlo := &WorkspaceListOutput{
			Workspaces: []WorkspaceListEntry{
				{
					Name:      "default",
					IsCurrent: false,
				},
				{
					Name:      "other",
					IsCurrent: false,
				},
			},
			Diagnostics: []Diagnostic{},
		}

		current := wlo.CurrentWorkspace()
		if current != nil {
			t.Fatalf("expected nil result, got: %#v", current)
		}
	})

	t.Run("returns nil when there are no workspaces", func(t *testing.T) {
		// This scenario could happen if there are no custom workspaces yet
		// and an apply hasn't yet happened in the default workspace.
		// E.g. user creates an empty Terraform project and immediately performs
		// `terraform workspace list -json`

		wlo := &WorkspaceListOutput{
			Workspaces: []WorkspaceListEntry{},
			Diagnostics: []Diagnostic{
				{
					// Mimicking `warnNoEnvsExistDiag` from the Core repo:
					// https://github.com/hashicorp/terraform/blob/527402d3fe2de2363c4587e7abd1a3b23669ca25/internal/command/workspace_command.go#L158
					Severity: "warning",
					Summary:  "Terraform cannot find any existing workspaces.",
					Detail:   "The \"default\" workspace is selected in your working directory. You can create this workspace by running \"terraform init\", by using the \"terraform workspace new\" subcommand or by including the \"-or-create\" flag with the \"terraform workspace select\" subcommand.",
				},
			},
		}

		current := wlo.CurrentWorkspace()
		if current != nil {
			t.Fatalf("expected nil result, got: %#v", current)
		}
	})
}
