package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	var msg Message

	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(msg)
}
