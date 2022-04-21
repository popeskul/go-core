package storage

import "go-search/pkg/crawler"

type Interface interface {
	Add([]crawler.Document)
	Search([]int) []crawler.Document
}
