package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

const ReadingParentRoute = "/reading"

func (api *API) loadReadingHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", api.getReading)
	r.Post("/", api.createReading)
	r.Route("/{readingID}", func(r chi.Router) {
		r.Get("/", api.getReadings)
	})

	return r
}

func (api *API) getReading(w http.ResponseWriter, r *http.Request) {
	_, _ = api.App.CreateSmoke(nil)
	w.Write([]byte(fmt.Sprintf("title:%d", 123)))
}

func (api *API) getReadings(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%d", 123)))
}

func (api *API) createReading(w http.ResponseWriter, r *http.Request) {
	// Validate Incoming smoke

	// save smoke

	// (model) api -> (model) service --> datastore
}
