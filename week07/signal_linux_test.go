package week07

import (
	"context"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

const TEST_SIGNAL = syscall.SIGUSR2

func TestManager_Run_Signals(t *testing.T) {
	t.Run("interruption by a signal", func(t *testing.T) {

		shirt := &Product{Quantity: 1}
		cap := &Product{Quantity: 2}
		short := &Product{Quantity: 1}

		count := 3

		ctx := context.Background()

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		sigCtx, cancel := signal.NotifyContext(ctx, TEST_SIGNAL)
		defer cancel()

		go Run(ctx, count, shirt, cap, short)

		go func() {
			time.Sleep(time.Second)
			syscall.Kill(syscall.Getpid(), TEST_SIGNAL)
		}()

		select {
		case <-ctx.Done():
		case <-sigCtx.Done():
			return
		}

		err := ctx.Err()
		if err == nil {
			return
		}

		if err == context.DeadlineExceeded {
			t.Fatal("unexpected error", err)
		}
	})
}
