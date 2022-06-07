package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-search/hw13/pkg/crawler"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func (api *Api) docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(api.Docs)
}

func (api *Api) findDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	doc, err := api.store.FindById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(doc)
}

func (api *Api) createDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var d crawler.Document

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: maybe we need validation service
	if d.Title == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	// TODO: maybe we need a service to generate id
	rand.Seed(time.Now().UnixNano())
	d.ID = rand.Intn(100000)

	api.Docs = append(api.Docs, d)

	doc := api.Docs[len(api.Docs)-1]
	_ = json.NewEncoder(w).Encode(doc)
}

func (api *Api) putDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var d crawler.Document

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.mux.Lock()
	res, err := api.store.FullUpdate(id, d)
	api.mux.Unlock()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}

func (api *Api) patchDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var d crawler.Document

	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.mux.Lock()
	updatedDoc, err := api.store.PartialUpdate(id, d)
	api.mux.Unlock()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(updatedDoc)
}

func (api *Api) deleteDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	api.mux.Lock()
	err = api.store.Delete(id)
	api.mux.Unlock()

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
