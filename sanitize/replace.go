package sanitize

func applyReplacement(value interface{}, replaceWith interface{}) interface{} {
	switch x := replaceWith.(type) {
	case func(value interface{}) interface{}:
		return x(value)
	default:
		return x
	}
}
