package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/api", user)

	log.Println("Server is running on port:", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("[ERROR]: %v", err.Error())
	}
}

func user(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	slack_name := r.URL.Query().Get("slack_name")
	if slack_name != "noornee" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	data := map[string]any{
		"slack_name":      "noornee",
		"current_day":     time.Now().Format("Monday"),
		"utc_time":        time.Now().UTC(),
		"track":           "backend",
		"github_file_url": "https://github.com/noornee/hngx_stage-1/blob/main/main.go",
		"github_repo_url": "https://github.com/noornee/hngx_stage-1",
		"status_code":     200,
	}

	jsonResp, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	fmt.Fprintln(w, string(jsonResp))
}
