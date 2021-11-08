package week06

import (
	"fmt"
	"testing"
)

func TestManager_Start(t *testing.T) {
	t.Parallel()

	//create new manager
	manager := NewManager()

	testcases := []struct {
		name string
		//m     *Manager
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

			go func() {
				got := manager.Start(tt.count)
				if got != nil {
					manager.Errors() <- got
					manager.Stop()
				}
			}()

			for {
				select {
				case err := <-manager.Errors():
					if err != nil {
						if err.Error() != tt.err.Error() {
							t.Fatalf("expected %v, got %v", tt.err, err)
						}
					}
				case <-manager.Done():
					return
				}
			}

		})
	}
}

func TestManager_Assign(t *testing.T) {
	t.Parallel()

	//Create new manager & stop the manager from further work
	m := NewManager()
	m.Stop()

	fakeProduct := Product{}

	readyProduct := Product{Quantity: 1}
	readyProduct.Build(0)

	fullProduct := Product{Quantity: 1}
	fullProduct.Build(1)

	testcases := []struct {
		name     string
		manager  *Manager
		products *Product
		err      error
	}{
		{
			name:     "manager stopped",
			manager:  m,
			products: &fakeProduct,
			err:      ErrManagerStopped{},
		},
		{
			name:     "manager working, invalid products quantity",
			manager:  NewManager(),
			products: &fakeProduct,
			err:      ErrInvalidQuantity(0),
		},
		{
			name:     "manager working, invalid employee",
			manager:  NewManager(),
			products: &readyProduct,
			err:      ErrInvalidEmployee(Employee(0)),
		},
		{
			name:     "manager working, valid products",
			manager:  NewManager(),
			products: &fullProduct,
			err:      nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.manager.Start(1)
			if err != nil {
				t.Fatal(err)
			}

			got := tt.manager.Assign(tt.products)
			if got != nil {
				if got.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, got)
				}
				tt.manager.Stop()
			}
			tt.manager.Stop()
		})
	}
}

func TestManager_Complete(t *testing.T) {
	t.Parallel()

	manager := NewManager()

	fakeProduct := Product{}

	readyProduct := Product{Quantity: 1}
	readyProduct.Build(0)

	completeProduct := readyProduct
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
			product:  &fakeProduct,
			err:      ErrInvalidEmployee(-1),
		},
		{
			name:     "invalid product, zero quantity",
			employee: 1,
			product:  &fakeProduct,
			err:      ErrInvalidQuantity(0),
		},
		{
			name:     "invalid product, not build",
			employee: 1,
			product:  &readyProduct,
			err:      ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", readyProduct)),
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

			err := manager.Start(1)
			if err != nil {
				t.Fatal(err)
			}

			go func() {

				got := manager.Complete(Employee(tt.employee), tt.product)
				if got != nil {
					manager.Errors() <- got
					return
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
