package week07

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestManager_Run(t *testing.T) {
	t.Parallel()

	t.Run("timeout after 5 seconds if nothing happens", func(t *testing.T) {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		//A product with large quantity that will take more than 5 secs to process
		prod := &Product{Quantity: 10000}

		exp := context.DeadlineExceeded.Error()

		_, err := Run(ctx, 1, prod)

		if err != nil {
			t.Fatal(err)
		}

		<-ctx.Done()

		if ctx.Err().Error() != exp {
			t.Fatalf("unexpected value, got: %v, exp: %v", ctx.Err().Error(), exp)
		}

	})

	t.Run("invalid product quantity", func(t *testing.T) {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		count := 1
		prod := &Product{Quantity: 0}
		exp := ErrInvalidQuantity(0)

		cp, err := Run(ctx, count, prod)

		if err != nil {
			if err.Error() != exp.Error() {
				t.Fatalf("unexpected value, got: %v, exp: %v", err, exp)
			}
		}

		if cp != nil {
			t.Fatalf("unexpected value, got: %v, exp: %v", cp, nil)
		}

	})

	t.Run("invalid employee count", func(t *testing.T) {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		count := 0
		prod := &Product{Quantity: 1}
		exp := ErrInvalidEmployeeCount(0)

		cp, err := Run(ctx, count, prod)

		if err != nil {
			if err.Error() != exp.Error() {
				t.Fatalf("unexpected value, got: %v, exp: %v", err, exp)
			}
		}

		if cp != nil {
			t.Fatalf("unexpected value, got: %v, exp: %v", cp, nil)
		}

	})

	t.Run("test that the output is correct", func(t *testing.T) {

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		shirt := &Product{Quantity: 1}
		cap := &Product{Quantity: 2}
		short := &Product{Quantity: 1}

		count := 3

		cp, err := Run(ctx, count, shirt, cap, short)

		if err != nil {
			t.Fatal(err)
		}

		if count != len(cp) {
			t.Fatalf("unexpected value, got: %v, exp: %v", len(cp), count)
		}

	})

}

func TestManager_Assign_Stop(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Create new manager & stop the manager from further work
	manager := NewManager()
	manager.Start(ctx, 1)
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

func TestManager_Complete_Stop(t *testing.T) {
	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//Create new manager & stop the manager from further work
	manager := NewManager()
	manager.Start(ctx, 1)
	manager.Stop()

	ready := Product{Quantity: 1}
	ready.Build(1)

	exp := ErrManagerStopped{}

	got := manager.Complete(Employee(1), &ready)
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

	complete := ready
	complete.Build(1)

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
			product:  &complete,
			err:      nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			manager := NewManager()

			ctx, err := manager.Start(ctx, 1)
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
				case <-ctx.Done():
					return
				}
			}
		})
	}

}
