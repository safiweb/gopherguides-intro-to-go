package week05

import (
	"fmt"
	"testing"
)

func TestErrors_TableNotFound(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name     string
		errTable ErrTableNotFound
		want     string
	}{
		{
			name:     "empty err table",
			errTable: ErrTableNotFound{},
			want:     "",
		},
		{
			name:     "err table with data",
			errTable: ErrTableNotFound{table: "users"},
			want:     "users",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			got := tt.errTable.TableNotFound()

			if got != tt.want {
				t.Fatalf("unexpected error: got %v, want %v", got, tt.want)
			}

		})
	}
}

func TestErrors_IsErrTableNotFound(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "not ErrTableNotFound error",
			err:  fmt.Errorf("hello"),
			want: false,
		},
		{
			name: "not not ErrTableNotFound error",
			err:  &ErrTableNotFound{table: "users"},
			want: true,
		},
		{
			name: "empty ErrTableNotFound error",
			err:  &ErrTableNotFound{},
			want: true,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			got := IsErrTableNotFound(tt.err)

			if got != tt.want {
				t.Fatalf("unexpected error: got %v, want %v", got, tt.want)
			}

		})
	}
}

func TestErrors_Clauses(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name string
		err  *errNoRows
		want Clauses
	}{
		{
			name: "empty clauses provided",
			err:  &errNoRows{nil, "users"},
			want: nil,
		},
		{
			name: "clauses provided correct expected",
			err:  &errNoRows{Clauses{"hobby": "golf", "name": "Webbs"}, "users"},
			want: Clauses{"name": "Webbs", "hobby": "golf"},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.err.Clauses()
			assertEqualClause(t, got, tt.want)

		})
	}
}

func TestErrors_RowNotFound(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name    string
		err     *errNoRows
		want    string
		clauses Clauses
	}{
		{
			name:    "empty clauses and table name",
			err:     &errNoRows{Clauses{}, ""},
			want:    "",
			clauses: Clauses{},
		},
		{
			name:    "empty clauses, correct table name provided",
			err:     &errNoRows{Clauses{}, "users"},
			want:    "users",
			clauses: Clauses{},
		},
		{
			name:    "correct clauses, correct table provided",
			err:     &errNoRows{Clauses{"hobby": "golf", "name": "Webbs"}, "users"},
			want:    "users",
			clauses: Clauses{"hobby": "golf", "name": "Webbs"},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got, c := tt.err.RowNotFound()

			if got != tt.want {
				t.Fatalf("unexpected error: got %v, want %v", got, tt.want)
			}

			assertEqualClause(t, c, tt.clauses)

		})
	}

}

func TestErrors_IsErrNoRows(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "errNoRows error type, empty clauses",
			err:  &errNoRows{Clauses{}, "users"},
			want: true,
		},
		{
			name: "errNoRows error type, correct data",
			err:  &errNoRows{Clauses{"hobby": "golf", "name": "Webbs"}, "users"},
			want: true,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			got := IsErrNoRows(tt.err)

			if got != tt.want {
				t.Fatalf("unexpected error: got %v, want %v", got, tt.want)
			}

		})
	}

}

func TestErrors_AsErrNoRows(t *testing.T) {

	t.Parallel()

	testcases := []struct {
		name    string
		err     error
		errRows ErrNoRows
		want    bool
	}{
		{
			name:    "not errNoRows error type",
			err:     fmt.Errorf("hello"),
			errRows: &errNoRows{},
			want:    false,
		},
		{
			name:    "errNoRows error type, empty query clauses",
			err:     &errNoRows{Clauses{}, "users"},
			errRows: &errNoRows{Clauses{}, "users"},
			want:    true,
		},
		{
			name:    "errNoRows error type, wrong query clauses",
			err:     &errNoRows{Clauses{"hobby": "golf", "name": "Webbs"}, "users"},
			errRows: &errNoRows{Clauses{"hobby": "golf", "name": "Webbs"}, "users"},
			want:    true,
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {

			got, gotWant := AsErrNoRows(tt.err)

			if got != nil {
				if got.Error() != tt.errRows.Error() {
					t.Fatalf("unexpected error: got %v, want %v", got.Error(), tt.errRows.Error())
				}
			}

			if gotWant != tt.want {
				t.Fatalf("unexpected error: got %v, want %v", gotWant, tt.want)
			}

		})
	}

}

func assertEqualClause(t testing.TB, got Clauses, want Clauses) {
	t.Helper()

	for k, v := range want {
		if got[k] != v {
			t.Fatalf("unexpected error: got %v, want %v", got, want)
		}
	}
}
