package storage

import "go-search/hw2/pkg/crawler"

type Interface interface {
	Add([]crawler.Document)
	Search([]int) []crawler.Document
}
