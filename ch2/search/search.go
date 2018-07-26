package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchItem string) {
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	results := make(chan *Result)
	var waitGroup sync.WaitGroup

	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
	}
	go func(matcher Matcher, feed *Feed) {
		Match(matcher, feed, searchItem, results)
		waitGroup.Done()
	}(Matcher, feed)

	go func() {
		waitGroup.Wait()
		close(results)
	}()
	Display(results)
}
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}
	log.Println("Register", feedType, "matcher")
	matcher[feedType] = matcher
}
