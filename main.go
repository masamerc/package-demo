package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"strings"
)

var port string = os.Getenv("PORT")

func main() {
	server := http.NewServeMux()

	server.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"message": "ok",
		}
		json.NewEncoder(w).Encode(data)
	})

	server.HandleFunc("POST /upper", func(w http.ResponseWriter, r *http.Request) {
		var requestData struct {
			Input string `json:"input"`
		}

		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		responseData := map[string]string{
			"result": strings.ToUpper(requestData.Input),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseData)
	})

	if port == "" {
		port = "8080"
	}

	slog.Info("Starting server on port " + port)

	http.ListenAndServe(":"+port, server)
}
