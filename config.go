package tfjson

// Config represents the complete configuration source
type Config struct {
	// A map of all provider instances across all modules in the
	// configuration, indexed in the format NAME.ALIAS. Default
	// providers will not have an alias.
	ProviderConfigs map[string]ProviderConfig `json:"provider_config,omitempty"`

	// The root module in the configuration. Any child modules descend
	// off of here.
	RootModule ConfigModule `json:"root_module,omitempty"`
}

// ProviderConfig describes a provider configuration instance.
type ProviderConfig struct {
	Name          string                 `json:"name,omitempty"`
	Alias         string                 `json:"alias,omitempty"`
	ModuleAddress string                 `json:"module_address,omitempty"`
	Expressions   map[string]interface{} `json:"expressions,omitempty"`
}
