package tfjson

// Config represents the complete configuration source
type Config struct {
	// A map of all provider instances across all modules in the
	// configuration, indexed in the format NAME.ALIAS. Default
	// providers will not have an alias.
	ProviderConfigs map[string]ProviderConfig `json:"provider_config,omitempty"`

	// The root module in the configuration. Any child modules descend
	// off of here.
	RootModule *ConfigModule `json:"root_module,omitempty"`
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
	Expressions map[string]Expression `json:"expressions,omitempty"`
}

// ConfigModule describes a module in Terraform configuration.
type ConfigModule struct {
	// The outputs defined in the module.
	Outputs map[string]ConfigOutput `json:"outputs,omitempty"`

	// The resources defined in the module.
	Resources []ConfigResource `json:"resources,omitempty"`

	// Any "module" stanzas within the specific module.
	ModuleCalls []ModuleCall `json:"module_calls,omitempty"`
}

// ConfigOutput defines an output as defined in configuration.
type ConfigOutput struct {
	// Indicates whether or not the output was marked as sensitive.
	Sensitive bool `json:"sensitive,omitempty"`

	// The defined value of the output.
	Expression *Expression `json:"expression,omitempty"`
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

	// The provider address used for this resource. This address should
	// be able to be referenced in the ProviderConfig key in the
	// root-level Config structure.
	ProviderConfigKey string `json:"provider_config_key,omitempty"`

	// The list of provisioner defined for this configuration. This
	// will be nil if no providers are defined.
	Provisioners []ConfigProvisioner `json:"provisioners,omitempty"`

	// Any non-special configuration values in the resource, indexed by
	// key.
	Expressions map[string]Expression `json:"expressions,omitempty"`

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

// ConfigProvisioner describes a provisioner declared in a resource
// configuration.
type ConfigProvisioner struct {
	// The type of the provisioner, ie: "local-exec".
	Type string `json:"type,omitempty"`

	// Any non-special configuration values in the provisioner, indexed by
	// key.
	Expressions map[string]Expression `json:"expressions,omitempty"`
}

// ModuleCall describes a declared "module" within a configuration.
// It also contains the data for the module itself.
type ModuleCall struct {
	// The resolved contents of the "source" field.
	ResolvedSource string `json:"resolved_source,omitempty"`

	// Any non-special configuration values in the module, indexed by
	// key.
	Expressions map[string]Expression `json:"expressions,omitempty"`

	// The expression data for the "count" value in the module.
	CountExpression *Expression `json:"count_expression,omitempty"`

	// The expression data for the "for_each" value in the module.
	ForEachExpression *Expression `json:"for_each_expression,omitempty"`

	// The configuration data for the module itself.
	Module *ConfigModule `json:"module,omitempty"`
}
