package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	name := query.Get("name")
	age := query.Get("age")

	if name == "" {
		name = "Гость"
	}

	response := map[string]string{
		"message": fmt.Sprintf("Привет, %s! Ваш возвраст %s", name, age),
	}
	json.NewEncoder(w).Encode(response)
}
