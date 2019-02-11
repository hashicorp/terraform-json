package tfjson

import (
	"encoding/json"
	"os"
	"testing"
)

func TestStateValidate(t *testing.T) {
	f, err := os.Open("test-fixtures/no_changes/plan.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var plan *Plan
	if err := json.NewDecoder(f).Decode(&plan); err != nil {
		t.Fatal(err)
	}

	if err := plan.PriorState.Validate(); err != nil {
		t.Fatal(err)
	}
}
