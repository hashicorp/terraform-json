package tfjson

import "testing"

func TestSchemaAttributeTypeName(t *testing.T) {
	cases := []struct {
		name    string
		subject SchemaAttributeType
	}{
		{
			name:    "bool",
			subject: SchemaAttributeTypeBool,
		},
		{
			name:    "number",
			subject: SchemaAttributeTypeNumber,
		},
		{
			name:    "string",
			subject: SchemaAttributeTypeString,
		},
		{
			name:    "list",
			subject: SchemaAttributeTypeList(SchemaAttributeTypeString),
		},
		{
			name:    "set",
			subject: SchemaAttributeTypeSet(SchemaAttributeTypeString),
		},
		{
			name:    "map",
			subject: SchemaAttributeTypeMap(SchemaAttributeTypeString),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual := tc.subject.Name()
			if tc.name != actual {
				t.Fatalf("expected %q, got %q", tc.name, actual)
			}
		})
	}
}

func TestSchemaAttributeTypeEquals(t *testing.T) {
	cases := []struct {
		name     string
		subject  SchemaAttributeType
		other    SchemaAttributeType
		mismatch bool
	}{
		{
			name:    "bool",
			subject: SchemaAttributeTypeBool,
			other:   SchemaAttributeTypeBool,
		},
		{
			name:    "number",
			subject: SchemaAttributeTypeNumber,
			other:   SchemaAttributeTypeNumber,
		},
		{
			name:    "string",
			subject: SchemaAttributeTypeString,
			other:   SchemaAttributeTypeString,
		},
		{
			name:    "list",
			subject: SchemaAttributeTypeList(SchemaAttributeTypeString),
			other:   SchemaAttributeTypeList(SchemaAttributeTypeString),
		},
		{
			name:    "set",
			subject: SchemaAttributeTypeSet(SchemaAttributeTypeString),
			other:   SchemaAttributeTypeSet(SchemaAttributeTypeString),
		},
		{
			name:    "map",
			subject: SchemaAttributeTypeMap(SchemaAttributeTypeString),
			other:   SchemaAttributeTypeMap(SchemaAttributeTypeString),
		},
		{
			name:    "nested list",
			subject: SchemaAttributeTypeList(SchemaAttributeTypeList(SchemaAttributeTypeString)),
			other:   SchemaAttributeTypeList(SchemaAttributeTypeList(SchemaAttributeTypeString)),
		},
		{
			name:    "nested set",
			subject: SchemaAttributeTypeSet(SchemaAttributeTypeList(SchemaAttributeTypeString)),
			other:   SchemaAttributeTypeSet(SchemaAttributeTypeList(SchemaAttributeTypeString)),
		},
		{
			name:    "nested map",
			subject: SchemaAttributeTypeMap(SchemaAttributeTypeList(SchemaAttributeTypeString)),
			other:   SchemaAttributeTypeMap(SchemaAttributeTypeList(SchemaAttributeTypeString)),
		},
		{
			name:     "mismatch primitive",
			subject:  SchemaAttributeTypeString,
			other:    SchemaAttributeTypeNumber,
			mismatch: true,
		},
		{
			name:     "mismatch collection (different collection type)",
			subject:  SchemaAttributeTypeList(SchemaAttributeTypeString),
			other:    SchemaAttributeTypeSet(SchemaAttributeTypeString),
			mismatch: true,
		},
		{
			name:     "mismatch collection (different element type)",
			subject:  SchemaAttributeTypeList(SchemaAttributeTypeString),
			other:    SchemaAttributeTypeList(SchemaAttributeTypeNumber),
			mismatch: true,
		},
		{
			name:     "mismatch primitive and collection",
			subject:  SchemaAttributeTypeString,
			other:    SchemaAttributeTypeList(SchemaAttributeTypeString),
			mismatch: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			switch {
			case tc.mismatch:
				if tc.subject.Equals(tc.other) {
					t.Fatal("should not be equal")
				}

			default:
				if tc.subject != tc.other {
					t.Fatalf("expected %#v, got %#v", tc.subject, tc.other)
				}

				if !tc.subject.Equals(tc.other) {
					t.Fatal("could not validate equality with Equals method")
				}
			}
		})
	}
}

