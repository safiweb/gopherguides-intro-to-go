package week08

import (
	"context"
	"testing"
)

func TestWarehouse_Stop(t *testing.T) {
	t.Parallel()

	w := &Warehouse{}
	ctx := context.Background()

	exp := "context canceled"

	ctx = w.Start(ctx)
	w.Stop()

	if ctx.Err().Error() != exp {
		t.Fatalf("expected %v, got %v", exp, ctx.Err().Error())
	}

}
