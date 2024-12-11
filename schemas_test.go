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
