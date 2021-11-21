package week08

import (
	"context"
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
			name:  "valid employee",
			count: 1,
			err:   nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()
			manager := &Manager{}

			_, err := manager.Start(ctx, tt.count)
			if err != nil {
				if tt.err.Error() != err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			}

		})
	}

}

func TestManager_Assign(t *testing.T) {
	t.Parallel()

	t.Run("manager stop", func(t *testing.T) {

		ctx := context.Background()
		exp := ErrManagerStopped{}

		manager := &Manager{}

		_, err := manager.Start(ctx, 1)
		if err != nil {
			t.Fatal("unexpected error")
		}
		manager.Stop()

		err = manager.Assign(&Product{
			Materials: Materials{
				Wood: 2,
				Oil:  3,
			},
		})

		if err != nil {
			if exp.Error() != err.Error() {
				t.Fatalf("expected %v, got %v", exp, err)
			}
		}

	})

	testcases := []struct {
		name string
		prod *Product
		err  error
	}{
		{
			name: "invalid product",
			prod: &Product{},
			err:  ErrInvalidMaterials(0),
		},
		{
			name: "valid product",
			prod: &Product{Materials: Materials{Oil: 2}},
			err:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()
			manager := &Manager{}

			_, err := manager.Start(ctx, 1)
			if err != nil {
				t.Fatal("unexpected error")
			}

			err = manager.Assign(tt.prod)
			if err != nil {
				if tt.err.Error() != err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			}

		})
	}

}

func TestManager_Complete(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name string
		e    Employee
		p    *Product
		err  error
	}{
		{
			name: "invalid employee",
			e:    Employee(0),
			p:    &Product{Materials: Materials{Wood: 2}},
			err:  ErrInvalidEmployee(0),
		},
		{
			name: "invalid quantity",
			e:    Employee(1),
			p:    &Product{},
			err:  ErrInvalidMaterials(0),
		},
		{
			name: "invalid product",
			e:    Employee(1),
			p:    &Product{Materials: Materials{Oil: 1}, builtBy: 0},
			err:  ErrProductNotBuilt(fmt.Sprintf("product is not built: %v", &Product{Materials: Materials{Oil: 1}, builtBy: 0})),
		},
		{
			name: "valid product/employee",
			e:    Employee(1),
			p:    &Product{Materials: Materials{Wood: 2, Oil: 3}},
			err:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()
			manager := &Manager{}

			manager.Start(ctx, 1)

			tt.p.Build(tt.e, &Warehouse{})

			go func() {
				got := manager.Complete(Employee(tt.e), tt.p)
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
				case <-ctx.Done():
					return
				case err := <-manager.Errors():
					if err != nil {
						if err.Error() != tt.err.Error() {
							t.Fatalf("expected %v, got %v", tt.err, err)
						}
					}
					return
				}
			}

		})
	}
}
