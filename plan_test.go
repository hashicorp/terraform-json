package tfjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestPlanValidate(t *testing.T) {
	f, err := os.Open("test-fixtures/basic/plan.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var plan *Plan
	if err := json.NewDecoder(f).Decode(&plan); err != nil {
		t.Fatal(err)
	}

	if err := plan.Validate(); err != nil {
		t.Fatal(err)
	}
}
