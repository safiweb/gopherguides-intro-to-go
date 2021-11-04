package week05

import (
	"errors"
	"testing"
)

func TestStore_All(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name  string
		table string
		store *Store
		err   error
	}{
		{
			name:  "store with no data",
			table: "users",
			store: &Store{},
			err:   ErrTableNotFound{},
		},
		{
			name:  "store with data",
			table: "users",
			store: &Store{data: data{}},
			err:   ErrTableNotFound{},
		},
		{
			name:  "store with data, users",
			table: "users",
			store: &Store{data: data{"users": Models{}}},
			err:   nil,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			_, err := tt.store.All(tt.table)

			if tt.err != nil {
				ok := errors.Is(err, tt.err)
				if !ok {
					t.Fatalf("expected error %v, got %v", tt.err, err)
				}
				return
			}

			assertError(t, err)
		})
	}

}

func TestStore_Len(t *testing.T) {

	t.Parallel()

	john := Model{
		"name":   "John",
		"gender": "M",
	}

	tom := Model{
		"name":   "Tom",
		"gender": "M",
	}

	myStore := &Store{}
	myStore.Insert("users", john, tom)

	testcases := []struct {
		name  string
		table string
		err   error
		want  int
	}{
		{
			name:  "table exists",
			table: "users",
			err:   nil,
			want:  2,
		},
		{
			name:  "table doesn't exist",
			table: "people",
			err:   &ErrTableNotFound{table: "people"},
			want:  0,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			got, err := myStore.Len(tt.table)

			if tt.err != nil {
				if tt.err.Error() != err.Error() {
					t.Fatalf("unexpected error: got %v, want %v", err, tt.err)
				}
				return
			}

			assertError(t, err)

			if got != tt.want {
				t.Fatalf("unexpected error: got %v, want %v", got, tt.want)
			}

		})
	}

}

func TestStore_Insert(t *testing.T) {

	t.Parallel()

	want := 2
	myStore := &Store{}
	myStore.Insert("users", Model{"name": "Jane", "gender": "F"}, Model{"name": "Anthony", "gender": "M"})

	got, err := myStore.Len("users")
	assertError(t, err)

	if got != want {
		t.Fatalf("unexpected error: got %v, want %v", got, want)
	}

}

func TestStore_Select(t *testing.T) {

	t.Parallel()

	john := Model{
		"name":   "John",
		"gender": "M",
	}

	tom := Model{
		"name":   "Tom",
		"gender": "M",
	}

	myStore := &Store{}
	myStore.Insert("users", john, tom)
	myStore.Insert("items")

	testcases := []struct {
		name  string
		table string
		query Clauses
		err   error
		want  Models
	}{
		{
			name:  "table doesn't exist",
			table: "orders",
			query: Clauses{},
			err:   ErrTableNotFound{table: "orders"},
			want:  Models{},
		},
		{
			name:  "row doesn't exist",
			table: "users",
			query: Clauses{"name": "Jane", "gender": "F"},
			err:   &errNoRows{table: "users"},
			want:  Models{},
		},
		{
			name:  "correct table, row requested",
			table: "users",
			query: Clauses{"name": "Tom", "gender": "M"},
			err:   nil,
			want:  Models{Model{"name": "Tom", "gender": "M"}},
		},
		{
			name:  "no data",
			table: "users",
			query: Clauses{},
			err:   nil,
			want:  Models{},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			got, err := myStore.Select(tt.table, tt.query)

			if tt.err != nil {
				ok := errors.Is(err, tt.err)
				if !ok {
					t.Fatalf("expected error %v, got %v", tt.err, err)
				}
				return
			}

			assertError(t, err)

			for i, model := range tt.want {
				assertEqualModel(t, got[i], model)
			}

		})
	}

}

func assertError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertEqualModel(t testing.TB, got Model, want Model) {
	t.Helper()
	for k, v := range want {
		if got[k] != v {
			t.Fatalf("unexpected error: got %v, want %v", got, want)
		}
	}
}
