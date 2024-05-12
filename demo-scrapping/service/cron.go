package service

import (
	"demo-scrapping/config"
	"demo-scrapping/repository"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
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
		j.scrapping()
		fmt.Println()
	})
	c.Start()
	defer c.Stop()

	select {}
}

func (j *cronJob) scrapping() error {
	log.Println("five second job executed from mysql for Scrapping")

	if allResult, err := j.repository.ViewAll(); err != nil {
		return err
	} else if len(allResult) == 0 {
		return errors.New("all Result zero")
	} else {
		for _, r := range allResult {
			log.Println("Try Scrapping URL : %s", r.URL)
			log.Println("Try Scrapping CardSelect : %s", r.CardSelector)
			log.Println("Try Scrapping Tag : %s", r.Tag)

			fmt.Println()
			j.scrappingHTML(r.URL, r.CardSelector, r.InnerSelector, strings.Split(r.Tag, " "))
		}
		return nil
	}
}

func (j *cronJob) scrappingHTML(url, cardSelector, innerSelect string, tag []string) {
	client := http.Client{Timeout: time.Second * 3}
	if request, err := http.NewRequest("GET", url, nil); err != nil {
		log.Println("Failed To Make Request", "err", err)
	} else {
		request.Header.Set("User-Agent", "M")

		if response, err := client.Do(request); err != nil {
			log.Println("Failed To Call GET API", "err", err)
		} else {
			defer response.Body.Close()

			if doc, err := goquery.NewDocumentFromReader(response.Body); err != nil {
				log.Println("Failed To Read response", "err", err)
			} else {
				fmt.Println(doc.Html())
			}

			fmt.Print(response.Body)
		}
	}
}
