package app

import (
	"github.com/rydhshlkhn/restaurant-management/codebase/factory"
)

type App struct {
	service factory.ServiceFactory
}

func New(service factory.ServiceFactory) *App {
	return &App{
		service: service,
	}
}

func (a *App) Run() {
	a.service.GetApplication().Serve()
}
