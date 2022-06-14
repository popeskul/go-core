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

// @Summary Get all documents
// @Description Get all documents
// @Tags docs
// @Accept  json
// @Produce  json
// @Success 200 {array} crawler.Document
// @Router /docs [get]
func (api *Api) docs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(api.store.GetAll())
}

// @Summary Get document by id
// @Description Get document by id
// @Tags docs
// @Accept  json
// @Produce  json
// @Param id path int true "Document id"
// @Success 200 {object} crawler.Document
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Not found"
// @Router /docs/{id} [get]
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

// @Summary Create document
// @Description Create document
// @Tags docs
// @Accept  json
// @Produce  json
// @Param doc body crawler.Document true "Document"
// @Success 201 {object} crawler.Document
// @Failure 400 {object} string "Bad request"
// @Router /docs [post]
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

	api.store.Add([]crawler.Document{d})
	doc := api.store.GetAll()[len(api.store.GetAll())-1]

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(doc)
}

// @Summary Update document
// @Description Update document
// @Tags docs
// @Accept  json
// @Produce  json
// @Param id path int true "Document id"
// @Param doc body crawler.Document true "Document"
// @Success 200 {object} crawler.Document
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Not found"
// @Router /docs/{id} [put]
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

	res, err := api.store.FullUpdate(id, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(res)
}

// @Summary Patch document
// @Description Patch document
// @Tags docs
// @Accept  json
// @Produce  json
// @Param id path int true "Document id"
// @Param doc body crawler.Document true "Document"
// @Success 200 {object} crawler.Document
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Not found"
// @Router /docs/{id} [patch]
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

	updatedDoc, err := api.store.PartialUpdate(id, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_ = json.NewEncoder(w).Encode(updatedDoc)
}

// @Summary Delete document
// @Description Delete document
// @Tags docs
// @Accept  json
// @Produce  json
// @Param id path int true "Document id"
// @Success 200 {object} crawler.Document
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Not found"
// @Router /docs/{id} [delete]
func (api *Api) deleteDoc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = api.store.Delete(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
}
