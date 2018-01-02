package store

import "github.com/tomdevito/sizzle/store/sqlstore"

type SimpleStore struct {
}

func (ss *SimpleStore) Smoke() SmokeStore {
	return sqlstore.NewSmokeStore(nil)
}
