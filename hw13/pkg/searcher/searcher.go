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
	Storage storage.Interface
	Index   index.Interface
	Scanner crawler.Interface
	sites   []string
	depth   int
}

const fileName = "Storage.json"

func New(
	scanner crawler.Interface,
	storage storage.Interface,
	index index.Interface,
	sites []string,
	depth int,
) *Searcher {
	return &Searcher{
		Scanner: scanner,
		Storage: storage,
		Index:   index,
		sites:   sites,
		depth:   depth,
	}
}

func (s *Searcher) Scan() ([]crawler.Document, error) {
	docs, err := s.read(fileName)
	if err == nil {
		return docs, nil
	}

	docs = s.scanUrls()
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = s.Storage.Write(f, docs)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (s *Searcher) read(fileName string) ([]crawler.Document, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	docs, err := s.Storage.Read(f)
	if err != nil {
		return nil, err
	}

	return docs, nil
}

func (s *Searcher) scanUrls() []crawler.Document {
	randInit := rand.NewSource(time.Now().UnixNano())
	var allDocs []crawler.Document

	for _, url := range s.sites {
		docs, errs := s.Scanner.Scan(url, s.depth)
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
