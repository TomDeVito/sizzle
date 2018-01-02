package sqlstore

import (
	"database/sql"

	"github.com/tomdevito/sizzle/model"
)

type ReadingStore struct {
	db *sql.DB
}

func NewReadingStore(db *sql.DB) *SmokeStore {
	return &SmokeStore{
		db: db,
	}
}

func (ss *SmokeStore) CreateReading(reading *model.Reading) (*model.Reading, error) {
	return &model.Reading{
		UUID: "123",
	}, nil
}
