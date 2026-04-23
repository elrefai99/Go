package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	http.ListenAndServe(":8080", nil)
}

func main() {
	allApis()
	startServer()
}
