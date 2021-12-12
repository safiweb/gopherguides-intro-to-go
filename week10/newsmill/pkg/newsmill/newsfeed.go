package newsmill

import (
	"context"
	"sync"
)

//News feed mock articles
var (
	Article1 = &Article{
		Id: 1,
		Categories: Categories{
			"sport": 22,
			"news":  25,
		},
	}

	Article2 = &Article{
		Id: 2,
		Categories: Categories{
			"politics": 22,
			"breaking": 23,
		},
	}

	Article3 = &Article{
		Id: 3,
		Categories: Categories{
			"breaking": 23,
			"world":    24,
		},
	}

	Article4 = &Article{
		Id: 4,
		Categories: Categories{
			"sport":  22,
			"gossip": 26,
		},
	}
)

type NewsFeed struct {
	config SourceConfig
	cancel context.CancelFunc
	sync.RWMutex
	Articles []Article
}

//Run the news feed source
func (n *NewsFeed) Run(ctx context.Context, cfg SourceConfig) context.Context {
	n.RLock()
	ctx, n.cancel = context.WithCancel(ctx)
	n.config = cfg
	n.RUnlock()
	return ctx
}

//Get the news feed news stories
func (n *NewsFeed) Fetch(articles ...Article) error {

	if len(articles) == 0 {
		return ErrNoArticles(len(articles))
	}

	n.Articles = append(n.Articles, articles...)

	return nil

}

//Publish the newsfeed news stories
func (n *NewsFeed) Publish(s *Service, cats ...Category) (context.Context, error) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	if len(cats) == 0 {
		return nil, ErrNoCategories(len(cats))
	}

	articles := n.Articles
	if len(articles) == 0 {
		return nil, ErrNoArticles(len(articles))
	}

	for _, j := range cats {
		for i := range articles {

			if err := articles[i].IsValid(); err != nil {
				return nil, err
			}

			_, ok := articles[i].Categories[j]
			if !ok {
				continue
			}

			articles[i].Source = n.config.Name

			go func(i int) {
				select {
				case s.broadcastCh() <- &articles[i]:
				default:
				}
			}(i)

			//s.broadcast <- &articles[i]

		}
	}

	return ctx, nil

}

// Close the news feed source
func (n *NewsFeed) Close() {
	n.RLock()
	defer n.RUnlock()

	if n.cancel != nil {
		n.cancel()
	}

}
