package tfjson

// UnknownValueType is a string representation of JSON value types used to
// supply type information for unknown values.
type UnknownValueType string

const (
	UnknownValueTypeBool   UnknownValueType = "bool"
	UnknownValueTypeNumber UnknownValueType = "number"
	UnknownValueTypeString UnknownValueType = "string"
	UnknownValueTypeArray  UnknownValueType = "array"
	UnknownValueTypeObject UnknownValueType = "object"
)

// UnknownValue represents an unknown value, supplied with a Change.
type UnknownValue struct {
	// The type of the value. Can be one of "bool", "number", "string",
	// "array", or "object".
	Type UnknownValueType `json:"type,"`

	// Whether or not the value is unknown.
	Unknown bool `json:"unknown,"`

	// If Type is "array", this will be populated with a list of
	// UnknownValues that one can use to determine the status of the
	// individual elements within the array. This will only be present
	// if Unknown is false.
	ArrayValues []UnknownValue `json:"array_values,omitempty"`

	// If Type is "object", this will be populated with a map of
	// UnknownValues, indexed on object key, that one can use to
	// determine the status of the individual elements within the
	// object. This will only be present if Unknown is false.
	ObjectValues map[string]UnknownValue `json:"object_values,omitempty"`
}
