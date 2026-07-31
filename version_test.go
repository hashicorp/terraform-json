// Copyright IBM Corp. 2019, 2026
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestVersionOutput_013(t *testing.T) {
	errOutput := `{
  "terraform_version": "0.13.5",
  "terraform_revision": "",
  "provider_selections": {
    "registry.terraform.io/hashicorp/github": "2.9.2",
    "registry.terraform.io/hashicorp/random": "3.0.0"
  },
  "terraform_outdated": true
}`
	var parsed VersionOutput
	if err := json.Unmarshal([]byte(errOutput), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &VersionOutput{
		Version: "0.13.5",
		ProviderSelections: map[string]string{
			"registry.terraform.io/hashicorp/github": "2.9.2",
			"registry.terraform.io/hashicorp/random": "3.0.0",
		},
		Outdated: true,
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

func TestVersionOutput_015(t *testing.T) {
	errOutput := `{
  "terraform_version": "0.15.0-dev",
  "terraform_revision": "ae025248cc0712bf53c675dc2fe77af4276dd5cc",
  "platform": "darwin_amd64",
  "provider_selections": {
    "registry.terraform.io/hashicorp/github": "2.9.2",
    "registry.terraform.io/hashicorp/random": "3.0.0"
  },
  "terraform_outdated": false
}`
	var parsed VersionOutput
	if err := json.Unmarshal([]byte(errOutput), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &VersionOutput{
		Version:  "0.15.0-dev",
		Revision: "ae025248cc0712bf53c675dc2fe77af4276dd5cc",
		Platform: "darwin_amd64",
		ProviderSelections: map[string]string{
			"registry.terraform.io/hashicorp/github": "2.9.2",
			"registry.terraform.io/hashicorp/random": "3.0.0",
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}

// In TF v1.17 we added the format_version field,
// to match how other static JSON outputs are versioned.
func TestVersionOutput_117(t *testing.T) {
	output := `{
  "format_version": "1.0",
  "terraform_version": "4.5.6-foo",
  "platform": "aros_riscv64",
  "provider_selections": {
    "registry.terraform.io/hashicorp/test1": "7.8.9-beta.2",
    "registry.terraform.io/hashicorp/test2": "1.2.3"
  },
  "terraform_outdated": false
}`
	var parsed VersionOutput
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Fatal(err)
	}

	expected := &VersionOutput{
		FormatVersion: "1.0",
		Version:       "4.5.6-foo",
		Platform:      "aros_riscv64",
		ProviderSelections: map[string]string{
			"registry.terraform.io/hashicorp/test1": "7.8.9-beta.2",
			"registry.terraform.io/hashicorp/test2": "1.2.3",
		},
	}
	if diff := cmp.Diff(expected, &parsed); diff != "" {
		t.Fatalf("output mismatch: %s", diff)
	}
}
