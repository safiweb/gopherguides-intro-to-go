package week07

import (
	"context"
	"testing"
)

func TestEmployee_work(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name     string
		employee Employee
		err      error
	}{
		{
			name:     "bad employee",
			employee: Employee(0),
			err:      ErrInvalidEmployee(0),
		},
		{
			name:     "good employee",
			employee: Employee(1),
			err:      nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			manager := NewManager()
			//defer manager.Stop()

			prod := Product{Quantity: 1}

			go tt.employee.work(ctx, manager)

			err := manager.Assign(&prod)
			if err != nil {
				t.Fatal(err)
			}

			select {
			case err := <-manager.Errors():
				if err.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
				return
			case cp := <-manager.Completed():
				if check := cp.IsValid(); check != nil {
					t.Fatalf("expected %v, got %v", tt.err, cp)
				}
			}

		})
	}
}
