package newsmill

import (
	"context"
	"testing"
)

func TestNewsFeed_Fetch(t *testing.T) {
	t.Parallel()

	t.Run("no articles", func(t *testing.T) {

		source := &NewsFeed{}
		err := source.Fetch()
		want := ErrNoArticles(0)

		if err != nil {
			if want.Error() != err.Error() {
				t.Fatalf("expected %v, got %v", want, err)
			}
		}

	})

}

func TestNewsFeed_Cancel(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	source := &NewsFeed{}

	source.Run(ctx, SourceConfig{articlesFile: "Mock FIles"})

	source.Fetch(*Article1)

	source.Close()

	want := []Article{}
	got := source.Articles

	if len(got) != len(want) {
		t.Fatalf("expected %v, got %v", len(want), len(got))
	}
}

func TestNewsFeed_articlesToPublish(t *testing.T) {
	t.Parallel()

	t.Run("no categories", func(t *testing.T) {

		source := &NewsFeed{}
		_, err := source.articlesToPublish()
		want := ErrNoCategories(0)

		if err != nil {
			if want.Error() != err.Error() {
				t.Fatalf("expected %v, got %v", want, err)
			}
		}

	})

	t.Run("no articles", func(t *testing.T) {

		source := &NewsFeed{}

		_, err := source.articlesToPublish("fresh")
		want := ErrNoArticles(0)

		if err != nil {
			if want.Error() != err.Error() {
				t.Fatalf("expected %v, got %v", want, err)
			}
		}

	})

	t.Run("success articles", func(t *testing.T) {

		source := &NewsFeed{}
		source.Fetch(*Article1, *Article2, *Article3)

		_, err := source.articlesToPublish("sport")
		want := error(nil)

		if err != nil {
			if want.Error() != err.Error() {
				t.Fatalf("expected %v, got %v", want, err)
			}
		}

	})

}
