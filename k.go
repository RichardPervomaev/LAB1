package main

import (
	"fmt"
	"net/http"
	"log"
)

func FIO(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ФИО 791")
}
func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HTTP-статус 200")
}
func main() {
	http.HandleFunc("/OK/", status) // Статус страницы
	http.HandleFunc("/Info/", FIO) // Информация о студенте
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	// Если хотите использовать https, то вместо ListenAndServe используйте ListenAndServeTLS
	// err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
