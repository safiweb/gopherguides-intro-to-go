package newsmill

import (
	"context"
	"fmt"
	"testing"
)

func TestService_Start(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	s := NewService()
	got, _ := s.Start(ctx)

	want := "context.Background.WithCancel"

	if fmt.Sprint(got) != want {
		t.Fatalf("expected %v, got %v", want, got)
	}

}

func TestService_Subscribe(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name string
		sub  Subscription
		err  error
	}{
		{
			name: "invalid subscription",
			sub:  Subscription(""),
			err:  ErrInvalidSubscription(Subscription("")),
		},
		{
			name: "subscription exists",
			sub:  Subscription("tech"),
			err:  ErrSubscriptionExist(fmt.Sprintf("subscription %s already exist", "tech")),
		},
		{
			name: "valid subscription",
			sub:  Subscription("sports"),
			err:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()
			s := NewService()
			s.Start(ctx)

			s.Subscribe("tech")

			_, err := s.Subscribe(tt.sub)

			if err != nil {
				if tt.err.Error() != err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			}

		})
	}

}

func TestService_Unsubscribe(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name string
		sub  Subscription
		err  error
	}{
		{
			name: "subscription not found",
			sub:  Subscription("tech"),
			err:  ErrSubscriptionNotFound("tech"),
		},
		{
			name: "unsubscribe successfull",
			sub:  Subscription("sports"),
			err:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()
			s := NewService()
			s.Start(ctx)

			s.Subscribe("sports")

			err := s.Unsubscribe(tt.sub)

			if err != nil {
				if tt.err.Error() != err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			}

		})
	}

}

func TestService_Dispatch(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	s := NewService()
	s.Start(ctx)

	s.Subscribe("sports")

	source := &NewsFeed{}
	ctx = source.Run(ctx, SourceConfig{articlesFile: "Mock FIles"})
	source.Fetch(*Article1, *Article2)
	publishCh := source.Publish("sports")

	s.Dispatch(ctx, publishCh)

	want := 1
	got := <-s.subs["sports"]

	if len(got.([]Article)) != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestService_Stop(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	s := NewService()
	s.Start(ctx)

	s.Subscribe("sports")

	s.Stop()

	want := true
	got := s.stopped

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
