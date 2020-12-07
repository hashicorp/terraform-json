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
