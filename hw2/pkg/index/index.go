package index

import "go-search/pkg/crawler"

type Index interface {
	Add([]crawler.Document)
}
