package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestGetNotes(t *testing.T) {
	rqt, err := http.NewRequest("GET", "note/", nil)
	if err != nil {
		t.Errorf("test request fail: %v", err)
	}

	rpt := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("note", GetNotes)
}
