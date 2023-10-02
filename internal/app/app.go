package app

import (
	"contentService/internal/api"
	"contentService/internal/config"
)

type App struct {
	config config.Config
	api    api.Api
}

func NewApp() *App {
	return &App{}
}

func (a *App) Initialize() {

}
