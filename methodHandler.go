package main

import (
	"fmt"
	"net/http"
)

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
