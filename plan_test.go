// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfjson

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPlanValidate(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/basic/plan.json"))
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

func TestPlan_015(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/basic/plan-0.15.json"))
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

	expectedChange := &Change{
		Actions:         Actions{"create"},
		After:           map[string]interface{}{"ami": "boop"},
		AfterUnknown:    map[string]interface{}{"id": true},
		BeforeSensitive: false,
		AfterSensitive:  map[string]interface{}{"ami": true},
	}
	if diff := cmp.Diff(expectedChange, plan.ResourceChanges[0].Change); diff != "" {
		t.Fatalf("unexpected change: %s", diff)
	}

	expectedVariable := map[string]*ConfigVariable{
		"test_var": {
			Default:   "boop",
			Sensitive: true,
		},
	}
	if diff := cmp.Diff(expectedVariable, plan.Config.RootModule.Variables); diff != "" {
		t.Fatalf("unexpected variables: %s", diff)
	}
}

func TestPlan_withChecks(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/has_checks/plan.json"))
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

	if len(plan.Checks) == 0 {
		t.Fatal("expected checks to not be empty")
	}

	for _, c := range plan.Checks {
		for _, instance := range c.Instances {
			k := reflect.TypeOf(instance.Address.InstanceKey).Kind()
			switch k {
			case reflect.Int, reflect.Float32, reflect.Float64, reflect.String:
				t.Log("instance key is a valid type")
			default:
				t.Fatalf("unexpected type %s, expected string or int", k.String())
			}
		}
	}
}

func TestPlan_movedBlock(t *testing.T) {
	f, err := os.Open(filepath.FromSlash("testdata/moved_block/plan.json"))
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

	if plan.ResourceChanges[0].PreviousAddress != "random_id.test" {
		t.Fatalf("unexpected previous address %s, expected is random_id.test", plan.ResourceChanges[0].PreviousAddress)
	}
}

func TestPlan_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	b, err := os.ReadFile("testdata/numerics/plan.json")
	if err != nil {
		t.Fatal(err)
	}

	testCases := map[string]struct {
		useJSONNumber bool
		expected      any
	}{
		"float64": {
			expected: 1.23,
		},
		"json-number": {
			useJSONNumber: true,
			expected:      json.Number("1.23"),
		},
	}

	for name, testCase := range testCases {
		name, testCase := name, testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			plan := &Plan{}

			plan.UseJSONNumber(testCase.useJSONNumber)

			err = plan.UnmarshalJSON(b)

			if err != nil {
				t.Fatal(err)
			}

			after, ok := plan.ResourceChanges[0].Change.After.(map[string]any)

			if !ok {
				t.Fatal("plan.ResourceChanges[0].Change.After cannot be asserted as map[string]any")
			}

			attr, ok := after["configurable_attribute"]

			if !ok {
				t.Fatal("configurable attribute not found")
			}

			if diff := cmp.Diff(attr, testCase.expected); diff != "" {
				t.Errorf("unexpected difference: %s", diff)
			}
		})
	}
}
