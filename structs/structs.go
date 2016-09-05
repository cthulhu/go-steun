package structs

import (
	"errors"
	"fmt"
	"reflect"
)

func Set(structObj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(structObj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}

	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}

	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		invalidTypeError := errors.New("Provided value type didn't match obj field type")
		return invalidTypeError
	}

	structFieldValue.Set(val)
	return nil
}

func Get(structObj interface{}, name string) (interface{}, error) {
	structValue := reflect.ValueOf(structObj).Elem()
	structFieldValue := structValue.FieldByName(name)

	if !structFieldValue.IsValid() {
		return nil, fmt.Errorf("No such field: %s in obj", name)
	}
	return structFieldValue.Interface(), nil
}
