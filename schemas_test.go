// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/zclconf/go-cty/cty"
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
		"a provider schema including list resource schemas is validated": {
			testDataPath: "testdata/list_resources/schemas.json",
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

func TestProviderSchemas_action(t *testing.T) {
	expectedAction := &ActionSchema{
		Block: &SchemaBlock{
			DescriptionKind: SchemaDescriptionKindPlain,
			Attributes: map[string]*SchemaAttribute{
				"program": {
					AttributeType:   cty.List(cty.String),
					Description:     "A list of strings, whose first element is the program to run and whose subsequent elements are optional command line arguments to the program.",
					DescriptionKind: SchemaDescriptionKindPlain,
					Required:        true,
				},
				"query": {
					AttributeType:   cty.Map(cty.String),
					Description:     "A map of string values to pass to the external program as the query arguments. If not supplied, the program will receive an empty object as its input.",
					DescriptionKind: SchemaDescriptionKindPlain,
					Optional:        true,
				},
				"working_dir": {
					AttributeType:   cty.String,
					Description:     "Working directory of the program. If not supplied, the program will run in the current directory.",
					DescriptionKind: SchemaDescriptionKindPlain,
					Optional:        true,
				},
			},
		},
	}

	f, err := os.Open("testdata/actions/schemas.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var schemas *ProviderSchemas
	if err := json.NewDecoder(f).Decode(&schemas); err != nil {
		t.Fatal(err)
	}

	gotAction := schemas.Schemas["registry.terraform.io/hashicorp/external"].ActionSchemas["external"]
	if diff := cmp.Diff(gotAction, expectedAction, cmpopts.EquateComparable(cty.Type{})); diff != "" {
		t.Errorf("Unexpected diff (+wanted, -got): %s", diff)
		return
	}
}
