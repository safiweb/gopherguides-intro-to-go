package week06

import "testing"

func TestEmployee_IsValid(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name     string
		employee Employee
		err      error
	}{
		{
			name:     "invalid employee",
			employee: Employee(-1),
			err:      ErrInvalidEmployee(Employee(-1)),
		},
		{
			name:     "valid employee",
			employee: Employee(1),
			err:      nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.employee.IsValid()
			if got != nil {
				if got.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, got)
				}
			}
		})
	}
}
