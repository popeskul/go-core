package index

import "go-search/hw11/server/pkg/crawler"

type Interface interface {
	Add([]crawler.Document)
	Search(string) []int
}
