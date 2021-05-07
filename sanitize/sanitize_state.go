package sanitize

import (
	"fmt"

	tfjson "github.com/hashicorp/terraform-json"
)

type SanitizeStateModuleChangeMode string

const (
	SanitizeStateModuleChangeModeBefore SanitizeStateModuleChangeMode = "before_sensitive"
	SanitizeStateModuleChangeModeAfter  SanitizeStateModuleChangeMode = "after_sensitive"
)

// SanitizeStateModule traverses a StateModule, consulting the
// supplied ResourceChange set for resources to determine whether or
// not particular values should be obfuscated.
//
// Use mode to supply the SanitizeStateModuleChangeMode that
// represents what sensitive field should be consulted to determine
// whether or not the value should be obfuscated:
//
// * SanitizeStateModuleChangeModeBefore for before_sensitive
// * SanitizeStateModuleChangeModeAfter for after_sensitive
//
// Sensitive values are replaced with the supplied replaceWith value.
// A new state module tree is issued.
func SanitizeStateModule(
	old *tfjson.StateModule,
	resourceChanges []*tfjson.ResourceChange,
	mode SanitizeStateModuleChangeMode,
	replaceWith interface{},
) (*tfjson.StateModule, error) {
	result := &tfjson.StateModule{
		Resources:    make([]*tfjson.StateResource, len(old.Resources)),
		Address:      old.Address,
		ChildModules: make([]*tfjson.StateModule, len(old.ChildModules)),
	}

	for i := range old.Resources {
		var err error
		result.Resources[i], err = sanitizeStateResource(
			old.Resources[i],
			findResourceChange(resourceChanges, old.Resources[i].Address),
			mode,
			replaceWith,
		)
		if err != nil {
			return nil, err
		}
	}

	for i := range old.ChildModules {
		var err error
		result.ChildModules[i], err = SanitizeStateModule(
			old.ChildModules[i],
			resourceChanges,
			mode,
			replaceWith,
		)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

func sanitizeStateResource(
	old *tfjson.StateResource,
	rc *tfjson.ResourceChange,
	mode SanitizeStateModuleChangeMode,
	replaceWith interface{},
) (*tfjson.StateResource, error) {
	result, err := copyStateResource(old)
	if err != nil {
		return nil, err
	}

	if rc == nil {
		return result, nil
	}

	var sensitive interface{}
	switch mode {
	case SanitizeStateModuleChangeModeBefore:
		sensitive = rc.Change.BeforeSensitive

	case SanitizeStateModuleChangeModeAfter:
		sensitive = rc.Change.AfterSensitive

	default:
		panic(fmt.Sprintf("invalid change mode %q", mode))
	}

	// We can re-use sanitizeChangeValue here to do the sanitization.
	result.AttributeValues = sanitizeChangeValue(result.AttributeValues, sensitive, replaceWith).(map[string]interface{})
	return result, nil
}

func findResourceChange(resourceChanges []*tfjson.ResourceChange, addr string) *tfjson.ResourceChange {
	// Linear search here, unfortunately :P
	for _, rc := range resourceChanges {
		if rc.Address == addr {
			return rc
		}
	}

	return nil
}

// SanitizeStateOutputs scans the supplied map of StateOutputs and
// replaces any values of outputs marked as Sensitive with the value
// supplied in replaceWith.
//
// A new copy of StateOutputs is returned.
func SanitizeStateOutputs(old map[string]*tfjson.StateOutput, replaceWith interface{}) (map[string]*tfjson.StateOutput, error) {
	result, err := copyStateOutputs(old)
	if err != nil {
		return nil, err
	}

	for k := range result {
		if result[k].Sensitive {
			result[k].Value = replaceWith
		}
	}

	return result, nil
}
