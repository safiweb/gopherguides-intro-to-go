package week06

import (
	"fmt"
	"testing"
)

func TestManager_Start(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name  string
		count int
		err   error
	}{
		{
			name:  "invalid employee count",
			count: 0,
			err:   ErrInvalidEmployeeCount(0),
		},
		{
			name:  "valid employee count",
			count: 1,
			err:   nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			//create new manager
			manager := NewManager()
			defer manager.Stop()

			err := manager.Start(tt.count)
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			}

		})
	}
}

func TestManager_Assign(t *testing.T) {
	t.Parallel()

	fake := Product{}

	ready := Product{Quantity: 1}
	ready.Build(0)

	full := Product{Quantity: 1}
	full.Build(1)

	testcases := []struct {
		name     string
		manager  *Manager
		products *Product
		err      error
	}{
		{
			name:     "manager working, invalid products quantity",
			products: &fake,
			err:      ErrInvalidQuantity(0),
		},
		{
			name:     "manager working, invalid employee",
			products: &ready,
			err:      ErrInvalidEmployee(Employee(0)),
		},
		{
			name:     "manager working, valid products",
			products: &full,
			err:      nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			manager := NewManager()
			defer manager.Stop()

			err := manager.Start(1)
			if err != nil {
				t.Fatal(err)
			}

			go func() {
				got := manager.Assign(tt.products)
				if got != nil {
					manager.Errors() <- got
				}
			}()

			select {
			case err := <-manager.Errors():
				if err.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			case <-manager.Jobs():
				manager.Stop()
				return
			}

			manager.Stop()
		})
	}
}

func TestManager_Assign_ManagerStop(t *testing.T) {
	t.Parallel()

	//Create new manager & stop the manager from further work
	manager := NewManager()
	manager.Stop()

	fake := Product{}
	exp := ErrManagerStopped{}

	got := manager.Assign(&fake)
	if got != nil {
		if got.Error() != exp.Error() {
			t.Fatalf("expected %v, got %v", exp, got)
			return
		}
	}

}

func TestManager_Complete(t *testing.T) {
	t.Parallel()

	fake := Product{}

	ready := Product{Quantity: 1}
	ready.Build(0)

	completeProduct := ready
	completeProduct.Build(1)

	testcases := []struct {
		name     string
		employee int
		product  *Product
		err      error
	}{
		{
			name:     "invalid employee",
			employee: -1,
			product:  &fake,
			err:      ErrInvalidEmployee(-1),
		},
		{
			name:     "invalid product, zero quantity",
			employee: 1,
			product:  &fake,
			err:      ErrInvalidQuantity(0),
		},
		{
			name:     "invalid product, not build",
			employee: 1,
			product:  &ready,
			err:      ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", ready)),
		},
		{
			name:     "complete product ready for dispatch",
			employee: 1,
			product:  &completeProduct,
			err:      nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			manager := NewManager()
			defer manager.Stop()

			err := manager.Start(1)
			if err != nil {
				t.Fatal(err)
			}

			go func() {

				got := manager.Complete(Employee(tt.employee), tt.product)
				if got != nil {
					manager.Errors() <- got
				}

			}()

			go func() {
				cp := <-manager.Completed()
				if check := cp.IsValid(); check == nil {
					manager.Stop()
				}
			}()

			for {
				select {
				case err := <-manager.Errors():
					if err.Error() != tt.err.Error() {
						t.Fatalf("expected %v, got %v", tt.err, err)
					}
					return
				case <-manager.Done():
					return
				}
			}
		})
	}

}
