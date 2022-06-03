package main

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testMux = mux.NewRouter()

func TestMain(m *testing.M) {
	testMux = mux.NewRouter()
	endpoints(testMux)
	m.Run()
}

func Test_handler(t *testing.T) {
	data := []int{}
	payload, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, "/bob", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected %d, got %d", http.StatusMethodNotAllowed, rr.Code)
	}

	//
	req, _ = http.NewRequest(http.MethodGet, "/bob", nil)
	req.Header.Set("Content-Type", "application/json")

	rr = httptest.NewRecorder()

	testMux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected %d, got %d", http.StatusOK, rr.Code)
	}

	t.Log("Response: ", rr.Body.String())
}
