package week08

import "testing"

func TestProduct_BuiltBy(t *testing.T) {
	t.Parallel()

	fake := Product{}

	prod := Product{Materials: Materials{Oil: 2}}
	prod.Build(1, &Warehouse{})

	got := prod.BuiltBy()
	exp := Employee(1)

	if got != exp {
		t.Fatalf("expected %d, got %d", exp, got)
	}

	got = fake.BuiltBy()
	exp = Employee(0)

	if got != exp {
		t.Fatalf("expected %d, got %d", exp, got)
	}

}
