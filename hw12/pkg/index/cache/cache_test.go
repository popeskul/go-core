package cache

import (
	"go-search/hw12/pkg/crawler"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestNew(t *testing.T) {
	index := New()
	if index == nil {
		t.Errorf("New() returns nil")
	}
}

func TestAdd(t *testing.T) {
	index := New()
	docs := []crawler.Document{
		{ID: 1, Title: "title1"},
		{ID: 2, Title: "title2"},
	}
	index.Add(docs)
	if len(index.store) != 2 {
		t.Errorf("Add() failed")
	}
}

func TestSearch(t *testing.T) {
	index := New()
	docs := []crawler.Document{
		{ID: 1, Title: "title1"},
		{ID: 2, Title: "title2"},
	}
	index.Add(docs)
	if len(index.Search("title1")) != 1 {
		t.Errorf("Search() failed")
	}
}

func Test_splitWords(t *testing.T) {
	s := "title1"
	words := splitWords(s)
	if len(words) != 1 {
		t.Errorf("splitWords() failed")
	}
}

func Test_isExist(t *testing.T) {
	ids := []int{1, 2, 3}
	item := 1
	if !isExist(ids, item) {
		t.Errorf("isExist() failed")
	}
}
