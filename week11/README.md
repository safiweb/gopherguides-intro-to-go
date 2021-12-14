# Newsmill

Newsmill is a Go package project that implements a news service. This service can have many subscribers listening to multiple news stories and  also have many sources for content.

It's also implemets a cli to manage the service.

## Goals

* **Easy** to understand and run.

## Install
```sh
go get github.com/safiweb/gopherguides-intro-to-go/week11/pkg/newsmill
```
## Getting Started
After installing, one can start using the news service package by calling below methods:

Starting the Service
```go
ctx := context.Background()
service := NewService()
service.Start(ctx)
```

Adding Subscriptions
```go
service.Subscribe("sports")
```

Unsubscribing Subscriptions
```go
service.Unsubscribe("sports")
```

Disptching news stories
```go
service.Dispatch(...)
```

Adding sources
```go
source := &NewsFeed{}
source.Run(...)
```

Fetching new stores
```go
source.Fetch(...)
```

Publishing new stories
```go
source.Publish(..)
```

## Full Example Code
```go
func TestService_Newservice(t *testing.T) {
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
```

## Testing
```go
go test ./... -v -cover -race
```

## Cli
Work in progress
