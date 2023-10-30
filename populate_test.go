package linkit

import (
	"testing"
)

type TableRecord struct {
	name      string
	container dependencyContainer
	wantErr   bool
	err       string
}

type TableRecords []TableRecord

func TestPopulate(t *testing.T) {
	tests := TableRecords{
		getWithOneAuxiliaryDeep(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := populate(tt.container)
			if (err != nil) != tt.wantErr {
				t.Errorf("populate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && tt.err != err.Error() {
				t.Errorf("populate() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

func TestPopulate_WithErr(t *testing.T) {
	tests := TableRecords{
		getWithPointerErr(),
		getWithoutStruct(),
		getWithMissingAuxDependency(),
		getWithInvalidAuxDependency(),
		getWithMissingField(),
		// withNotImplementInterface
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := populate(tt.container)
			if (err != nil) != tt.wantErr {
				t.Errorf("populate() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr && tt.err != err.Error() {
				t.Errorf("populate() error = %v, wantErr %v", err, tt.err)
			}
		})
	}
}

type RoleUseCase interface {
	GetRole() string
}

type PermissionUseCase interface {
	GetPermission() string
}

type User struct {
	Role RoleUseCase
}

type Role struct {
	Permission PermissionUseCase
}

func (r Role) GetRole() string {
	return "role"
}

type Permission struct{}

func (p Permission) GetPermission() string {
	return "permission"
}

func getWithOneAuxiliaryDeep() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	set[*User](c, WithAuxiliaryDependencies(
		map[string]string{
			"Role": Name(Role{}),
		},
	))

	set[*Role](c, WithAuxiliaryDependencies(
		map[string]string{
			"Permission": Name(Permission{}),
		},
	))

	set[*Permission](c)

	return TableRecord{
		name:      "one auxiliary dep",
		container: c,
		wantErr:   false,
	}
}

func getWithPointerErr() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	set[User](c, WithAuxiliaryDependencies(
		map[string]string{
			"Role": Name(Role{}),
		},
	))

	set[*Role](c, WithAuxiliaryDependencies(
		map[string]string{
			"Permission": Name(Permission{}),
		},
	))

	set[*Permission](c)

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "linkit.populate(): could not assign auxiliary dependencies to linkit.User of type struct, you must pass a struct pointer",
	}
}

func getWithoutStruct() TableRecord {
	var c dependencyContainer = make(dependencyContainer)

	set[int](
		c,
		WithName(Name(User{})),
		WithValue(0),
		WithAuxiliaryDependencies(
			map[string]string{
				"Role": Name(Role{}),
			},
		))

	set[*Role](c, WithAuxiliaryDependencies(
		map[string]string{
			"Permission": Name(Permission{}),
		},
	))

	set[*Permission](c)

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "linkit.populate(): could not assign auxiliary dependencies to linkit.User of type int, you must pass a struct pointer",
	}
}

func getWithMissingAuxDependency() TableRecord {
	var c dependencyContainer = make(dependencyContainer)

	set[*User](c, WithAuxiliaryDependencies(
		map[string]string{
			"Role": Name(Role{}),
		},
	))

	set[*Role](c, WithAuxiliaryDependencies(
		map[string]string{
			"Permission": Name(Permission{}),
		},
	))

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "linkit.populate(): missing auxiliary dependency with name linkit.Permission for linkit.Role",
	}
}

func getWithInvalidAuxDependency() TableRecord {
	var c dependencyContainer = make(dependencyContainer)

	set[*User](c, WithAuxiliaryDependencies(
		map[string]string{
			"Role": Name(Role{}),
		},
	))

	set[*Role](c, WithAuxiliaryDependencies(
		map[string]string{
			"Permission": Name(Permission{}),
		},
	))

	set[int](
		c,
		WithName(Name(Permission{})),
		WithValue(10),
	)

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "linkit.populate(): setAuxDependency(): cannot set auxiliary dependency with name linkit.Permission to field Permission of struct linkit.Role",
	}
}

func getWithMissingField() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	set[*User](c, WithAuxiliaryDependencies(
		map[string]string{
			"Role": Name(Role{}),
		},
	))

	set[*Role](c, WithAuxiliaryDependencies(
		map[string]string{
			"MissingField": Name(Permission{}),
		},
	))

	set[*Permission](c)

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "linkit.populate(): setAuxDependency(): getFieldByName(): field MissingField not found in struct linkit.Role",
	}
}
