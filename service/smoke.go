package service

import (
	"github.com/tomdevito/sizzle/model"
	"github.com/tomdevito/sizzle/store"
)

type SmokeService struct {
	store store.SmokeStore
}

func NewSmokeService(store store.SmokeStore) *SmokeService {
	return &SmokeService{
		store: store,
	}
}
func (ss *SmokeService) createSmoke(smoke *model.Smoke) (*model.Smoke, error) {
	//validate ?

	return ss.store.CreateSmoke(smoke)
}
