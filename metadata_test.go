// Copyright IBM Corp. 2019, 2025
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestMetadataFunctionsValidate(t *testing.T) {
	f, err := os.Open("testdata/basic/functions.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var functions *MetadataFunctions
	if err := json.NewDecoder(f).Decode(&functions); err != nil {
		t.Fatal(err)
	}

	if err := functions.Validate(); err != nil {
		t.Fatal(err)
	}
}
