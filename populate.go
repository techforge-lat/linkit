package dependor

import (
	"fmt"
	"reflect"
)

// SetAuxiliarDependencies populates auxiliar dependencies of every dependency in the container
func SetAuxiliarDependencies() error {
	if err := populate(container); err != nil {
		return err
	}

	return nil
}

// populate exists to be tested easily by receiving the dependency container
func populate(container dependencyContainer) error {
	for _, v := range container {
		for fieldName, dependencyName := range v.dependsOn {
			structValue := reflect.ValueOf(v.value)
			if !isPointer(structValue) {
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

// setAuxDependency sets an auxiliary dependency to a struct's field
func setAuxDependency(structValue reflect.Value, fieldName string, dependency any) error {
	field, err := getFieldByName(fieldName, structValue)
	if err != nil {
		return fmt.Errorf("setAuxDependency(): %v", err)
	}

	field.Set(reflect.ValueOf(dependency))

	return nil
}

// getFieldByName gets a struct's field by providing it's name
func getFieldByName(fieldName string, structValue reflect.Value) (reflect.Value, error) {
	fieldByName, _ := structValue.Type().FieldByName(fieldName)
	if len(fieldByName.Index) == 0 {
		return reflect.Value{}, fmt.Errorf("getFieldByName(): field %s not found in struct %s", fieldName, structValue.Type().String())
	}

	return structValue.Field(fieldByName.Index[0]), nil
}

// isPointer validates if value is a pointer
func isPointer(value reflect.Value) bool {
	return value.Kind() == reflect.Pointer
}
