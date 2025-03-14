// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestProviderSchemasValidate(t *testing.T) {
	cases := map[string]struct {
		testDataPath string
	}{
		"a basic provider schema is validated": {
			testDataPath: "testdata/basic/schemas.json",
		},
		"a provider schema including functions is validated": {
			testDataPath: "testdata/functions/schemas.json",
		},
		"a provider schema including ephemeral resources is validated": {
			testDataPath: "testdata/ephemeral_resources/schemas.json",
		},
		"a provider schema including a resource with write-only attribute(s) is validated": {
			testDataPath: "testdata/write_only_attribute_on_resource/schemas.json",
		},
		"a provider schema including resource identity schemas is validated": {
			testDataPath: "testdata/identity/schemas.json",
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {
			f, err := os.Open(tc.testDataPath)
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
		})
	}
}

// TestProviderSchemas_writeOnlyAttribute asserts that write-only attributes in a resource in a
// provider schema JSON file are marked as WriteOnly once decoded into a ProviderSchemas struct
func TestProviderSchemas_writeOnlyAttribute(t *testing.T) {
	f, err := os.Open("testdata/write_only_attribute_on_resource/schemas.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var schemas *ProviderSchemas
	if err := json.NewDecoder(f).Decode(&schemas); err != nil {
		t.Fatal(err)
	}

	resourceSchema := schemas.Schemas["terraform.io/builtin/terraform"].ResourceSchemas["terraform_example"]
	if resourceSchema.Block.Attributes["wo_attr"].WriteOnly != true {
		t.Fatal("expected terraform_example.wo_attr to be marked as write-only")
	}
	if resourceSchema.Block.Attributes["foo"].WriteOnly != false {
		t.Fatal("expected terraform_example.foo to not be marked as write-only")
	}
}
