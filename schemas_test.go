package tfjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestProviderSchemasValidate(t *testing.T) {
	f, err := os.Open("testdata/basic/schemas.json")
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
	f, err := os.Open("testdata/nested_attributes/schemas.json")
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
