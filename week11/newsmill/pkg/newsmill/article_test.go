package newsmill

import (
	"testing"
)

func TestArticle_IsInvalid(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		name string
		art  Article
		err  error
	}{
		{
			name: "invalid article",
			art:  Article{},
			err:  ErrInvalidCategories(0),
		},
		{
			name: "valid article",
			art:  Article{Id: 2, Category: "sports"},
			err:  nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			err := tt.art.IsValid()

			if err != nil {
				if tt.err.Error() != err.Error() {
					t.Fatalf("expected %v, got %v", tt.err, err)
				}
			}

		})
	}
}
