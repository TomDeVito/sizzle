package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/tomdevito/sizzle/service"
)

const SMOKE_PARENT_ROUTE = "/smoke"

type SmokeAPI struct {
	service *service.SmokeService
}

func NewSmokeAPI(service *service.SmokeService) *SmokeAPI {
	return &SmokeAPI{
		service: service,
	}
}

func (sr *SmokeAPI) LoadAPIHandler() http.Handler {
	r := chi.NewRouter()
	r.Get("/", sr.getSmokes)
	r.Post("/", sr.CreateSmoke)
	r.Route("/{smokeID}", func(r chi.Router) {
		r.Get("/", sr.getSmoke)
	})

	return r
}

func (sr *SmokeAPI) GetAPIHandlerParentRoute() string {
	return SMOKE_PARENT_ROUTE
}

func (sr *SmokeAPI) getSmoke(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%d", sr.service)))
}

func (sr *SmokeAPI) getSmokes(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("title:%d", sr.service)))
}
func (sr *SmokeAPI) CreateSmoke(w http.ResponseWriter, r *http.Request) {
	// Validate Incoming smoke

	// save smoke

	// (model) api -> (model) service --> datastore

}
