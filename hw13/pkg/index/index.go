package index

import "go-search/hw13/pkg/crawler"

type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
