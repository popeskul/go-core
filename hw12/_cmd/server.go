package main

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := mux.NewRouter()
	endpoints(mux)

	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}

func endpoints(r *mux.Router) {
	r.HandleFunc("/{name}", handler).Methods(http.MethodGet)
}

func handler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := template.New("main")

	t, err := t.Parse(`<html><body/><h1>Hello {{.}}</h1></body></html>`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	t.Execute(w, vars["name"])
}
