package main

import (
	"fmt"
	"net/http"
)

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

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/method", methodHandler)
	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
