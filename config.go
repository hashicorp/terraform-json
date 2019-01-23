package tfjson

// Config represents the complete configuration source
type Config struct {
	// A map of all provider instances across all modules in the
	// configuration, indexed in the format NAME.ALIAS. Default
	// providers will not have an alias.
	ProviderConfigs map[string]ProviderConfig `json:"provider_config,omitempty"`

	// The root module in the configuration. Any child modules descend
	// off of here.
	// RootModule ConfigModule `json:"root_module,omitempty"`
}

// ProviderConfig describes a provider configuration instance.
type ProviderConfig struct {
	// The name of the provider, ie: "aws".
	Name string `json:"name,omitempty"`

	// The alias of the provider, ie: "us-east-1".
	Alias string `json:"alias,omitempty"`

	// The address of the module the provider is declared in.
	ModuleAddress string `json:"module_address,omitempty"`

	// All configuration expressions in the module, by key.
	Expressions map[string]Expression `json:"expressions,omitempty"`
}
