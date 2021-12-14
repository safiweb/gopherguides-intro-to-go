package newsmill

import (
	"testing"
)

func TestCategory_Name(t *testing.T) {
	t.Parallel()

	want := "people"

	c := Category("people")

	got := c.Name()

	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}
