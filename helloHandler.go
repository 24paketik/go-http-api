package main

import (
	"encoding/json"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string{"message": "hello world"}
	json.NewEncoder(w).Encode(response)
}
