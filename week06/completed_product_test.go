package week06

import (
	"fmt"
	"testing"
)

func TestCompletedProduct_IsValid(t *testing.T) {

	t.Parallel()

	//Build a product of 1 Quantity for 1 Employee
	myProduct := Product{Quantity: 1}
	myProduct.Build(1)

	//Build a product of 0 Quantity for 1 Employee
	myProductNoQuantity := Product{}
	myProductNoQuantity.Build(1)

	//Different Complete Products for different scenarios
	emptyCompleteProduct := CompletedProduct{}
	completeProductNoEmployee := CompletedProduct{Product: myProduct, Employee: Employee(0)}
	completeProductNotBuild := CompletedProduct{Product: Product{Quantity: 1}, Employee: Employee(1)}
	completeProductNoQuantity := CompletedProduct{Product: myProductNoQuantity, Employee: Employee(1)}
	completeProductFull := CompletedProduct{Product: myProduct, Employee: Employee(1)}

	testcases := []struct {
		name string
		cp   CompletedProduct
		err  error
	}{
		{
			name: "empty complete product",
			cp:   emptyCompleteProduct,
			err:  fmt.Errorf("invalid employee number: 0"),
		},
		{
			name: "complete product no employee",
			cp:   completeProductNoEmployee,
			err:  fmt.Errorf("invalid employee number: 0"),
		},
		{
			name: "complete product no buildby",
			cp:   completeProductNotBuild,
			err:  fmt.Errorf("product is not built: {1 0}"),
		},
		{
			name: "complete product no quantity",
			cp:   completeProductNoQuantity,
			err:  fmt.Errorf("quantity must be greater than 0, got 0"),
		},
		{
			name: "full complete product",
			cp:   completeProductFull,
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
