package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/method", methodHandler)

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/echo", echoHandler)

	http.HandleFunc("/greet", greetHandler)

	http.HandleFunc("/user/", userHandler)

	fmt.Println("Сервер запущен на порту 8080")
	http.ListenAndServe(":8080", nil)
}
