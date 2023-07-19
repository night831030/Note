package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
)

type Note struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("env file not found %v", err)
	}

	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	connectString := fmt.Sprintf("%s:%s@tcp(localhost:3306)/notebook", username, password)

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		log.Fatalf("sql connect fail: %v\n", err)
	}
	defer db.Close()

	handler := NewHandler(db)

	router := mux.NewRouter()

	router.HandleFunc("/note", handler.GetNotes).Methods("GET")
	router.HandleFunc("/note/{id}", handler.GetNote).Methods("GET")
	router.HandleFunc("/note", handler.PostNote).Methods("POST")
	// router.HandleFunc("/note/{id}", PutNote).Methods("PUT")
	// router.HandleFunc("/note/{id}", PatchNote).Methods("PATCH")
	router.HandleFunc("/note/{id}", handler.DeleteNote).Methods("DELETE")

	http.ListenAndServe("localhost:8080", router)
}

// func check(w http.ResponseWriter, err error) {
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 		log.Printf("錯誤：%v\n", err)
// 	}
// }

// func errlog(err error) {
// 	if err != nil {
// 		log.Printf("錯誤：%v\n", err)
// 	}
// }

// func PostNote(w http.ResponseWriter, r *http.Request) {
// 	var note Note
// 	err := json.NewDecoder(r.Body).Decode(&note)
// 	check(w, err)
// 	var noteID int
// 	note.Database("POST", noteID)
// 	w.WriteHeader(http.StatusCreated)
// }

// func PutNote(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	noteID, err := strconv.Atoi(vars["id"])
// 	check(w, err)

// 	var updatenote Note
// 	err = json.NewDecoder(r.Body).Decode(&updatenote)
// 	check(w, err)

// 	updatenote.Database("PUT", noteID)
// 	w.WriteHeader(http.StatusOK)
// }

// func PatchNote(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	noteID, err := strconv.Atoi(vars["id"])
// 	check(w, err)

// 	var updatenote Note
// 	err = json.NewDecoder(r.Body).Decode(&updatenote)
// 	check(w, err)

// 	updatenote.Database("PATCH", noteID)
// 	w.WriteHeader(http.StatusOK)
// }

// func DeleteNote(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	noteID, err := strconv.Atoi(vars["id"])
// 	check(w, err)

// 	var note Note
// 	note.Database("DELETE", noteID)
// 	w.WriteHeader(http.StatusOK)
// }
