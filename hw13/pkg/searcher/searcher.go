package searcher

import (
	"fmt"
	"go-search/hw13/pkg/crawler"
	"go-search/hw13/pkg/index"
	"go-search/hw13/pkg/storage"
	"math/rand"
	"os"
	"time"
)

type Searcher struct {
	storage storage.Interface
	index   index.Interface
	scanner crawler.Interface
	sites   []string
	depth   int
}

const fileName = "storage.json"

func New(
	scanner crawler.Interface,
	storage storage.Interface,
	index index.Interface,
	sites []string,
	depth int,
) *Searcher {
	return &Searcher{
		scanner: scanner,
		storage: storage,
		index:   index,
		sites:   sites,
		depth:   depth,
	}
}

func (s *Searcher) ScanForDocuments() ([]crawler.Document, error) {
	docs, err := s.readDocumentsFromFile(fileName)
	if err == nil {
		return docs, nil
	}

	docs = s.scanUrls()
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = s.storage.Write(f, docs)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (s *Searcher) readDocumentsFromFile(fileName string) ([]crawler.Document, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	docs, err := s.storage.Read(f)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (s *Searcher) scanUrls() []crawler.Document {
	randInit := rand.NewSource(time.Now().UnixNano())
	var allDocs []crawler.Document

	for _, url := range s.sites {
		docs, errs := s.scanner.Scan(url, s.depth)
		if errs != nil {
			fmt.Println("Error scan: ", errs)
		}

		for i := range docs {
			docs[i].ID = rand.New(randInit).Intn(1000)
		}

		allDocs = append(docs, docs...)
	}

	return allDocs
}

func (s *Searcher) AddDocumentsToStorage(docs []crawler.Document) {
	s.storage.Add(docs)
}

func (s *Searcher) AddDocumentsToIndex(docs []crawler.Document) {
	s.index.Add(docs)
}

func (s *Searcher) Storage() storage.Interface {
	return s.storage
}
