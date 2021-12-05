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
	ID       int
	Email    string
	Category []Category
}

//News stories categories to subscribe
type Category struct {
	Name string
}

func (s *Subscriber) AddSubscription(category ...string) error {
	return nil
}
func (s *Subscriber) RemoveSubscription(category ...string) error {
	return nil
}

/*
//Subscribe will subscribe a subscriber to the news service
func (ns *NewsService) Subscribe(ctx context.Context) {

}

//UnSubscribe will unsubscribe a subscriber from the news service
func (ns *NewsService) Unsubscribe(ctx context.Context) {

}
*/
