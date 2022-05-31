package memstore

import (
	"encoding/json"
	"go-search/hw11/server/pkg/crawler"
	"io"
	"sort"
)

type DB struct {
	docs []crawler.Document
}

func New() *DB {
	return &DB{
		docs: []crawler.Document{},
	}
}

func (db *DB) Add(docs []crawler.Document) {
	db.docs = append(db.docs, docs...)

	sort.Slice(db.docs, func(i, j int) bool {
		return db.docs[i].ID < db.docs[j].ID
	})
}

func (db *DB) Search(ids []int) []crawler.Document {
	var results []crawler.Document

	for _, id := range ids {
		index := sort.Search(len(db.docs), func(index int) bool {
			return db.docs[index].ID >= id
		})

		if index >= len(db.docs) || db.docs[index].ID != id {
			return nil
		}

		results = append(results, db.docs[index])
	}

	return results
}

func (db *DB) Read(r io.Reader) ([]crawler.Document, error) {
	content, err := io.ReadAll(r)
	if err != nil {
		return nil, err
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
