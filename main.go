package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Привет мир!")

}

func methodHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Метод GET")
	case "POST":
		fmt.Fprintln(w, "Метод POST")
	case "PUT":
		fmt.Fprintln(w, "Метод PUT")
	case "DELETE":
		fmt.Fprintln(w, "Метод DELETE")
	default:
		fmt.Fprintln(w, "Метод не поддерживается")
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	response := map[string]string{"message": "hello world"}
	json.NewEncoder(w).Encode(response)
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

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/method", methodHandler)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/echo", echoHandler)

	http.HandleFunc("/greet", greetHandler)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
