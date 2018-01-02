package app

import (
	"github.com/go-kit/kit/metrics/provider"
	"github.com/tomdevito/sizzle/store"
)

type App struct {
	Store  store.Store
	Metric provider.Provider
}
