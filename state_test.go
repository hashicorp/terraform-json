package tfjson

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
)

func TestStateValidate_raw(t *testing.T) {
	f, err := os.Open("testdata/no_changes/state.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var state State
	if err := json.NewDecoder(f).Decode(&state); err != nil {
		t.Fatal(err)
	}

	if err := state.Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestStateUnmarshal_valid(t *testing.T) {
	f, err := os.Open("testdata/no_changes/state.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	var state State
	err = json.Unmarshal(b, &state)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStateUnmarshal_internalState(t *testing.T) {
	f, err := os.Open("testdata/no_changes/terraform.tfstate")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}

	var state State
	err = json.Unmarshal(b, &state)
	if err == nil {
		t.Fatal("expected unmarshal to fail")
	}
	got := err.Error()
	expected := "unexpected state input, format version is missing"
	if expected != got {
		t.Fatalf("error mismatch.\nexpected: %q\ngot: %q\n", expected, got)
	}
}

func TestStateValidate_fromPlan(t *testing.T) {
	f, err := os.Open("testdata/no_changes/plan.json")
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

func TestStateValidate_fromPlan110(t *testing.T) {
	f, err := os.Open("testdata/110_basic/plan.json")
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
