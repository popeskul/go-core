package storage

import (
	"go-search/hw4/pkg/crawler"
	"io"
)

type Interface interface {
	Add([]crawler.Document)
	Search([]int) []crawler.Document
	Read(r io.Reader) ([]crawler.Document, error)
	Write(w io.Writer, docs []crawler.Document) (int, error)
}
