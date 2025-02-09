package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) < 3 {
		http.Error(w, "Некорректный запрос", http.StatusBadRequest)
		return
	}

	userID := parts[2]

	response := map[string]string{
		"message": fmt.Sprintf("Информация о пользователе с ID = %s", userID),
	}
	json.NewEncoder(w).Encode(response)
}
