package memstore

import (
	"encoding/json"
	"fmt"
	"go-search/hw13/pkg/crawler"
	"io"
	"os"
	"sort"
	"sync"
)

type DB struct {
	sync.Mutex
	docs []crawler.Document
}

func New() *DB {
	return &DB{
		docs: []crawler.Document{},
	}
}

func (db *DB) Add(docs []crawler.Document) {
	db.Lock()
	defer db.Unlock()

	db.docs = append(db.docs, docs...)

	sort.Slice(db.docs, func(i, j int) bool {
		return db.docs[i].ID < db.docs[j].ID
	})
}

func (db *DB) GetAll() []crawler.Document {
	return db.docs
}

func (db *DB) Search(ids []int) []crawler.Document {
	var results []crawler.Document

	for _, id := range ids {
		index := db.findIndex(id)
		if index == -1 {
			continue
		}

		results = append(results, db.docs[index])
	}

	return results
}

func (db *DB) Read(r io.Reader) ([]crawler.Document, error) {
	if r == nil {
		return nil, os.ErrNotExist
	}

	content, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	if len(content) == 0 {
		return nil, fmt.Errorf("empty file")
	}

	var docs []crawler.Document
	err = json.Unmarshal(content, &docs)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (db *DB) Write(w io.Writer, docs []crawler.Document) (int, error) {
	content, _ := json.Marshal(docs)
	return w.Write(content)
}

func (db *DB) FindById(id int) (crawler.Document, error) {
	index := db.findIndex(id)
	if index == -1 {
		return crawler.Document{}, fmt.Errorf("document with id %d not found", id)
	}

	return db.docs[index], nil
}

func (db *DB) FullUpdate(id int, doc crawler.Document) (crawler.Document, error) {
	index := db.findIndex(id)
	if index == -1 {
		return crawler.Document{}, fmt.Errorf("document with id %d not found", id)
	}

	db.Lock()
	db.docs[index].Body = doc.Body
	db.docs[index].Title = doc.Title
	db.docs[index].URL = doc.URL
	db.Unlock()

	// TODO: save to file

	return db.docs[index], nil
}

func (db *DB) PartialUpdate(id int, doc crawler.Document) (crawler.Document, error) {
	index := db.findIndex(id)
	if index == -1 {
		return crawler.Document{}, fmt.Errorf("document with id %d not found", id)
	}

	// TODO: maybe we need service to update only changed fields
	db.Lock()
	if doc.Body != "" {
		db.docs[index].Body = doc.Body
	}
	if doc.Title != "" {
		db.docs[index].Title = doc.Title
	}
	if doc.URL != "" {
		db.docs[index].URL = doc.URL
	}
	db.Unlock()

	// TODO: save to file

	return db.docs[index], nil
}

func (db *DB) Delete(id int) error {
	db.Lock()
	defer db.Unlock()

	index := db.findIndex(id)
	if index == -1 {
		return fmt.Errorf("document with id %d not found", id)
	}

	db.docs = append(db.docs[:index], db.docs[index+1:]...)

	return nil
}

func (db *DB) findIndex(id int) int {
	index := sort.Search(len(db.docs), func(index int) bool {
		return db.docs[index].ID >= id
	})

	if index >= len(db.docs) || db.docs[index].ID != id {
		return -1
	}

	return index
}
