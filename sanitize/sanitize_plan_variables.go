package sanitize

import (
	tfjson "github.com/spacelift-io/terraform-json"
)

// SanitizePlanVariables traverses a map of PlanVariable and replaces
// any sensitive values with the value supplied in replaceWith.
// configs should be the map of ConfigVariables from the root module
// (so Plan.Config.RootModule.Variables).
//
// A new copy of the PlanVariable map is returned.
func SanitizePlanVariables(
	old map[string]*tfjson.PlanVariable,
	configs map[string]*tfjson.ConfigVariable,
	replaceWith interface{},
) (map[string]*tfjson.PlanVariable, error) {
	return SanitizePlanVariablesDynamic(old, configs, NewValueSanitizer(replaceWith))
}

// SanitizePlanVariablesDynamic traverses a map of PlanVariable and replaces
// any sensitive values by using the sanitizer provideds.
// configs should be the map of ConfigVariables from the root module
// (so Plan.Config.RootModule.Variables).
//
// A new copy of the PlanVariable map is returned.
func SanitizePlanVariablesDynamic(
	old map[string]*tfjson.PlanVariable,
	configs map[string]*tfjson.ConfigVariable,
	sanitizer Sanitizer,
) (map[string]*tfjson.PlanVariable, error) {
	result := make(map[string]*tfjson.PlanVariable, len(old))
	for k := range old {
		v, err := sanitizeVariable(old[k], configs[k], sanitizer)
		if err != nil {
			return nil, err
		}

		result[k] = v
	}

	return result, nil
}

func sanitizeVariable(
	old *tfjson.PlanVariable,
	config *tfjson.ConfigVariable,
	sanitizer Sanitizer,
) (*tfjson.PlanVariable, error) {
	result, err := copyPlanVariable(old)
	if err != nil {
		return nil, err
	}

	if config != nil && config.Sensitive {
		result.Value = sanitizer(result.Value)
	}

	return result, nil
}
