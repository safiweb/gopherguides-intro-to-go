package newsmill

import (
	"context"
	"os"
	"sync"
)

type FileFeed struct {
	config SourceConfig
	cancel context.CancelFunc
	sync.RWMutex
	Articles []Article
}

// Run the file feed source
func (f *FileFeed) Run(ctx context.Context, cfg SourceConfig) context.Context {
	f.RLock()
	ctx, f.cancel = context.WithCancel(ctx)
	f.config = cfg
	f.RUnlock()
	return ctx
}

//Get the news feed news stories from file
func (f *FileFeed) WatchFileArticles() error {

	if f.config.DirPath == "" {
		return ErrNoFilePath{}
	}

	_, err := os.ReadDir(f.config.DirPath)
	if err != nil {
		return ErrNoDirFound{}
	}

	/*if len(articles) == 0 {
		return ErrNoArticles(len(articles))
	}

	n.Articles = append(n.Articles, articles...)*/

	return nil

}

//Get the file feed news stories
func (n *FileFeed) Get(articles ...Article) error {

	if len(articles) == 0 {
		return ErrNoArticles(len(articles))
	}

	n.Articles = append(n.Articles, articles...)

	return nil

}

//Publish the file feed news stories
func (n *FileFeed) Publish(cats ...Category) (<-chan interface{}, error) {

	if len(cats) == 0 {
		return nil, ErrNoCategories(len(cats))
	}

	articles := n.Articles
	if len(articles) == 0 {
		return nil, ErrNoArticles(len(articles))
	}

	publish := make(chan interface{})

	go func() {
		for _, j := range cats {
			for i := range articles {
				_, ok := articles[i].Categories[j]
				if !ok {
					continue
				}
				articles[i].Source = n.config.Name
				publish <- articles[i]
				//fmt.Println(articles[i])
			}
			close(publish)
		}
	}()

	return publish, nil

}

// Close the file feed source
func (f *FileFeed) Close() {
	f.RLock()
	defer f.RUnlock()

	if f.cancel != nil {
		f.cancel()
	}

}
