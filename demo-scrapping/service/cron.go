package service

import (
	"demo-scrapping/config"
	"demo-scrapping/repository"
	"fmt"

	"github.com/robfig/cron"
)

type cronJob struct {
	cfg        *config.Config
	repository repository.RepositoryImpl
	c          *cron.Cron
}

func NewCronJob(cfg *config.Config, repository repository.RepositoryImpl) *cronJob {
	c := &cronJob{cfg: cfg, repository: repository, c: cron.New()}

	/*
		1. 메인 프로세서는 계속 진행 할 필요가 있음
		2.
	*/

	go c.runJobs()

	return c
}

func (j *cronJob) runJobs() {
	c := j.c
	//db := j.repository

	c.AddFunc("*/5 * * * * *", func() {
		fmt.Println("5초 주기로 작동합니다.")
	})
	c.Start()
	defer c.Stop()

	select {}
}
