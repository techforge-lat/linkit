package dependor

import (
	"fmt"
	"reflect"
)

// Populate populates auxiliar dependencies of every dependency in the container
func Populate() error {
	for _, v := range container {
		for fieldName, dependencyName := range v.dependsOn {
			structValue := reflect.ValueOf(v.value)
			if !isPointer(structValue.Kind()) {
				return fmt.Errorf("dependor.Populate(): cound not assign dependencies to struct %s, you must pass a struct pointer", structValue.String())
			}
			structValue = structValue.Elem()

			if err := setAuxDependency(structValue, fieldName, container[dependencyName].value); err != nil {
				return fmt.Errorf("dependor.Populate(): %v", err)
			}
		}
	}

	return nil
}

func setAuxDependency(structValue reflect.Value, fieldName string, dependency any) error {
	field, err := getFieldByName(fieldName, structValue)
	if err != nil {
		return fmt.Errorf("setAuxDependency(): %v", err)
	}

	field.Set(reflect.ValueOf(dependency))

	return nil
}

func getFieldByName(fieldName string, structValue reflect.Value) (reflect.Value, error) {
	fieldByName, _ := structValue.Type().FieldByName(fieldName)
	if len(fieldByName.Index) == 0 {
		return reflect.Value{}, fmt.Errorf("getFieldByName(): field %s not found in struct %s", fieldName, structValue.Type().String())
	}

	return structValue.Field(fieldByName.Index[0]), nil
}

func isPointer(k reflect.Kind) bool {
	return k == reflect.Pointer
}
