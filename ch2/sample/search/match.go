package search

import (
	"log"
)

type Result struct {
	Field   string
	Content string
}
type Matcher interface {
	Search(feed *Feed, searchItem string) ([]*Result, error)
}
