package tfjson

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func TestSchemaAttributeTypeMarshalUnmarshal(t *testing.T) {
	cases := []struct {
		name     string
		in       string
		expected SchemaAttributeType
	}{
		{
			name:     "bool",
			in:       `"bool"`,
			expected: SchemaAttributeTypeBool,
		},
		{
			name:     "number",
			in:       `"number"`,
			expected: SchemaAttributeTypeNumber,
		},
		{
			name:     "string",
			in:       `"string"`,
			expected: SchemaAttributeTypeString,
		},
		{
			name:     "list",
			in:       `["list","string"]`,
			expected: SchemaAttributeTypeList(SchemaAttributeTypeString),
		},
		{
			name:     "set",
			in:       `["set","string"]`,
			expected: SchemaAttributeTypeSet(SchemaAttributeTypeString),
		},
		{
			name:     "map",
			in:       `["map","string"]`,
			expected: SchemaAttributeTypeMap(SchemaAttributeTypeString),
		},
		{
			name:     "nested list",
			in:       `["list",["list","string"]]`,
			expected: SchemaAttributeTypeList(SchemaAttributeTypeList(SchemaAttributeTypeString)),
		},
		{
			name:     "nested set",
			in:       `["set",["list","string"]]`,
			expected: SchemaAttributeTypeSet(SchemaAttributeTypeList(SchemaAttributeTypeString)),
		},
		{
			name:     "nested map",
			in:       `["map",["list","string"]]`,
			expected: SchemaAttributeTypeMap(SchemaAttributeTypeList(SchemaAttributeTypeString)),
		},
		{
			name: "object",
			in:   `["object",{"bar":"string","foo":"number"}]`,
			expected: SchemaAttributeTypeObject(
				map[string]SchemaAttributeType{
					"bar": SchemaAttributeTypeString,
					"foo": SchemaAttributeTypeNumber,
				},
			),
		},
		{
			name: "object nested in list",
			in:   `["list",["object",{"bar":"string","foo":"number"}]]`,
			expected: SchemaAttributeTypeList(
				SchemaAttributeTypeObject(
					map[string]SchemaAttributeType{
						"bar": SchemaAttributeTypeString,
						"foo": SchemaAttributeTypeNumber,
					},
				),
			),
		},
		{
			name: "object nested in set",
			in:   `["set",["object",{"bar":"string","foo":"number"}]]`,
			expected: SchemaAttributeTypeSet(
				SchemaAttributeTypeObject(
					map[string]SchemaAttributeType{
						"bar": SchemaAttributeTypeString,
						"foo": SchemaAttributeTypeNumber,
					},
				),
			),
		},
		{
			name: "object with nested types",
			in:   `["object",{"bar":"string","foo":["list",["object",{"baz":"string"}]]}]`,
			expected: SchemaAttributeTypeObject(
				map[string]SchemaAttributeType{
					"bar": SchemaAttributeTypeString,
					"foo": SchemaAttributeTypeList(
						SchemaAttributeTypeObject(
							map[string]SchemaAttributeType{
								"baz": SchemaAttributeTypeString,
							},
						),
					),
				},
			),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var actual SchemaAttributeType
			if err := json.Unmarshal([]byte(tc.in), &actual); err != nil {
				t.Fatal(err)
			}

			actualOut, err := json.Marshal(actual)
			if err != nil {
				t.Fatalf("marshal err: %s", err)
			}

			if !tc.expected.Equals(actual) {
				t.Fatalf("\nEquals: failed\nexpected:\n\n%s\n\ngot:\n\n%s\n\n", spew.Sdump(tc.expected), spew.Sdump(actual))
			}

			if tc.in != string(actualOut) {
				t.Fatalf("JSON output mismatch: expected %s, got %s", tc.in, actualOut)
			}
		})
	}
}
