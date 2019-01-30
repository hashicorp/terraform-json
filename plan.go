package tfjson

// PlanFormatVersion is the version of the JSON plan format that is
// supported by this package.
const PlanFormatVersion = "0.1"

// ResourceMode is a string representation of the resource type found
// in certain fields in the plan.
type ResourceMode string

const (
	// DataSourceMode is the resource mode for data sources.
	DataSourceMode ResourceMode = "data"

	// ManagedResourceMode is the resource mode for managed resources.
	ManagedResourceMode ResourceMode = "managed"
)

// Plan represents the entire contents of an output Terraform plan.
type Plan struct {
	// The version of the plan format. This should always match the
	// PlanFormatVersion constant in this package, or else an unmarshal
	// will be unstable.
	FormatVersion string `json:"format_version,omitempty"`

	// The common state representation of resources within this plan.
	// This is a product of the existing state merged with the diff for
	// this plan.
	PlannedValues *StateValues `json:"planned_values,omitempty"`

	// The change operations for resources and data sources within this
	// plan.
	ResourceChanges []ResourceChange `json:"resource_changes,omitempty"`

	// The change operations for outputs within this plan.
	OutputChanges map[string]Change `json:"output_changes,omitempty"`

	// The Terraform state prior to the plan operation. This is the
	// same format as PlannedValues, without the current diff merged.
	PriorState *State `json:"prior_state,omitempty"`

	// The Terraform configuration used to make the plan.
	Config *Config `json:"configuration,omitempty"`
}

// ResourceChange is a description of an individual change action
// that Terraform plans to use to move from the prior state to a new
// state matching the configuration.
type ResourceChange struct {
	// The absolute resource address.
	Address string `json:"address,omitempty"`

	// The module portion of the above address. Omitted if the instance
	// is in the root module.
	ModuleAddress string `json:"module_address,omitempty"`

	// The resource mode.
	Mode ResourceMode `json:"mode,omitempty"`

	// The resource type, example: "aws_instance" for aws_instance.foo.
	Type string `json:"type,omitempty"`

	// The resource name, example: "foo" for aws_instance.foo.
	Name string `json:"name,omitempty"`

	// The instance key for any resources that have been created using
	// "count" or "for_each". If neither of these apply the key will be
	// empty.
	//
	// This value can be either an integer (int) or a string.
	Index interface{} `json:"index,omitempty"`

	// If set, indicates that this action applies to a "deposed" object
	// of the given instance rather than to its "current" object.
	// Omitted for changes to the current object.
	Deposed bool `json:"deposed,omitempty"`

	// The data describing the change that will be made to this object.
	Change *Change `json:"change,omitempty"`
}

// Change is the representation of a proposed change for an object.
type Change struct {
	// The action to be carried out by this change.
	Actions Actions `json:"actions,omitempty"`

	// Before and After are representations of the object value both
	// before and after the action. For create and delete actions,
	// either Before or After is unset (respectively). For no-op
	// actions, both values will be identical. After will be incomplete
	// if there are values within it that won't be known until after
	// apply.
	Before interface{} `json:"before,"`
	After  interface{} `json:"after,omitempty"`

	// A deep object of booleans that denotes any values that are
	// unknown in a resource. These values were previously referred to
	// as "computed" values.
	//
	// If the value cannot be found in this map, then its value should
	// be available within After, so long as the operation supports it.
	AfterUnknown interface{} `json:"after_unknown,omitempty"`
}
