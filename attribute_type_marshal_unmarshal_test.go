package tfjson

import (
	"encoding/json"
	"testing"
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

			if tc.expected != actual {
				t.Fatalf("expected %#v, got %#v", tc.expected, actual)
			}

			if !tc.expected.Equals(actual) {
				t.Fatal("could not validate equality with Equals method")
			}

			if tc.in != string(actualOut) {
				t.Fatalf("JSON output mismatch: expected %q, got %q", tc.in, actualOut)
			}
		})
	}
}
