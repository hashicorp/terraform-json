// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestProviderSchemasValidate(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/basic/schemas.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var schemas *ProviderSchemas
	if err := json.NewDecoder(f).Decode(&schemas); err != nil {
		t.Fatal(err)
	}

	if err := schemas.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestProviderSchemasValidate_functions(t *testing.T) {
	f, err := os.Open("testdata/functions/schemas.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var schemas *ProviderSchemas
	if err := json.NewDecoder(f).Decode(&schemas); err != nil {
		t.Fatal(err)
	}

	if err := schemas.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestProviderSchemasValidate_ephemeralResources(t *testing.T) {
	f, err := os.Open("testdata/ephemeral_resources/schemas.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var schemas *ProviderSchemas
	if err := json.NewDecoder(f).Decode(&schemas); err != nil {
		t.Fatal(err)
	}

	if err := schemas.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestProviderSchemasValidate_nestedAttributes(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/nested_attributes/schemas.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var schemas *ProviderSchemas
	if err := json.NewDecoder(f).Decode(&schemas); err != nil {
		t.Fatal(err)
	}

	if err := schemas.Validate(); err != nil {
		t.Fatal(err)
	}
}
