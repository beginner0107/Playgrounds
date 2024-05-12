package service

import (
	"demo-scrapping/config"
	"demo-scrapping/repository"
	"errors"
	"fmt"
	"log"

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
	db := j.repository

	c.AddFunc("*/5 * * * * *", func() {
		scrapping(db)
		fmt.Println()
	})
	c.Start()
	defer c.Stop()

	select {}
}

func scrapping(db repository.RepositoryImpl) error {
	log.Println("five second job executed from mysql for Scrapping")

	if allResult, err := db.ViewAll(); err != nil {
		return err
	} else if len(allResult) == 0 {
		return errors.New("all Result zero")
	} else {
		for _, r := range allResult {
			log.Println("Try Scrapping URL : %s", r.URL)
			log.Println("Try Scrapping CardSelect : %s", r.CardSelector)
			log.Println("Try Scrapping Tag : %s", r.Tag)
		}
		return nil
	}
}
