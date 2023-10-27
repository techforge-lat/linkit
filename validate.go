package dependor

import (
	"fmt"
	"reflect"
)

// Validate validates that every auxiliar dependency is set
func Validate() error {
	for _, v := range container {
		depType := reflect.TypeOf(v.value)
		if isPointer(depType.Kind()) {
			depType = depType.Elem()
		}

		for i := 0; i < depType.NumField(); i++ {
			if !isAuxDependency(depType, i) {
				continue
			}

			depValue := reflect.ValueOf(v.value)
			if isPointer(depType.Kind()) {
				depValue = depValue.Elem()
			}

			field := depValue.Field(i)
			if !isValueSet(field) {
				return fmt.Errorf("dependor: missing %s dependency in struct %s", field.Type().Name(), depType.String())
			}
		}
	}

	return nil

}

func isAuxDependency(t reflect.Type, fieldIndex int) bool {
	_, exist := t.Field(fieldIndex).Tag.Lookup("dependor")
	return exist
}

func isPointer(k reflect.Kind) bool {
	return k == reflect.Pointer
}

func isValueSet(field reflect.Value) bool {
	return !field.IsNil() && !field.IsZero()
}
