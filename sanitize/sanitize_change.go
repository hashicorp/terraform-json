package sanitize

import (
	tfjson "github.com/hashicorp/terraform-json"
)

// SanitizeChange traverses a Change and replaces all values at
// the particular locations marked by BeforeSensitive AfterSensitive
// with the value supplied as replaceWith.
//
// A new change is issued.
func SanitizeChange(old *tfjson.Change, replaceWith interface{}) (*tfjson.Change, error) {
	result, err := copyChange(old)
	if err != nil {
		return nil, err
	}

	result.Before = sanitizeChangeValue(result.Before, result.BeforeSensitive, replaceWith)
	result.After = sanitizeChangeValue(result.After, result.AfterSensitive, replaceWith)

	return result, nil
}

func sanitizeChangeValue(old, sensitive, replaceWith interface{}) interface{} {
	// Only expect deep types that we would normally see in JSON, so
	// arrays and objects.
	switch x := old.(type) {
	case []interface{}:
		if filterSlice, ok := sensitive.([]interface{}); ok {
			for i := range filterSlice {
				if i >= len(x) {
					break
				}

				x[i] = sanitizeChangeValue(x[i], filterSlice[i], replaceWith)
			}
		}
	case map[string]interface{}:
		if filterMap, ok := sensitive.(map[string]interface{}); ok {
			for filterKey := range filterMap {
				if value, ok := x[filterKey]; ok {
					x[filterKey] = sanitizeChangeValue(value, filterMap[filterKey], replaceWith)
				}
			}
		}
	}

	if shouldFilter, ok := sensitive.(bool); ok && shouldFilter {
		return replaceWith
	}

	return old
}
