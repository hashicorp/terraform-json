package sanitize

type Sanitizer func(old interface{}) interface{}

func NewValueSanitizer(value interface{}) Sanitizer {
	return func(old interface{}) interface{} {
		return value
	}
}
