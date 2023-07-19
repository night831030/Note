package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	noteID, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("get note id fail: %v", err)
		return
	}

	row, err := h.db.Query("SELECT * FROM note WHERE id = ?", noteID)
	if err != nil {
		log.Printf("sql select error: %v", err)
		return
	}

	note := Note{}
	for row.Next() {
		if err := row.Scan(&note.ID, &note.Title, &note.Text); err != nil {
			log.Printf("sql scan is fail %v", err)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(note); err != nil {
		log.Printf("unable to ecode json error: %v", err)
	}

}

func (h *Handler) PostNote(w http.ResponseWriter, r *http.Request) {
	note := Note{}
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		log.Printf("get note fail: %v", err)
		return
	}

	result, err := h.db.Exec("INSERT INTO Note (`list`,`log`) VALUES (`%s`,`%s`)", note.Title, note.Text)
	if err != nil {
		log.Printf("sql insert fail: %v", err)
		return
	}
	fmt.Println(result)

	w.WriteHeader(http.StatusCreated)
}
