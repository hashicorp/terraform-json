// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"testing"
)

func TestStateValidate_raw(t *testing.T) {
	cases := map[string]struct {
		statePath string
	}{
		"basic state": {
			statePath: "testdata/no_changes/state.json",
		},
		"state with identity": {
			statePath: "testdata/identity/state.json",
		},
	}

	for tn, tc := range cases {
		t.Run(tn, func(t *testing.T) {
			f, err := os.Open(tc.statePath)
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
		})
	}
}

func TestStateUnmarshal(t *testing.T) {
	testCases := map[string]struct {
		filePath      string
		expectedError error
	}{
		"valid state JSON": {
			filePath: "testdata/no_changes/state.json",
		},
		"invalid state JSON": {
			filePath:      "testdata/invalid/state.json",
			expectedError: errors.New("input is not a valid JSON"),
		},
	}

	for tn, tc := range testCases {
		t.Run(tn, func(t *testing.T) {
			f, err := os.Open(tc.filePath)
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()

			b, err := io.ReadAll(f)
			if err != nil {
				t.Fatal(err)
			}

			var state State
			err = state.UnmarshalJSON(b)

			if tc.expectedError != nil {
				if err.Error() != tc.expectedError.Error() {
					t.Fatalf("expected error %v; got %v", tc.expectedError.Error(), err.Error())
				} else if err == nil {
					t.Fatalf("expected error %v; got nil", tc.expectedError.Error())
				}
			}

			if tc.expectedError == nil && err != nil {
				t.Errorf("expected no error, got %q", err.Error())
			}
		})
	}
}

func TestStateUnmarshal_internalState(t *testing.T) {
	f, err := os.Open("testdata/no_changes/terraform.tfstate")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	b, err := io.ReadAll(f)
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
