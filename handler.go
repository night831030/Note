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
	row, err := h.db.Query("SELECT * FROM note")
	if err != nil {
		log.Printf("sql select error: %v", err)
		return
	}

	var notes []Note
	for row.Next() {
		note := Note{}
		if err := row.Scan(&note.ID, &note.Title, &note.Text); err != nil {
			log.Printf("sql scan is fail %v", err)
			return
		}

		notes = append(notes, note)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notes); err != nil {
		log.Printf("unable to ecode json error: %v", err)
	}
}
