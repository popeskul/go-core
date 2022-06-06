package memstore

import (
	"errors"
	"fmt"
	"go-search/hw13/pkg/crawler"
	"os"
	"path"
	"reflect"
	"runtime"
	"testing"
)

const fileName = "storage.json"

var (
	store       *DB
	defaultDocs = []crawler.Document{
		{ID: 0},
		{ID: 1},
		{ID: 2},
	}
)

func TestMain(m *testing.M) {
	store = New()
	store.Add(defaultDocs)

	m.Run()
}

func TestDB_New(t *testing.T) {
	db := New()
	if db == nil {
		t.Errorf("New() = nil, want *DB")
	}
}

func TestDB_Add(t *testing.T) {
	store.Add([]crawler.Document{
		{ID: 3},
	})
	if len(store.docs) != 4 {
		t.Errorf("len(db.docs) = %d, want 4", len(store.docs))
	}
}

func TestDB_Search(t *testing.T) {
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

func TestDB_Read(t *testing.T) {
	// because of incorrect path to the file
	dir, err := correctCurrentLocation("../../")
	if err != nil {
		t.Errorf("os.Chdir(%q) = %v, want nil", dir, err)
	}

	var f *os.File
	defer f.Close()

	if _, err = os.Stat(path.Join(path.Dir(dir), fileName)); errors.Is(err, os.ErrNotExist) {
		f, err = os.Create(fileName)
		if err != nil {
			t.Errorf("os.Create(%q) = %v, want nil", fileName, err)
		}
	}

	// -------------------------------------------

	// happy path, file exists and has correct data
	_, err = store.Write(f, defaultDocs)
	if err != nil {
		t.Errorf("db.Write(nil) = %v, want nil", err)
	}

	f, err = os.Open(fileName)
	if err != nil {
		t.Errorf("os.Open(%q) = %v, want nil", fileName, err)
	}

	docs, err := store.Read(f)
	if err != nil {
		t.Errorf("db.Read(nil) = %v, want nil", err)
	}

	if !reflect.DeepEqual(docs, defaultDocs) {
		t.Errorf("db.Read(nil) = %v, want %v", docs, defaultDocs)
	}

	// -------------------------------------------

	// if reader is nil
	err = os.Remove(fileName)
	if err != nil {
		t.Errorf("os.Remove(%q) = %v, want nil", fileName, err)
	}

	_, err = store.Read(nil)
	if err != os.ErrNotExist {
		t.Errorf("Must be error: %v", err)
	}

	// -------------------------------------------

	// if content is empty
	write, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		t.Errorf("os.OpenFile(%q) = %v, want nil", fileName, err)
	}

	_, err = write.Write([]byte(""))
	if err != nil {
		t.Errorf("write.Write(nil) = %v, want nil", err)
	}

	f, err = os.Open(fileName)
	if err != nil {
		t.Errorf("os.Open(%q) = %v, want nil", fileName, err)
	}

	_, err = store.Read(f)
	if err.Error() != fmt.Errorf("empty file").Error() {
		t.Errorf("db.Read(nil) = %v, want nil", err)
	}

	// -------------------------------------------

	// if impossible to unmarshal
	f, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		t.Errorf("os.OpenFile(%q) = %v, want nil", fileName, err)
	}

	_, err = f.Write([]byte("/{}/"))
	if err != nil {
		t.Errorf("write.Write(nil) = %v, want nil", err)
	}

	f, err = os.Open(fileName)
	if err != nil {
		t.Errorf("os.Open(%q) = %v, want nil", fileName, err)
	}

	_, err = store.Read(f)
	if err.Error() == fmt.Errorf("unmarshal error").Error() {
		t.Errorf("db.Read(nil) = %v, want nil", err)
	}
}

func TestDB_Write(t *testing.T) {
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

func TestDB_FindById(t *testing.T) {
	tests := []struct {
		id  int
		err error
	}{
		{0, nil},
		{1, nil},
		{100, fmt.Errorf("document with id %d not found", 100)},
	}

	for _, tt := range tests {
		doc, err := store.FindById(tt.id)

		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("db.FindById(%d): got %v, want %v", tt.id, err, tt.err)
		}

		if err == nil && doc.ID != tt.id {
			t.Errorf("db.FindById(%d): got %v, want %v", tt.id, doc.ID, tt.id)
		}
	}
}

func TestDB_Update(t *testing.T) {
	tests := []struct {
		id    int
		doc   crawler.Document
		error error
	}{
		{1, crawler.Document{ID: 1, Title: "new title"}, nil},
		{1000, crawler.Document{}, fmt.Errorf("document with id %d not found", 1000)},
	}

	for _, tt := range tests {
		status := store.Update(tt.id, tt.doc)

		if !reflect.DeepEqual(status, tt.error) {
			t.Errorf("db.Update(%d, %v): got %v, want %v", tt.id, tt.doc, status, tt.error)
		}
	}
}

func TestDB_Delete(t *testing.T) {
	tests := []struct {
		id    int
		error error
	}{
		{1, nil},
		{1000, fmt.Errorf("document with id %d not found", 1000)},
	}

	for _, tt := range tests {
		status := store.Delete(tt.id)

		if !reflect.DeepEqual(status, tt.error) {
			t.Errorf("db.Delete(%d): got %v, want %v", tt.id, status, tt.error)
		}
	}
}

// for the correct path to the file
func correctCurrentLocation(p string) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), p)
	err := os.Chdir(dir)
	return dir, err
}
