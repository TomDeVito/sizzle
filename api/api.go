package api

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/tomdevito/sizzle/app"
)

type API struct {
	App *app.App
}

func NewAPI(app *app.App) *API {
	return &API{
		App: app,
	}
}

func (api *API) Initialize() http.Handler {
	r := chi.NewRouter()
	r.Mount(SmokeParentRoute, api.loadSmokeHandler())
	r.Mount(ReadingParentRoute, api.loadReadingHandler())

	return r
}