func TestSchemaAttributeTypeIsPrimitiveType(t *testing.T) {
	cases := []struct {
		name         string
		subject      SchemaAttributeType
		notPrimitive bool
	}{
		{
			name:    "bool",
			subject: SchemaAttributeTypeBool,
		},
		{
			name:    "number",
			subject: SchemaAttributeTypeNumber,
		},
		{
			name:    "string",
			subject: SchemaAttributeTypeString,
		},
		{
			name:         "list",
			subject:      SchemaAttributeTypeList(SchemaAttributeTypeString),
			notPrimitive: true,
		},
		{
			name:         "set",
			subject:      SchemaAttributeTypeSet(SchemaAttributeTypeString),
			notPrimitive: true,
		},
		{
			name:         "map",
			subject:      SchemaAttributeTypeMap(SchemaAttributeTypeString),
			notPrimitive: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			switch {
			case tc.notPrimitive:
				if tc.subject.IsPrimitiveType() {
					t.Fatal("should not be primitive")
				}

			default:
				if !tc.subject.IsPrimitiveType() {
					t.Fatal("should be primitive")
				}
			}
		})
	}
}
func TestSchemaAttributeTypeIsCollectionType(t *testing.T) {
	cases := []struct {
		name          string
		subject       SchemaAttributeType
		notCollection bool
	}{
		{
			name:    "list",
			subject: SchemaAttributeTypeList(SchemaAttributeTypeString),
		},
		{
			name:    "set",
			subject: SchemaAttributeTypeSet(SchemaAttributeTypeString),
		},
		{
			name:    "map",
			subject: SchemaAttributeTypeMap(SchemaAttributeTypeString),
		},
		{
			name:          "bool",
			subject:       SchemaAttributeTypeBool,
			notCollection: true,
		},
		{
			name:          "number",
			subject:       SchemaAttributeTypeNumber,
			notCollection: true,
		},
		{
			name:          "string",
			subject:       SchemaAttributeTypeString,
			notCollection: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			switch {
			case tc.notCollection:
				if tc.subject.IsCollectionType() {
					t.Fatal("should not be collection")
				}

			default:
				if !tc.subject.IsCollectionType() {
					t.Fatal("should be collection")
				}
			}
		})
	}
}

func TestSchemaAttributeTypeIsListType(t *testing.T) {
	if !SchemaAttributeTypeList(SchemaAttributeTypeString).IsListType() {
		t.Fatal("list should be list type")
	}

	if SchemaAttributeTypeSet(SchemaAttributeTypeString).IsListType() {
		t.Fatal("set should not be list type")
	}

	if SchemaAttributeTypeString.IsListType() {
		t.Fatal("string should not be list type")
	}
}

func TestSchemaAttributeTypeIsSetType(t *testing.T) {
	if !SchemaAttributeTypeSet(SchemaAttributeTypeString).IsSetType() {
		t.Fatal("set should be set type")
	}

	if SchemaAttributeTypeList(SchemaAttributeTypeString).IsSetType() {
		t.Fatal("list should not be set type")
	}

	if SchemaAttributeTypeString.IsSetType() {
		t.Fatal("string should not be set type")
	}
}

func TestSchemaAttributeTypeIsMapType(t *testing.T) {
	if !SchemaAttributeTypeMap(SchemaAttributeTypeString).IsMapType() {
		t.Fatal("map should be map type")
	}

	if SchemaAttributeTypeList(SchemaAttributeTypeString).IsMapType() {
		t.Fatal("list should not be map type")
	}

	if SchemaAttributeTypeString.IsMapType() {
		t.Fatal("string should not be map type")
	}
}

func TestSchemaAttributeTypeElementType(t *testing.T) {
	actual := SchemaAttributeTypeList(SchemaAttributeTypeString).ElementType()
	if actual != SchemaAttributeTypeString {
		t.Fatalf("expected string, got %s", actual.Name())
	}
}
