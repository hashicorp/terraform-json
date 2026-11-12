// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestModulesOutput_v1_0(t *testing.T) {
	output := `{
  "format_version": "1.0",
  "modules": [
    { "key": "test", "source": "./mods/test", "version": "" },
    {
      "key": "other",
      "source": "registry.terraform.io/hashicorp/example/aws",
      "version": "1.2.3"
    }
  ]
}
`
	var parsed ModulesOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &ModulesOutput{
		FormatVersion: "1.0",
		Modules: []ModuleOutput{
			{
				Key:     "test",
				Source:  "./mods/test",
				Version: "",
			},
			{
				Key:     "other",
				Source:  "registry.terraform.io/hashicorp/example/aws",
				Version: "1.2.3",
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestModulesOutput_v1_1(t *testing.T) {
	output := `{
  "format_version": "1.1",
  "modules": [
    {
      "key": "test",
      "source": "./mods/test",
      "version": ""
    },
    {
        "key":     "other",
        "source":  "registry.terraform.io/hashicorp/example/aws",
        "version": "1.2.3"
    }
  ],
  "diagnostics": []
}
`
	var parsed ModulesOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &ModulesOutput{
		FormatVersion: "1.1",
		Modules: []ModuleOutput{
			{
				Key:     "test",
				Source:  "./mods/test",
				Version: "",
			},
			{
				Key:     "other",
				Source:  "registry.terraform.io/hashicorp/example/aws",
				Version: "1.2.3",
			},
		},
		Diagnostics: []Diagnostic{},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestModulesOutput_diagnostics(t *testing.T) {
	output := `{
  "format_version": "1.1",
  "modules": [],
  "diagnostics": [
    {
      "severity": "warning",
      "summary": "Example warning",
      "detail": "This is an example warning message."
    },
    {
      "severity": "error",
      "summary": "Module not installed",
      "detail": "This module is not yet installed. Run \"terraform init\" to install all modules required by this configuration."
    }
  ]
}`
	var parsed ModulesOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &ModulesOutput{
		FormatVersion: "1.1",
		Modules:       []ModuleOutput{},
		Diagnostics: []Diagnostic{
			{
				Severity: DiagnosticSeverityWarning,
				Summary:  "Example warning",
				Detail:   "This is an example warning message.",
			},
			{
				Severity: DiagnosticSeverityError,
				Summary:  "Module not installed",
				Detail:   "This module is not yet installed. Run \"terraform init\" to install all modules required by this configuration.",
			},
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestModulesOutput_missingFormatVersion(t *testing.T) {
	output := `{
  "modules": []
}`
	var parsed ModulesOutput
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

func TestModulesOutput_unsupportedFormatVersion(t *testing.T) {
	output := `{
  "format_version": "2.0",
  "modules": []
}`
	var parsed ModulesOutput
	err := json.Unmarshal([]byte(output), &parsed)
	if err == nil {
		t.Fatal("expected unmarshal to fail")
	}
	got := err.Error()
	expected := `unsupported modules format version: "2.0.0" does not satisfy ">= 1.0, < 2.0"`
	if expected != got {
		t.Fatalf("error mismatch.\nexpected: %q\ngot: %q\n", expected, got)
	}
}
