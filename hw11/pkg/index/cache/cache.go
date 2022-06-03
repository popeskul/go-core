package cache

import (
	"go-search/hw11/pkg/crawler"
	"sort"
	"strings"
)

type Index struct {
	store map[string][]int
}

func New() *Index {
	return &Index{
		store: make(map[string][]int),
	}
}

func (index *Index) Add(docs []crawler.Document) {
	for _, doc := range docs {
		for _, word := range splitWords(doc.Title) {
			if !isExist(index.store[word], doc.ID) {
				index.store[word] = append(index.store[word], doc.ID)
				sort.Ints(index.store[word])
			}
		}
	}
}

func (index *Index) Search(query string) []int {
	return index.store[strings.ToLower(query)]
}

func splitWords(s string) []string {
	var words []string

	for _, word := range strings.Split(s, " ") {
		words = append(words, strings.ToLower(word))
	}

	return words
}

func isExist(ids []int, item int) bool {
	for _, id := range ids {
		if id == item {
			return true
		}
	}

	return false
}
