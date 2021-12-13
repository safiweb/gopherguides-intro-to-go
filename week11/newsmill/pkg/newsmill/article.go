package newsmill

// Article is the single type that a source returns.
type Article struct {
	Id       int    `json:"id"`
	Category string `json:"category"`
	Source   string `json:"source"`
}

// IsValid returns an error if the article is invalid.
func (a *Article) IsValid() error {

	if len(a.Category) == 0 {
		return ErrInvalidCategories(len(a.Category))
	}

	return nil
}

// Articles is a list of articles.
//type Articles []Article
