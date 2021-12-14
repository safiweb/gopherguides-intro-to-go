package newsmill

/*
import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
)

type FileFeed struct {
	config SourceConfig
	cancel context.CancelFunc
	sync.RWMutex
	Articles []Article
}

//Run the file feed source
func (f *FileFeed) Run(ctx context.Context, cfg SourceConfig) context.Context {
	f.RLock()
	ctx, f.cancel = context.WithCancel(ctx)
	f.config = cfg
	f.RUnlock()
	return ctx
}

//Get the file feed news stories
func (f *FileFeed) Fetch(articles ...Article) error {

	if len(articles) == 0 {
		return ErrNoArticles(len(articles))
	}

	f.Articles = append(f.Articles, articles...)

	return nil

}

//Publish the file feed news stories
func (f *FileFeed) Publish(s *Service, cats ...Category) (context.Context, error) {

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	if len(cats) == 0 {
		return nil, ErrNoCategories(len(cats))
	}

	articles := f.Articles
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

			articles[i].Source = f.config.Name

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

// GetFileArticles get articles from the defined file
func (f *FileFeed) GetFileArticles() ([]Article, error) {

	if f.config.articlesFile == "" {
		return nil, nil
	}

	articles := []Article{}

	if buf, err := ioutil.ReadFile(f.config.articlesFile); os.IsNotExist(err) {
	} else if err != nil {
		return nil, err
	} else if err := json.Unmarshal(buf, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// Close the file feed source
func (f *FileFeed) Close() {
	f.RLock()
	defer f.RUnlock()

	if f.cancel != nil {
		f.cancel()
	}

}
*/
