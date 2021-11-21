package week08

import (
	"context"
	"testing"
)

func TestEmployee_work(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name string
		e    Employee
		prod *Product
		err  error
	}{
		{
			name: "invalid employee",
			e:    Employee(0),
			prod: &Product{Materials: Materials{Wood: 2, Oil: 3}},
			err:  ErrInvalidEmployee(0),
		},
		{
			name: "valid employee",
			prod: &Product{Materials: Materials{Wood: 1, Oil: 1}},
			e:    Employee(1),
			err:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()

			manager := &Manager{}
			manager.Warehouse = &Warehouse{}

			go tt.e.work(ctx, manager)

			err := manager.Assign(tt.prod)
			if err != nil {
				t.Fatal(err)
			}

			select {
			case err := <-manager.Errors():
				if err != nil {
					if err.Error() != tt.err.Error() {
						t.Fatalf("expected %v, got %v", tt.err, err)
					}
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
