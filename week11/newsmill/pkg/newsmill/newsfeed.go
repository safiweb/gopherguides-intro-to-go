package newsmill

import (
	"context"
	"sync"
)

//News feed mock articles
var (
	Article1 = &Article{
		Id:       1,
		Category: "sports",
	}

	Article2 = &Article{
		Id:       2,
		Category: "politics",
	}

	Article3 = &Article{
		Id:       3,
		Category: "breaking",
	}

	Article4 = &Article{
		Id:       4,
		Category: "tech",
	}

	Article5 = &Article{
		Id:       5,
		Category: "world",
	}
	Article6 = &Article{
		Id:       6,
		Category: "sports",
	}

	Article7 = &Article{
		Id:       7,
		Category: "politics",
	}
)

type NewsFeed struct {
	publisherCh chan interface{} // broadcast channel
	config      SourceConfig
	cancel      context.CancelFunc
	sync.RWMutex
	once     sync.Once
	Articles []Article
}

//Run the news feed source
func (n *NewsFeed) Run(ctx context.Context, cfg SourceConfig) context.Context {
	n.RLock()
	ctx, n.cancel = context.WithCancel(ctx)
	n.config = cfg
	n.publisherCh = make(chan interface{})
	n.RUnlock()
	return ctx
}

//Get the news feed news stories
func (n *NewsFeed) Fetch(articles ...Article) error {

	if len(articles) == 0 {
		return ErrNoArticles(len(articles))
	}

	if n.Articles == nil {
		n.Articles = []Article{}
	}

	n.Articles = append(n.Articles, articles...)

	return nil

}

// Publish the newsfeed news stories
// Gives a channel that generate the articles to publish
func (n *NewsFeed) Publish(cats ...string) <-chan []Article {
	out := make(chan []Article)
	articles, _ := n.articlesToPublish(cats...)
	go func() {
		//for _, n := range articles {
		out <- articles
		//}
		close(out)
	}()
	return out
}

// articlesToPublish get articles to publish
func (n *NewsFeed) articlesToPublish(cats ...string) ([]Article, error) {

	if len(cats) == 0 {
		return nil, ErrNoCategories(len(cats))
	}

	artToPublish := []Article{}
	articles := n.Articles

	if len(articles) == 0 {
		return nil, ErrNoArticles(len(articles))
	}

	for _, article := range articles {
		for _, cat := range cats {
			if cat == article.Category {
				n.config.articlesFile = "News Feed"
				artToPublish = append(artToPublish, article)
			}
		}
	}

	/*go func() {
		n.Publisher() <- artToPublish
	}()*/

	return artToPublish, nil

}

// Publisher will return a channel that can be listened to
// for new articles to be read.
func (n *NewsFeed) Publisher() chan interface{} {
	return n.publisherCh
}

// Close the news feed source
func (n *NewsFeed) Close() {
	n.RLock()
	defer n.RUnlock()

	n.once.Do(func() {

		n.Lock()
		defer n.Unlock()

		n.cancel()

		// close all channels
		if n.publisherCh != nil {
			close(n.publisherCh)
		}
	})

}
