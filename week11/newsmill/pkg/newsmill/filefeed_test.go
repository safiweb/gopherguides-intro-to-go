package newsmill

import (
	"context"
	"fmt"
	"testing"
)

func TestArticle_sss(t *testing.T) {
	ctx := context.Background()

	//initilize the service
	service := NewService()

	//starts the service
	ctx, err := service.Start(ctx)
	if err != nil {
		fmt.Println(err)
	}

	//Add subscribers
	service.Subscribe("sports")

	//Start a New source
	source := &NewsFeed{}
	ctx = source.Run(ctx, SourceConfig{Name: "Mock"})

	//Fetch articles
	source.Fetch(*Article1, *Article2)

	//Publish sports articles
	ch := source.Publish("sports")

	//Disptch the stories to subscribers
	service.Dispatch(ctx, ch)

	// News subscribe are getting
	for _, sub := range service.subs {
		fmt.Println(<-sub)
	}

	//Unsubscribe
	service.Unsubscribe("sports")

	//Stop the service
	service.Stop()

}
