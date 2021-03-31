package tfjson

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPlanValidate(t *testing.T) {
	f, err := os.Open("testdata/basic/plan.json")
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
	f, err := os.Open("testdata/basic/plan-0.15.json")
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
