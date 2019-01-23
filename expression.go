package tfjson

// Expression describes the format for an individual key in a
// Terraform configuration.
//
// This is usually indexed by key when referenced, ie:
// map[string]Expression, but exceptions exist, such as "count"
// expressions.
type Expression struct {
	// If the *entire* expression is a constant-defined value, this
	// will contain the Go representation of the expression's data.
	ConstantValue interface{} `json:"constant_value,omitempty"`

	// If any part of the expression contained values that were not
	// able to be resolved at parse-time, this will contain a list of
	// the referenced identifiers that caused the value to be unknown.
	References []string `json:"references,omitempty"`
}
