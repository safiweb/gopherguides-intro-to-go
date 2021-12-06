package newsmill

/*
Subscribers should be able to receive news stories for the categories they are subscribed to.
Subscribers should be able to subscribe to the news service and receive news stories for one or more categories.
Subscribers should be able to unsubscribe from the news service. Other subscribers should not be affected.
Subscribers should be cancelled by the news service when the news service is stopped.
Subscribers should not be aware of each other, nor should they have any direct contact with the news sources.
Subscribers should not be effected by the removal of a news source.
*/

//Subscriber to subscribe to news stories category/categories
type Subscriber struct {
	ID     int
	Name   string
	Topics []string
}

type Subscription struct {
	subID      int
	name       string
	Subscriber Subscriber
}
