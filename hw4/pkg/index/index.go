package index

import "go-search/hw4/pkg/crawler"

type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
