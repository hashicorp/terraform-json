package sanitize

import (
	"reflect"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/mitchellh/copystructure"
)

// copyStructureCopy is an internal function that wraps copystructure.Copy with
// a shallow copier for unknown values.
//
// Performing the shallow copy of the unknown values is important
// here, as unknown values are parsed in with the main terraform-json
// package as singletons, and must continue to be comparable.
func copyStructureCopy(v interface{}) (interface{}, error) {
	c := &copystructure.Config{
		ShallowCopiers: map[reflect.Type]struct{}{
			reflect.TypeOf(tfjson.UnknownConstantValue): struct{}{},
		},
	}

	return c.Copy(v)
}

// copyChange copies a Change value and returns the copy.
func copyChange(old *tfjson.Change) (*tfjson.Change, error) {
	c, err := copyStructureCopy(old)
	if err != nil {
		return nil, err
	}

	return c.(*tfjson.Change), nil
}

// copyPlan copies a Plan value and returns the copy.
func copyPlan(old *tfjson.Plan) (*tfjson.Plan, error) {
	c, err := copyStructureCopy(old)
	if err != nil {
		return nil, err
	}

	return c.(*tfjson.Plan), nil
}

// copyPlanVariable copies a PlanVariable value and returns the copy.
func copyPlanVariable(old *tfjson.PlanVariable) (*tfjson.PlanVariable, error) {
	c, err := copyStructureCopy(old)
	if err != nil {
		return nil, err
	}

	return c.(*tfjson.PlanVariable), nil
}

// copyStateResource copies a StateResource value and returns the copy.
func copyStateResource(old *tfjson.StateResource) (*tfjson.StateResource, error) {
	c, err := copyStructureCopy(old)
	if err != nil {
		return nil, err
	}

	return c.(*tfjson.StateResource), nil
}

// copyStateOutput copies a StateOutput value and returns the copy.
func copyStateOutputs(old map[string]*tfjson.StateOutput) (map[string]*tfjson.StateOutput, error) {
	c, err := copystructure.Copy(old)
	if err != nil {
		return nil, err
	}

	return c.(map[string]*tfjson.StateOutput), nil
}
