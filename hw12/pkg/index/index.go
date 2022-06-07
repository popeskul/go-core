package index

import "go-search/hw12/pkg/crawler"

type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
