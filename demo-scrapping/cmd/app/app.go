package app

import (
	"demo-scrapping/authenticator"
	"demo-scrapping/config"
	"demo-scrapping/network"
	"demo-scrapping/repository"
	"demo-scrapping/service"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg *config.Config

	network *network.Network

	authenticator authenticator.AuthenticatorImpl
	repository    repository.RepositoryImpl
	service       service.ServiceImpl

	stop chan struct{}
}

func NewApp(cfg *config.Config) *App {
	a := &App{cfg: cfg,
		stop: make(chan struct{})}

	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT)

	go func() {
		<-channel
		a.exit()
	}()

	return a
}

func (a *App) Wait() {
	<-a.stop
	os.Exit(1)
}

func (a *App) exit() {
	a.stop <- struct{}{}
}

func (a *App) Run() {
	// TODO Network 시작
}
