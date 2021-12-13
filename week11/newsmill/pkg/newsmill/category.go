package newsmill

// Category is the type of category an article belongs to.
type Category string

// Return the category name
func (c Category) Name() string {
	return string(c)
}

// Categories is a list of category with there ids.
type Categories map[Category]int
