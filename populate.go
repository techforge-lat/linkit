package dependor

import (
	"fmt"
	"reflect"
)

// Populate populates auxiliar dependencies of every dependency in the container
func Populate() error {
	for _, v := range container {
		for _, depName := range v.dependsOn {
			depType := reflect.TypeOf(v.value)
			if !isPointer(depType.Kind()) {
				return fmt.Errorf("dependor: cound not assign dependencies to struct %s, you must pass a struct pointer", depType.String())
			}
			depType = depType.Elem()

			for i := 0; i < depType.NumField(); i++ {
				if !isAuxDependency(depType, i) {
					continue
				}

				depValue := reflect.ValueOf(v.value).Elem()
				dependencyField := depValue.Field(i)
				dependencyField.Set(reflect.ValueOf(container[depName].value))
			}
		}
	}

	return nil
}
