package index

import "go-search/pkg/crawler"

type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
