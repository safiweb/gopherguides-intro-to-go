package week06

import (
	"fmt"
	"testing"
)

func TestProduct_BuiltBy(t *testing.T) {
	t.Parallel()

	fakeProduct := Product{}

	myProduct := Product{Quantity: 1}
	myProduct.Build(1)

	got := myProduct.BuiltBy()
	exp := Employee(1)

	if got != exp {
		t.Fatalf("expected %d, got %d", exp, got)
	}

	got = fakeProduct.BuiltBy()
	exp = Employee(0)

	if got != exp {
		t.Fatalf("expected %d, got %d", exp, got)
	}

}

func TestProduct_Build(t *testing.T) {
	t.Parallel()

	emptyProduct := Product{}
	readyProduct := Product{Quantity: 1}

	testcases := []struct {
		name    string
		product *Product
		worker  Employee
		err     error
	}{
		{
			name:    "invalid product",
			product: &emptyProduct,
			worker:  Employee(1),
			err:     ErrInvalidQuantity(emptyProduct.Quantity),
		},
		{
			name:    "valid product, bad employee",
			product: &readyProduct,
			worker:  Employee(-1),
			err:     ErrInvalidEmployee(Employee(-1)),
		},
		{
			name:    "valid product, good employee",
			product: &readyProduct,
			worker:  Employee(1),
			err:     nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.product.Build(tt.worker)
			if got != nil {
				if got.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, got)
				}
			}
		})
	}

}

func TestProduct_IsBuilt(t *testing.T) {
	t.Parallel()

	fakeProduct := Product{}

	readyProduct := Product{Quantity: 1}
	readyProduct.Build(0)

	completeProduct := readyProduct
	completeProduct.Build(1)

	testcases := []struct {
		name    string
		product Product
		err     error
	}{
		{
			name:    "empty fake product",
			product: fakeProduct,
			err:     ErrInvalidQuantity(fakeProduct.Quantity),
		},
		{
			name:    "ready product no worker",
			product: readyProduct,
			err:     ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", readyProduct)),
		},
		{
			name:    "complete product with worker",
			product: completeProduct,
			err:     nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.product.IsBuilt()
			if got != nil {
				if got.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, got)
				}
			}
		})
	}

}
