package main

import (
    "fmt"
    "net/http"
)

// Обработчик для главной страницы
func handleRequest(w http.ResponseWriter, r *http.Request) {
    // Ответ на запрос
    fmt.Fprintf(w, "<html><body><h1>Hello, World!</h1></body></html>")
}

func main() {
    // Определение маршрута и обработчика
    http.HandleFunc("/", handleRequest)

    // Запуск сервера
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Error starting server:", err)
    }
}