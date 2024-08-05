package main

import (
	"go-connor-ve-blog/cmd/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", handlers.Home)
	mux.HandleFunc("GET /blog", handlers.Blog)
	fileServer := http.FileServer(http.Dir("../../ui/css/"))
	mux.Handle("GET /css/", http.StripPrefix("/css", fileServer))
	imageServer := http.FileServer(http.Dir("ui/images/"))
	mux.Handle("/images/", http.StripPrefix("/images", imageServer))

	log.Print("starting server on :4000")

	err := http.ListenAndServe("localhost:4000", mux)

	log.Fatal(err)
}
