package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	//configure / create app

	r := setupRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func setupRouter() http.Handler {

	r := chi.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	// mount router from api package

	//r.Mount(smokeAPI.GetAPIHandlerParentRoute(), )

	return r
}
