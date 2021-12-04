package week08

import (
	"fmt"
	"testing"
)

func TestCompletedProduct_IsValid(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name string
		cp   CompletedProduct
		err  error
	}{
		{
			name: "invalid employee",
			cp: CompletedProduct{
				Product:  Product{Materials: Materials{Oil: 1}},
				Employee: Employee(0),
			},
			err: ErrInvalidEmployee(0),
		},
		{
			name: "not built",
			cp: CompletedProduct{
				Product:  Product{Materials: Materials{Oil: 2}},
				Employee: Employee(1),
			},
			err: ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", &Product{Materials: Materials{Oil: 2}})),
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cp.IsValid()
			if err != nil {
				if err.Error() != tt.err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			}
		})
	}
}
