package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tomdevito/sizzle/api"
	"github.com/tomdevito/sizzle/service"
	"github.com/tomdevito/sizzle/store/sqlstore"
)

func main() {
	r := setupRouter()
	log.Fatal(http.ListenAndServe(":8080", r))

}

func setupRouter() http.Handler {

	r := chi.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	smokeAPI := api.NewSmokeAPI(
		service.NewSmokeService(
			sqlstore.NewSmokeStore(nil),
		),
	)

	r.Mount(smokeAPI.GetAPIHandlerParentRoute(), smokeAPI.LoadAPIHandler())

	return r
}
