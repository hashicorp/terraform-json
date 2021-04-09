package sanitize

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	tfjson "github.com/hashicorp/terraform-json"
)

type testVariablesCase struct {
	name     string
	old      map[string]*tfjson.PlanVariable
	configs  map[string]*tfjson.ConfigVariable
	expected map[string]*tfjson.PlanVariable
}

func variablesCases() []testVariablesCase {
	return []testVariablesCase{
		{
			name: "basic",
			old: map[string]*tfjson.PlanVariable{
				"foo": &tfjson.PlanVariable{
					Value: "test-foo",
				},
				"bar": &tfjson.PlanVariable{
					Value: "test-bar",
				},
			},
			configs: map[string]*tfjson.ConfigVariable{
				"foo": &tfjson.ConfigVariable{
					Sensitive: false,
				},
				"bar": &tfjson.ConfigVariable{
					Sensitive: true,
				},
			},
			expected: map[string]*tfjson.PlanVariable{
				"foo": &tfjson.PlanVariable{
					Value: "test-foo",
				},
				"bar": &tfjson.PlanVariable{
					Value: DefaultSensitiveValue,
				},
			},
		},
	}
}

func TestSanitizePlanVariables(t *testing.T) {
	for i, tc := range variablesCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := SanitizePlanVariables(tc.old, tc.configs, DefaultSensitiveValue)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("SanitizePlanVariables() mismatch (-expected +actual):\n%s", diff)
			}

			if diff := cmp.Diff(variablesCases()[i].old, tc.old); diff != "" {
				t.Errorf("SanitizePlanVariables() altered original (-expected +actual):\n%s", diff)
			}
		})
	}
}
