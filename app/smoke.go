package app

import (
	"github.com/tomdevito/sizzle/model"
)

func (app *App) CreateSmoke(smoke *model.Smoke) (*model.Smoke, error) {
	//validate ?

	return app.Store.Smoke().CreateSmoke(nil)
}
