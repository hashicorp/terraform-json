package tfjson

import (
	"encoding/json"
	"errors"
	"strings"
)

// Config represents the complete configuration source
type Config struct {
	// A map of all provider instances across all modules in the
	// configuration.
	//
	// The index for this field is opaque and should not be parsed. Use
	// the individual fields in ProviderConfig to discern actual data
	// about the provider such as name, alias, or defined module.
	ProviderConfigs map[string]*ProviderConfig `json:"provider_config,omitempty"`

	// The root module in the configuration. Any child modules descend
	// off of here.
	RootModule *ConfigModule `json:"root_module,omitempty"`
}

// Validate checks to ensure that the config is present.
func (c *Config) Validate() error {
	if c == nil {
		return errors.New("config is nil")
	}

	return nil
}

// ProviderConfig describes a provider configuration instance.
type ProviderConfig struct {
	// The name of the provider, ie: "aws".
	Name string `json:"name,omitempty"`

	// The alias of the provider, ie: "us-east-1".
	Alias string `json:"alias,omitempty"`

	// The address of the module the provider is declared in.
	ModuleAddress string `json:"module_address,omitempty"`

	// Any non-special configuration values in the provider, indexed by
	// key.
	Expressions map[string]*Expression `json:"expressions,omitempty"`

	// The defined version constraint for this provider.
	VersionConstraint string `json:"version_constraint,omitempty"`
}

// ConfigModule describes a module in Terraform configuration.
type ConfigModule struct {
	// The outputs defined in the module.
	Outputs map[string]*ConfigOutput `json:"outputs,omitempty"`

	// The resources defined in the module.
	Resources []*ConfigResource `json:"resources,omitempty"`

	// Any "module" stanzas within the specific module.
	ModuleCalls map[string]*ModuleCall `json:"module_calls,omitempty"`

	// The variables defined in the module.
	Variables map[string]*ConfigVariable `json:"variables,omitempty"`
}

// ConfigOutput defines an output as defined in configuration.
type ConfigOutput struct {
	// Indicates whether or not the output was marked as sensitive.
	Sensitive bool `json:"sensitive,omitempty"`

	// The defined value of the output.
	Expression *Expression `json:"expression,omitempty"`

	// The defined description of this output.
	Description string `json:"description,omitempty"`

	// The defined dependencies tied to this output.
	DependsOn []string `json:"depends_on,omitempty"`
}

// ConfigResource is the configuration representation of a resource.
type ConfigResource struct {
	// The address of the resource relative to the module that it is
	// in.
	Address string `json:"address,omitempty"`

	// The resource mode.
	Mode ResourceMode `json:"mode,omitempty"`

	// The type of resource, ie: "null_resource" in
	// "null_resource.foo".
	Type string `json:"type,omitempty"`

	// The name of the resource, ie: "foo" in "null_resource.foo".
	Name string `json:"name,omitempty"`

	// An opaque key representing the provider configuration this
	// module uses. Note that there are more than one circumstance that
	// this key will not match what is found in the ProviderConfigs
	// field in the root Config structure, and as such should not be
	// relied on for that purpose.
	ProviderConfigKey string `json:"provider_config_key,omitempty"`

	// The list of provisioner defined for this configuration. This
	// will be nil if no providers are defined.
	Provisioners []*ConfigProvisioner `json:"provisioners,omitempty"`

	// Any non-special configuration values in the resource, indexed by
	// key.
	Expressions map[string]*Expression `json:"expressions,omitempty"`

	// The resource's configuration schema version. With access to the
	// specific Terraform provider for this resource, this can be used
	// to determine the correct schema for the configuration data
	// supplied in Expressions.
	SchemaVersion uint64 `json:"schema_version"`

	// The expression data for the "count" value in the resource.
	CountExpression *Expression `json:"count_expression,omitempty"`

	// The expression data for the "for_each" value in the resource.
	ForEachExpression *Expression `json:"for_each_expression,omitempty"`
}

// ConfigVariable defines a variable as defined in configuration.
type ConfigVariable struct {
	// The defined default value of the variable.
	Default interface{} `json:"default,omitempty"`

	// The defined text description of the variable.
	Description string `json:"description,omitempty"`
}

