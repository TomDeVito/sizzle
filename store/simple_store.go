package store

import "github.com/tomdevito/sizzle/store/sqlstore"
import "database/sql"

type SimpleStore struct {
	db *sql.DB
}

func (ss *SimpleStore) Smoke() SmokeStore {
	return sqlstore.NewSmokeStore(ss.db)
}

func (ss *SimpleStore) Reading() ReadingStore {
	return sqlstore.NewReadingStore(ss.db)
}
