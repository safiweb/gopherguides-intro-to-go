package newsmill

import "fmt"

// ErrInvalidSubscription is returned when subscription id is invalid.
type ErrInvalidSubscription string

func (e ErrInvalidSubscription) Error() string {
	return fmt.Sprintf("invalid subcription: %s", Subscription(e))
}

// ErrInvalidCategories is returned when article categories are less or equal to 0
type ErrInvalidCategories int

func (e ErrInvalidCategories) Error() string {
	return fmt.Sprintf("categories must be greater than: %d", e)
}

// ErrNoArticles is returned when articles provided are less or equal to 0
type ErrNoArticles int

func (e ErrNoArticles) Error() string {
	return fmt.Sprintf("articles must be greater than: %d", e)
}

// ErrNoCategories is returned when categories provided are less or equal to 0
type ErrNoCategories int

func (e ErrNoCategories) Error() string {
	return fmt.Sprintf("categories must be greater than: %d", e)
}

/*
type ErrNoFilePath struct{}

func (ErrNoFilePath) Error() string {
	return "no filepath declared"
}

type ErrNoDirFound struct{}

func (ErrNoDirFound) Error() string {
	return "The system cannot find the folder specified"
}
*/

type ErrSubscriptionNotFound string

func (e ErrSubscriptionNotFound) Error() string {
	return fmt.Sprintf("subcription with topic: %v not found", Subscription(e))
}

/*
// ErrInvalidTopicCount is returned when the subscription topics count is invalid.
type ErrInvalidTopicCount int

func (e ErrInvalidTopicCount) Error() string {
	return fmt.Sprintf("invalid topic(s) count: %d", e)
}
*/

// ErrSubscriptionExist is returned when the subscription topics count is invalid.
type ErrSubscriptionExist string

func (e ErrSubscriptionExist) Error() string {
	return string(e)
}
