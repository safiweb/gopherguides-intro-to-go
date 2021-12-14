package newsmill

import (
	"testing"
)

func TestSubscription_IsInvalid(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name string
		sub  Subscription
		err  error
	}{
		{
			name: "invalid article",
			sub:  Subscription(""),
			err:  ErrInvalidSubscription(""),
		},
		{
			name: "valid subcription",
			sub:  Subscription("new"),
			err:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.sub.IsValid()

			if err != nil {
				if tt.err.Error() != err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			}

		})
	}
}
