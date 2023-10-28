package dependor

import "testing"

type TableRecord struct {
	name      string
	container dependencyContainer
	wantErr   bool
	err       string
}

type TableRecords []TableRecord

func TestPopulate(t *testing.T) {
	tests := TableRecords{
		getWithOneAuxiliarDep(),
		getWithPointerErr(),
		getWithoutStruct(),
		getWithMissingAuxDependency(),
		getWithInvalidAuxDependency(),
		getWithMissingField(),
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

func getWithOneAuxiliarDep() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	user := &User{}
	set[*User](c, Name(User{}), user, map[string]string{
		"Role": Name(Role{}),
	})

	role := &Role{}
	set[*Role](c, Name(Role{}), role, map[string]string{
		"Permission": Name(Permission{}),
	})

	permission := &Permission{}
	set[*Permission](c, Name(Permission{}), permission, map[string]string{})

	return TableRecord{
		name:      "one auxiliar dep",
		container: c,
		wantErr:   false,
	}
}

func getWithPointerErr() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	user := User{}
	set[User](c, Name(User{}), user, map[string]string{
		"Role": Name(Role{}),
	})

	role := &Role{}
	set[*Role](c, Name(Role{}), role, map[string]string{
		"Permission": Name(Permission{}),
	})

	permission := &Permission{}
	set[*Permission](c, Name(Permission{}), permission, map[string]string{})

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "dependor.Populate(): cound not assign auxiliary dependencies to User of type struct, you must pass a struct pointer",
	}
}

func getWithoutStruct() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	user := 0
	set[int](c, Name(User{}), user, map[string]string{
		"Role": Name(Role{}),
	})

	role := &Role{}
	set[*Role](c, Name(Role{}), role, map[string]string{
		"Permission": Name(Permission{}),
	})

	permission := &Permission{}
	set[*Permission](c, Name(Permission{}), permission, map[string]string{})

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "dependor.Populate(): cound not assign auxiliary dependencies to int of type int, you must pass a struct pointer",
	}
}

func getWithMissingAuxDependency() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	user := &User{}
	set[*User](c, Name(User{}), user, map[string]string{
		"Role": Name(Role{}),
	})

	role := &Role{}
	set[*Role](c, Name(Role{}), role, map[string]string{
		"Permission": Name(Permission{}),
	})

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "dependor.Populate(): missing auxiliary dependency with name dependor.Permission for dependor.Role",
	}
}

func getWithInvalidAuxDependency() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	user := &User{}
	set[*User](c, Name(User{}), user, map[string]string{
		"Role": Name(Role{}),
	})

	role := &Role{}
	set[*Role](c, Name(Role{}), role, map[string]string{
		"Permission": Name(Permission{}),
	})

	set[int](c, Name(Permission{}), 10, map[string]string{})

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "dependor.Populate(): setAuxDependency(): cannot set auxiliary dependency with name dependor.Permission to field Permission of struct dependor.Role",
	}
}

func getWithMissingField() TableRecord {
	var c dependencyContainer = make(dependencyContainer)
	set[*User](c, Name(User{}), &User{}, map[string]string{
		"Role": Name(Role{}),
	})

	set[*Role](c, Name(Role{}), &Role{}, map[string]string{
		"MissingField": Name(Permission{}),
	})

	set[*Permission](c, Name(Permission{}), &Permission{}, map[string]string{})

	return TableRecord{
		name:      "pointer error",
		container: c,
		wantErr:   true,
		err:       "dependor.Populate(): setAuxDependency(): getFieldByName(): field MissingField not found in struct dependor.Role",
	}
}
