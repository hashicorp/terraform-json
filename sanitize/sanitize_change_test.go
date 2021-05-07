package sanitize

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	tfjson "github.com/hashicorp/terraform-json"
)

type testChangeCase struct {
	name     string
	old      *tfjson.Change
	expected *tfjson.Change
}

func changeCases() []testChangeCase {
	return []testChangeCase{
		{
			name: "basic",
			old: &tfjson.Change{
				Before: map[string]interface{}{
					"foo": map[string]interface{}{"a": "foo"},
					"bar": map[string]interface{}{"a": "foo"},
					"baz": map[string]interface{}{"a": "foo"},
					"qux": map[string]interface{}{
						"a": map[string]interface{}{
							"b": "foo",
						},
						"c": "bar",
					},
					"quxx": map[string]interface{}{
						"a": map[string]interface{}{
							"b": "foo",
						},
						"c": "bar",
					},
				},
				After: map[string]interface{}{
					"one":   map[string]interface{}{"x": "one"},
					"two":   map[string]interface{}{"x": "one"},
					"three": map[string]interface{}{"x": "one"},
					"four": map[string]interface{}{
						"x": map[string]interface{}{
							"y": "one",
						},
						"z": "two",
					},
					"five": map[string]interface{}{
						"x": map[string]interface{}{
							"y": "one",
						},
						"z": "two",
					},
				},
				BeforeSensitive: map[string]interface{}{
					"foo":  map[string]interface{}{},
					"bar":  true,
					"baz":  map[string]interface{}{"a": true},
					"qux":  map[string]interface{}{},
					"quxx": map[string]interface{}{"c": true},
				},
				AfterSensitive: map[string]interface{}{
					"one":   map[string]interface{}{},
					"two":   true,
					"three": map[string]interface{}{"x": true},
					"four":  map[string]interface{}{},
					"five":  map[string]interface{}{"z": true},
				},
			},
			expected: &tfjson.Change{
				Before: map[string]interface{}{
					"foo": map[string]interface{}{"a": "foo"},
					"bar": DefaultSensitiveValue,
					"baz": map[string]interface{}{"a": DefaultSensitiveValue},
					"qux": map[string]interface{}{
						"a": map[string]interface{}{
							"b": "foo",
						},
						"c": "bar",
					},
					"quxx": map[string]interface{}{
						"a": map[string]interface{}{
							"b": "foo",
						},
						"c": DefaultSensitiveValue,
					},
				},
				After: map[string]interface{}{
					"one":   map[string]interface{}{"x": "one"},
					"two":   DefaultSensitiveValue,
					"three": map[string]interface{}{"x": DefaultSensitiveValue},
					"four": map[string]interface{}{
						"x": map[string]interface{}{
							"y": "one",
						},
						"z": "two",
					},
					"five": map[string]interface{}{
						"x": map[string]interface{}{
							"y": "one",
						},
						"z": DefaultSensitiveValue,
					},
				},
				BeforeSensitive: map[string]interface{}{
					"foo":  map[string]interface{}{},
					"bar":  true,
					"baz":  map[string]interface{}{"a": true},
					"qux":  map[string]interface{}{},
					"quxx": map[string]interface{}{"c": true},
				},
				AfterSensitive: map[string]interface{}{
					"one":   map[string]interface{}{},
					"two":   true,
					"three": map[string]interface{}{"x": true},
					"four":  map[string]interface{}{},
					"five":  map[string]interface{}{"z": true},
				},
			},
		},
	}
}

func TestSanitizeChange(t *testing.T) {
	for i, tc := range changeCases() {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			actual, err := SanitizeChange(tc.old, DefaultSensitiveValue)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(tc.expected, actual); diff != "" {
				t.Errorf("SanitizeChange() mismatch (-expected +actual):\n%s", diff)
			}

			if diff := cmp.Diff(changeCases()[i].old, tc.old); diff != "" {
				t.Errorf("SanitizeChange() altered original (-expected +actual):\n%s", diff)
			}
		})
	}
}
