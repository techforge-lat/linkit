package dependor

import "testing"

func TestPopulate(t *testing.T) {
	tests := []struct {
		name      string
		container dependencyContainer
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := populate(tt.container); (err != nil) != tt.wantErr {
				t.Errorf("Populate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