// ConfigProvisioner describes a provisioner declared in a resource
// configuration.
type ConfigProvisioner struct {
	// The type of the provisioner, ie: "local-exec".
	Type string `json:"type,omitempty"`

	// Any non-special configuration values in the provisioner, indexed by
	// key.
	Expressions map[string]*Expression `json:"expressions,omitempty"`
}

// ModuleCall describes a declared "module" within a configuration.
// It also contains the data for the module itself.
type ModuleCall struct {
	// The contents of the "source" field.
	Source string `json:"source,omitempty"`

	// Any non-special configuration values in the module, indexed by
	// key.
	Expressions map[string]*Expression `json:"expressions,omitempty"`

	// The expression data for the "count" value in the module.
	CountExpression *Expression `json:"count_expression,omitempty"`

	// The expression data for the "for_each" value in the module.
	ForEachExpression *Expression `json:"for_each_expression,omitempty"`

	// The configuration data for the module itself.
	Module *ConfigModule `json:"module,omitempty"`

	// The version constraint for modules that come from the registry.
	VersionConstraint string `json:"version_constraint,omitempty"`
}

// Expression describes the format for an individual key in a
// Terraform configuration.
//
// This struct wraps ExpressionData to support custom JSON parsing.
type Expression struct {
	*ExpressionData
}

// ExpressionData describes the format for an individual key in a
// Terraform configuration.
type ExpressionData struct {
	// If the *entire* expression is a constant-defined value, this
	// will contain the Go representation of the expression's data.
	ConstantValue interface{} `json:"constant_value,omitempty"`

	// If any part of the expression contained values that were not
	// able to be resolved at parse-time, this will contain a list of
	// the referenced identifiers that caused the value to be unknown.
	References []string `json:"references,omitempty"`

	// A list of complex objects that were nested in this expression.
	// If this value is a nested block in configuration, sometimes
	// referred to as a "sub-resource", this field will contain those
	// values, and ConstantValue and References will be blank.
	NestedBlocks []map[string]*Expression `json:"-"`
}

// UnmarshalJSON implements json.Unmarshaler for Expression.
func (e *Expression) UnmarshalJSON(b []byte) error {
	result := new(ExpressionData)

	// Check to see if this is an array first. If it is, this is more
	// than likely a list of nested blocks.
	var rawNested []map[string]json.RawMessage
	if err := json.Unmarshal(b, &rawNested); err == nil {
		result.NestedBlocks, err = unmarshalExpressionBlocks(rawNested)
		if err != nil {
			return err
		}
	} else {
		// It's a non-nested expression block, parse normally
		if err := json.Unmarshal(b, &result); err != nil {
			return err
		}
	}

	e.ExpressionData = result
	return nil
}

func unmarshalExpressionBlocks(raw []map[string]json.RawMessage) ([]map[string]*Expression, error) {
	var result []map[string]*Expression

	for _, rawBlock := range raw {
		block := make(map[string]*Expression)
		for k, rawExpr := range rawBlock {
			var expr *Expression
			if err := json.Unmarshal(rawExpr, &expr); err != nil {
				return nil, err
			}

			block[k] = expr
		}

		result = append(result, block)
	}

	return result, nil
}

// MarshalJSON implements json.Marshaler for Expression.
func (e *Expression) MarshalJSON() ([]byte, error) {
	// Check for nested blocks, and marshal those instead if they exist.
	if len(e.ExpressionData.NestedBlocks) > 0 {
		return marshalExpressionBlocks(e.ExpressionData.NestedBlocks)
	}

	return json.Marshal(e.ExpressionData)
}

func marshalExpressionBlocks(nested []map[string]*Expression) ([]byte, error) {
	var rawNested []map[string]json.RawMessage
	for _, block := range nested {
		rawBlock := make(map[string]json.RawMessage)
		for k, expr := range block {
			raw, err := json.Marshal(expr)
			if err != nil {
				return nil, err
			}

			rawBlock[k] = raw
		}

		rawNested = append(rawNested, rawBlock)
	}

	return json.Marshal(rawNested)
}
