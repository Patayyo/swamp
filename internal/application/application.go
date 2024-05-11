package application

import (
	"github.com/gorepos/usercartv2/internal/store"
)

type Application struct {
	S store.Store
}

func NewApplication(store store.Store) *Application {
	return &Application{S: store}
}

