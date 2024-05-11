package app

import (
	"demo-scrapping/authenticator"
	"demo-scrapping/config"
	"demo-scrapping/network"
	"demo-scrapping/repository"
	"demo-scrapping/service"
)

type App struct {
	cfg *config.Config

	network *network.Network

	authenticator authenticator.AuthenticatorImpl
	repository    repository.RepositoryImpl
	service       service.ServiceImpl
}

func NewApp(cfg *config.Config) *App {
	a := &App{cfg: cfg}

	return a
}
