// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestStateIdentitiesOutput_basic(t *testing.T) {
	output := `{
  "format_version": "1.0",
  "identities": {
    "test_instance.bar": {
      "id": "my-bar-id"
    },
    "test_instance.foo": {
      "id": "my-foo-id"
    }
  }
}`
	var parsed StateIdentitiesOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &StateIdentitiesOutput{
		FormatVersion: "1.0",
		Identities: map[string]map[string]any{
			"test_instance.bar": {
				"id": "my-bar-id",
			},
			"test_instance.foo": {
				"id": "my-foo-id",
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestStateIdentitiesOutput_warning(t *testing.T) {
	output := `{
  "format_version": "1.0",
  "identities": {},
  "diagnostics": [
    {
      "severity": "warning",
      "summary": "Example warning",
      "detail": "This is an example warning message."
    }
  ]
}`
	var parsed StateIdentitiesOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &StateIdentitiesOutput{
		FormatVersion: "1.0",
		Identities:    map[string]map[string]any{},
		Diagnostics: []Diagnostic{
			{
				Severity: DiagnosticSeverityWarning,
				Summary:  "Example warning",
				Detail:   "This is an example warning message.",
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestStateIdentitiesOutput_error(t *testing.T) {
	output := `{
  "format_version": "1.0",
  "identities": {},
  "diagnostics": [
    {
      "severity": "error",
      "summary": "Example error",
      "detail": "This is an example error message."
    }
  ]
}`
	var parsed StateIdentitiesOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &StateIdentitiesOutput{
		FormatVersion: "1.0",
		Identities:    map[string]map[string]any{},
		Diagnostics: []Diagnostic{
			{
				Severity: DiagnosticSeverityError,
				Summary:  "Example error",
				Detail:   "This is an example error message.",
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestStateIdentitiesOutput_missingFormatVersion(t *testing.T) {
	output := `{
  "test_instance.foo": {"id": "my-foo-id"},
  "test_instance.bar": {"id": "my-bar-id"}
}`
	var parsed StateIdentitiesOutput
	err := json.Unmarshal([]byte(output), &parsed)
	if err == nil {
		t.Fatal("expected unmarshal to fail")
	}
	got := err.Error()
	expected := "unexpected input, format version is missing"
	if expected != got {
		t.Fatalf("error mismatch.\nexpected: %q\ngot: %q\n", expected, got)
	}
}

func TestStateIdentitiesOutput_unsupportedFormatVersion(t *testing.T) {
	output := `{
  "format_version": "2.0",
  "identities": {}
}`
	var parsed StateIdentitiesOutput
	err := json.Unmarshal([]byte(output), &parsed)
	if err == nil {
		t.Fatal("expected unmarshal to fail")
	}
	got := err.Error()
	expected := `unsupported state identities format version: "2.0.0" does not satisfy ">= 1.0, < 2.0"`
	if expected != got {
		t.Fatalf("error mismatch.\nexpected: %q\ngot: %q\n", expected, got)
	}
}
