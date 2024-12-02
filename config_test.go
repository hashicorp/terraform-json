// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestConfigValidate(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/basic/plan.json"))
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var plan *Plan
	if err := json.NewDecoder(f).Decode(&plan); err != nil {
		t.Fatal(err)
	}

	if err := plan.Config.Validate(); err != nil {
		t.Fatal(err)
	}
}
