package newsmill

import (
	"fmt"
	"sort"
	"strings"
)

// Category is the type of category an article belongs to.
type Category string

// Return the category name
func (c Category) Name() string {
	return string(c)
}

// Categories is a list of category with there ids.
type Categories map[Category]int

//
func (cats Categories) String() string {
	lines := make([]string, 0, len(cats))

	for c, id := range cats {
		lines = append(lines, fmt.Sprintf("{%s:%dx}", c, id))
	}
	sort.Strings(lines)
	s := strings.Join(lines, ", ")
	return fmt.Sprintf("[%s]", s)
}
