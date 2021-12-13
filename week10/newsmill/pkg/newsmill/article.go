package newsmill

// Article is the single type that a source returns.
type Article struct {
	Id         int        `json:"id"`
	Categories Categories `json:"categories"`
	Source     string     `json:"source"`
}

// IsValid returns an error if the article is invalid.
func (a *Article) IsValid() error {

	if len(a.Categories) == 0 {
		return ErrInvalidCategories(len(a.Categories))
	}

	return nil
}

// Articles is a list of articles.
//type Articles []Article
