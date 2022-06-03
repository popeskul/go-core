package memstore

import (
	"go-search/hw11/pkg/crawler"
	"os"
	"testing"
)

const fileName = "storage.json"

var (
	store       *DB
	defaultDocs = []crawler.Document{
		{ID: 1},
		{ID: 2},
	}
)

func TestMain(m *testing.M) {
	store = New()
	store.Add(defaultDocs)

	m.Run()
}

func TestNew(t *testing.T) {
	db := New()
	if db == nil {
		t.Errorf("New() = nil, want *DB")
	}
}

func TestAdd(t *testing.T) {
	store.Add([]crawler.Document{
		{ID: 3},
	})
	if len(store.docs) != 3 {
		t.Errorf("len(db.docs) = %d, want 3", len(store.docs))
	}
}

func TestSearch(t *testing.T) {
	tests := []struct {
		ids   []int
		count int
	}{
		{[]int{1}, 1},
		{[]int{1, 2}, 2},
		{[]int{100}, 0},
	}

	for _, tt := range tests {
		docs := store.Search(tt.ids)
		if len(docs) != tt.count {
			t.Errorf("len(db.Search(%v)) = %d, want %d", tt.ids, len(docs), tt.count)
		}
	}
}

func TestRead(t *testing.T) {
	f, err := os.Open(fileName)
	if err != nil {
		t.Errorf("os.Open(%q) = %v, want nil", fileName, err)
	}
	defer f.Close()

	docs, err2 := store.Read(f)
	if err2 != nil {
		t.Errorf("db.Read(nil) = %v, want nil", err2)
	}

	if len(docs) != 3 {
		t.Errorf("len(db.docs) = %d, want 3", len(store.docs))
	}
}

func TestWrite(t *testing.T) {
	f, err := os.Create(fileName)
	if err != nil {
		t.Errorf("os.Create(%q) = %v, want nil", fileName, err)
	}
	defer f.Close()

	_, err = store.Write(f, store.docs)
	if err != nil {
		t.Errorf("db.Write(nil) = %v, want nil", err)
	}
}
