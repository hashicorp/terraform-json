package sanitize

import (
	"errors"

	tfjson "github.com/hashicorp/terraform-json"
)

const DefaultSensitiveValue = "REDACTED_SENSITIVE"

var NilPlanError = errors.New("nil plan supplied")

// SanitizePlan sanitizes the entirety of a Plan, replacing sensitive
// values with the default value in DefaultSensitiveValue.
//
// See SanitizePlanWithValue for full detail on the where replacement
// takes place.
func SanitizePlan(old *tfjson.Plan) (*tfjson.Plan, error) {
	return SanitizePlanWithValue(old, DefaultSensitiveValue)
}

// SanitizePlanWithValue sanitizes the entirety of a Plan to the best
// of its ability, depending on the provided metadata on sensitive
// values. These are found in:
//
// * ResourceChanges: Sanitized based on BeforeSensitive and
// AfterSensitive fields.
//
// * Variables: Based on variable config data found in the root
// module of the Config.
//
// * PlannedValues: Sanitized based on the values found in
// AfterSensitive in ResourceChanges. Outputs are sanitized
// according to the appropriate sensitivity flags provided for the
// output.
//
// * PriorState: Sanitized based on the values found in
// BeforeSensitive in ResourceChanges. Outputs are sanitized according
// to the appropriate sensitivity flags provided for the output.
//
// * OutputChanges: Sanitized based on the values found in
// BeforeSensitive and AfterSensitive. This generally means that
// any sensitive output will have OutputChange fully obfuscated as
// the BeforeSensitive and AfterSensitive in outputs are opaquely the
// same.
//
// Sensitive values are replaced with the value supplied with
// replaceWith. A copy of the Plan is returned.
func SanitizePlanWithValue(old *tfjson.Plan, replaceWith interface{}) (*tfjson.Plan, error) {
	if old == nil {
		return nil, NilPlanError
	}

	result, err := copyPlan(old)
	if err != nil {
		return nil, err
	}

	// Sanitize ResourceChanges
	for i := range result.ResourceChanges {
		result.ResourceChanges[i].Change, err = SanitizeChange(result.ResourceChanges[i].Change, replaceWith)
		if err != nil {
			return nil, err
		}
	}

	// Sanitize Variables
	result.Variables, err = SanitizePlanVariables(result.Variables, result.Config.RootModule.Variables, replaceWith)
	if err != nil {
		return nil, err
	}

	// Sanitize PlannedValues
	result.PlannedValues.RootModule, err = SanitizeStateModule(
		result.PlannedValues.RootModule,
		result.ResourceChanges,
		SanitizeStateModuleChangeModeAfter,
		replaceWith)
	if err != nil {
		return nil, err
	}

	result.PlannedValues.Outputs, err = SanitizeStateOutputs(result.PlannedValues.Outputs, replaceWith)
	if err != nil {
		return nil, err
	}

	// Sanitize PriorState
	if result.PriorState != nil {
		result.PriorState.Values.RootModule, err = SanitizeStateModule(
			result.PriorState.Values.RootModule,
			result.ResourceChanges,
			SanitizeStateModuleChangeModeBefore,
			replaceWith)
		if err != nil {
			return nil, err
		}

		result.PriorState.Values.Outputs, err = SanitizeStateOutputs(result.PriorState.Values.Outputs, replaceWith)
		if err != nil {
			return nil, err
		}
	}

	// Sanitize OutputChanges
	for k := range result.OutputChanges {
		result.OutputChanges[k], err = SanitizeChange(result.OutputChanges[k], replaceWith)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
