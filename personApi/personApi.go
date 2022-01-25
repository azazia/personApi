package main

import (
	"personApi/handler"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)







func main() {
	http.HandleFunc("/", handler.GetHandler)
	http.HandleFunc("/insert", handler.PostHandler)

	log.Println("listening port 8000")
	log.Fatal(http.ListenAndServe(":8000",nil))
}