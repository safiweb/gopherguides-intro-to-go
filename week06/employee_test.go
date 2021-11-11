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

func TestEmployee_work(t *testing.T) {
	t.Parallel()

	/*manager := NewManager()
	defer manager.Stop()

	e := Employee(1)
	p := Product{Quantity: 1}

	go e.work(manager)

	err := manager.Assign(&p)
	if err != nil {
		t.Fatal(err)
	}

	select {
	case err := <-manager.Errors():
		t.Fatal(err)
		return
	case cp := <-manager.Completed():
		if check := cp.IsValid(); check != nil {
			t.Fatal(cp)
		}
	}*/

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

			manager := NewManager()
			defer manager.Stop()

			prod := Product{Quantity: 1}

			go tt.employee.work(manager)

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
