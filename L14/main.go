package main

import (
	"L14/database"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func getAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"data":    r.URL.Query().Get("data"),
		"newData": r.URL.RawQuery,
		"success": "true",
		"message": "hello",
	})

}

func postAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	json.NewEncoder(w).Encode(map[string]string{
		"data":    string(body),
		"success": "true",
		"message": "hello",
	})
}

func startServer() {
	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.ConnectDB()

	allApis()
	startServer()
}
