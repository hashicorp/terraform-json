package sanitize

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	tfjson "github.com/hashicorp/terraform-json"
)

type testStateCase struct {
	name            string
	old             *tfjson.StateModule
	resourceChanges []*tfjson.ResourceChange
	mode            SanitizeStateModuleChangeMode
	expected        *tfjson.StateModule
}

func stateCases() []testStateCase {
	return []testStateCase{
		{
			name: "before",
			old: &tfjson.StateModule{
				Resources: []*tfjson.StateResource{
					{
						Address: "null_resource.foo",
						AttributeValues: map[string]interface{}{
							"foo": "bar",
							"baz": "qux",
						},
					},
				},
				Address: "",
				ChildModules: []*tfjson.StateModule{
					&tfjson.StateModule{
						Resources: []*tfjson.StateResource{
							{
								Address: "module.foo.null_resource.bar",
								AttributeValues: map[string]interface{}{
									"a": "b",
									"c": "d",
								},
							},
						},
						Address:      "module.foo",
						ChildModules: []*tfjson.StateModule{},
					},
				},
			},
			resourceChanges: []*tfjson.ResourceChange{
				{
					Address: "null_resource.foo",
					Change: &tfjson.Change{
						BeforeSensitive: map[string]interface{}{
							"baz": true,
						},
						AfterSensitive: map[string]interface{}{
							"foo": true,
						},
					},
				},
				{
					Address: "module.foo.null_resource.bar",
					Change: &tfjson.Change{
						BeforeSensitive: map[string]interface{}{
							"a": true,
						},
						AfterSensitive: map[string]interface{}{
							"c": true,
						},
					},
				},
			},
			mode: SanitizeStateModuleChangeModeBefore,
			expected: &tfjson.StateModule{
				Resources: []*tfjson.StateResource{
					{
						Address: "null_resource.foo",
						AttributeValues: map[string]interface{}{
							"foo": "bar",
							"baz": DefaultSensitiveValue,
						},
					},
				},
				Address: "",
				ChildModules: []*tfjson.StateModule{
					&tfjson.StateModule{
						Resources: []*tfjson.StateResource{
							{
								Address: "module.foo.null_resource.bar",
								AttributeValues: map[string]interface{}{
									"a": DefaultSensitiveValue,
									"c": "d",
								},
							},
						},
						Address:      "module.foo",
						ChildModules: []*tfjson.StateModule{},
					},
				},
			},
		},
		{
			name: "after",
			old: &tfjson.StateModule{
				Resources: []*tfjson.StateResource{
					{
						Address: "null_resource.foo",
						AttributeValues: map[string]interface{}{
							"foo": "bar",
							"baz": "qux",
						},
					},
				},
				Address: "",
				ChildModules: []*tfjson.StateModule{
					&tfjson.StateModule{
						Resources: []*tfjson.StateResource{
							{
								Address: "module.foo.null_resource.bar",
								AttributeValues: map[string]interface{}{
									"a": "b",
									"c": "d",
								},
							},
						},
						Address:      "module.foo",
						ChildModules: []*tfjson.StateModule{},
					},
				},
			},
			resourceChanges: []*tfjson.ResourceChange{
				{
					Address: "null_resource.foo",
					Change: &tfjson.Change{
						BeforeSensitive: map[string]interface{}{
							"baz": true,
						},
						AfterSensitive: map[string]interface{}{
							"foo": true,
						},
					},
				},
				{
					Address: "module.foo.null_resource.bar",
					Change: &tfjson.Change{
						BeforeSensitive: map[string]interface{}{
							"a": true,
						},
						AfterSensitive: map[string]interface{}{
							"c": true,
						},
					},
				},
			},
			mode: SanitizeStateModuleChangeModeAfter,
			expected: &tfjson.StateModule{
				Resources: []*tfjson.StateResource{
					{
						Address: "null_resource.foo",
						AttributeValues: map[string]interface{}{
							"foo": DefaultSensitiveValue,
							"baz": "qux",
						},
					},
				},
				Address: "",
				ChildModules: []*tfjson.StateModule{
					&tfjson.StateModule{
						Resources: []*tfjson.StateResource{
							{
								Address: "module.foo.null_resource.bar",
								AttributeValues: map[string]interface{}{
									"a": "b",
									"c": DefaultSensitiveValue,
								},
							},
						},
						Address:      "module.foo",
						ChildModules: []*tfjson.StateModule{},
					},
				},
			},
		},
	}
}

func TestSanitizeStateModule(t *testing.T) {
	for i, tc := range stateCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := SanitizeStateModule(tc.old, tc.resourceChanges, tc.mode, DefaultSensitiveValue)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("SanitizeStateModule() mismatch (-expected +actual):\n%s", diff)
			}

			if diff := cmp.Diff(stateCases()[i].old, tc.old); diff != "" {
				t.Errorf("SanitizeStateModule() altered original (-expected +actual):\n%s", diff)
			}
		})
	}
}

type testOutputCase struct {
	name     string
	old      map[string]*tfjson.StateOutput
	expected map[string]*tfjson.StateOutput
}

func outputCases() []testOutputCase {
	return []testOutputCase{
		{
			name: "basic",
			old: map[string]*tfjson.StateOutput{
				"foo": {
					Value: "bar",
				},
				"a": {
					Value:     "b",
					Sensitive: true,
				},
			},
			expected: map[string]*tfjson.StateOutput{
				"foo": {
					Value: "bar",
				},
				"a": {
					Value:     DefaultSensitiveValue,
					Sensitive: true,
				},
			},
		},
	}
}

func TestSanitizeStateOutputs(t *testing.T) {
	for i, tc := range outputCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := SanitizeStateOutputs(tc.old, DefaultSensitiveValue)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("SanitizeStateOutputs() mismatch (-expected +actual):\n%s", diff)
			}

			if diff := cmp.Diff(outputCases()[i].old, tc.old); diff != "" {
				t.Errorf("SanitizeStateOutputs() altered original (-expected +actual):\n%s", diff)
			}
		})
	}
}
