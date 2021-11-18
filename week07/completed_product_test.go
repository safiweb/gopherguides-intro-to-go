package week07

import (
	"fmt"
	"testing"
)

func TestCompletedProduct_IsValid(t *testing.T) {

	t.Parallel()

	//Build a product of 1 Quantity for 1 Employee
	prod := Product{Quantity: 1}
	prod.Build(1)

	//Employee
	e := Employee(1)

	//Different Complete Products for different scenarios
	empty := CompletedProduct{}
	noEmployee := CompletedProduct{Product: prod, Employee: Employee(0)}
	notBuild := CompletedProduct{Product: Product{Quantity: 1}, Employee: e}
	good := CompletedProduct{Product: prod, Employee: e}

	testcases := []struct {
		name string
		cp   CompletedProduct
		err  error
	}{
		{
			name: "empty complete product",
			cp:   empty,
			err:  fmt.Errorf("invalid employee number: 0"),
		},
		{
			name: "complete product no employee",
			cp:   noEmployee,
			err:  ErrInvalidEmployee(0),
		},
		{
			name: "complete product no buildby",
			cp:   notBuild,
			err:  ErrProductNotBuilt("product is not built: {1 0}"),
		},
		{
			name: "full complete product",
			cp:   good,
			err:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cp.IsValid()
			if got != nil {
				if got.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, got)
					return
				}
			}
		})
	}

}
