package data

import (
	"errors"
	"reflect"
)

// ValidateRequiredFields Generic function to check if required fields are non-zero (i.e., not empty/default)
func ValidateRequiredFields[T any](input T, requiredFields []string) error {
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for _, field := range requiredFields {
		f := v.FieldByName(field)
		if !f.IsValid() {
			return errors.New("missing field: " + field)
		}

		// Check if zero value
		if reflect.DeepEqual(f.Interface(), reflect.Zero(f.Type()).Interface()) {
			return errors.New("field is required: " + field)
		}
	}

	return nil
}
