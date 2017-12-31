package sqlstore

import (
	"database/sql"

	"github.com/tomdevito/sizzle/model"
)

type SmokeStore struct {
	db *sql.DB
}

func NewSmokeStore(db *sql.DB) *SmokeStore {
	return &SmokeStore{
		db: db,
	}
}

// GetSmokeByID getSmokeByID
func (ss *SmokeStore) GetSmokeByID(smokeID int64) (*model.Smoke, error) {
	return nil, nil
}

func (ss *SmokeStore) CreateSmoke(smoke *model.Smoke) (*model.Smoke, error) {
	return &model.Smoke{
		UUID: "123",
		Name: "Smoke One",
	}, nil
}
