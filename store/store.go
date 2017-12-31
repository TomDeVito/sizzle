package store

import "github.com/tomdevito/sizzle/model"

// Any datastore will need to satisfy this store
type Store interface {
	Smoke() SmokeStore
}

//  Store for main smoke
type SmokeStore interface {
	GetSmokeByID(int64) (*model.Smoke, error)
	CreateSmoke(*model.Smoke) (*model.Smoke, error)
}
