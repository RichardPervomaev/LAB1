package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func infoFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "\"Айказян Ашот Ашотович 791\"")
}
func okFunc(w http.ResponseWriter, r *http.Request) {
}
func statusFunc(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(w, "%s\n %s\n %s\n" , "Айказян Ашот Ашотович 791 ", currentTime.Format("2006-01-02 3:4:5 pm"), net.IPv4(31, 0, 0, 54))

}

func main(){
	http.HandleFunc("/ok", okFunc) // Статус страницы
	http.HandleFunc("/Info", infoFunc) // Информация о студенте
	http.HandleFunc("/status", statusFunc)
	err := http.ListenAndServe(":8080", nil) // устанавливаем порт веб-сервера
	// Если хотите использовать https, то вместо ListenAndServe используйте ListenAndServeTLS
	// err := http.ListenAndServeTLS(":8080", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
