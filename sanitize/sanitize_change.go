package sanitize

import (
	tfjson "github.com/spacelift-io/terraform-json"
)

// SanitizeChange traverses a Change and replaces all values at
// the particular locations marked by BeforeSensitive AfterSensitive
// with the value supplied as replaceWith.
//
// A new change is issued.
func SanitizeChange(old *tfjson.Change, replaceWith interface{}) (*tfjson.Change, error) {
	return SanitizeChangeDynamic(old, NewValueSanitizer(replaceWith))
}

// SanitizeChangeDynamic traverses a Change and replaces all values at
// the particular locations marked by BeforeSensitive AfterSensitive
// by using the sanitization function passed in.
//
// A new change is issued.
func SanitizeChangeDynamic(old *tfjson.Change, sanitizer Sanitizer) (*tfjson.Change, error) {
	result, err := copyChange(old)
	if err != nil {
		return nil, err
	}

	result.Before = sanitizeChangeValueDynamic(result.Before, result.BeforeSensitive, sanitizer)
	result.After = sanitizeChangeValueDynamic(result.After, result.AfterSensitive, sanitizer)

	return result, nil
}

func sanitizeChangeValue(old, sensitive, replaceWith interface{}) interface{} {
	return sanitizeChangeValueDynamic(old, sensitive, NewValueSanitizer(replaceWith))
}

func sanitizeChangeValueDynamic(old, sensitive interface{}, replaceWith Sanitizer) interface{} {
	// Only expect deep types that we would normally see in JSON, so
	// arrays and objects.
	switch x := old.(type) {
	case []interface{}:
		if filterSlice, ok := sensitive.([]interface{}); ok {
			for i := range filterSlice {
				if i >= len(x) {
					break
				}

				x[i] = sanitizeChangeValueDynamic(x[i], filterSlice[i], replaceWith)
			}
		}
	case map[string]interface{}:
		if filterMap, ok := sensitive.(map[string]interface{}); ok {
			for filterKey := range filterMap {
				if value, ok := x[filterKey]; ok {
					x[filterKey] = sanitizeChangeValueDynamic(value, filterMap[filterKey], replaceWith)
				}
			}
		}
	}

	if shouldFilter, ok := sensitive.(bool); ok && shouldFilter {
		replacement := replaceWith(old)
		return replacement
	}

	return old
}
