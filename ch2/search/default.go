package search

import (
	"log"
)

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

type Result struct {
	Field   string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchItem string) ([]*Result, error)
}
