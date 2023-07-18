package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GetNotes(w http.ResponseWriter, r *http.Request) {
	notes := []Note{
		{
			ID:    1,
			Title: "frist",
			Text:  "some text",
		},
		{
			ID:    2,
			Title: "second",
			Text:  "some text 2",
		},
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(notes)
	if err != nil {
		log.Printf("Get error: %v", err)
	}
}
