package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const SmokeParentRoute = "/smoke"

func (api *API) loadSmokeHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", api.getSmokes)
	r.Post("/", api.createSmoke)
	r.Route("/{smokeID}", func(r chi.Router) {
		r.Get("/", api.getSmoke)
	})

	return r
}

func (api *API) getSmoke(w http.ResponseWriter, r *http.Request) {
	_, _ = api.App.CreateSmoke(nil)
	w.Write([]byte(fmt.Sprintf("title:%d", 123)))
}

func (api *API) getSmokes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%d", 123)))
}

func (api *API) createSmoke(w http.ResponseWriter, r *http.Request) {
	// Validate Incoming smoke

	// save smoke

	// (model) api -> (model) service --> datastore
}
