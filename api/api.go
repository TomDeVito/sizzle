package api

import (
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
