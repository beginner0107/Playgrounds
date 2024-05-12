package service

import (
	"demo-scrapping/config"
	"demo-scrapping/repository"
	"demo-scrapping/types/schema"
	"log"
)

type service struct {
	cfg        *config.Config
	repository repository.RepositoryImpl
}

type ServiceImpl interface {
	Add(url, cardSelector, innerSelector string, tag []string) error
	Update(url, cardSelector, innerSelector string, tag []string) error
	Delete(url string) error
	View(url string) (*schema.Admin, error)
	ViewAll() ([]*schema.Admin, error)
}

func NewService(cfg *config.Config, repository repository.RepositoryImpl) ServiceImpl {
	s := &service{}
	return s
}

func (s *service) Add(url, cardSelector, innerSelector string, tag []string) error {
	if err := s.repository.Add(url, cardSelector, innerSelector, tag); err != nil {
		log.Println("Failed To Call Add Admin Data", "err", err)
		return err
	} else {
		return nil
	}
}

func (s *service) Update(url, cardSelector, innerSelector string, tag []string) error {
	if err := s.repository.Update(url, cardSelector, innerSelector, tag); err != nil {
		log.Println("Failed To Call Update Admin Data", "err", err)
		return err
	} else {
		return nil
	}
}

func (s *service) Delete(url string) error {
	if err := s.repository.Delete(url); err != nil {
		log.Println("Failed To Call Delete Admin Data", "err", err)
		return err
	} else {
		return nil
	}
}

func (s *service) View(url string) (*schema.Admin, error) {
	if res, err := s.repository.View(url); err != nil {
		log.Println("Failed To Call View Admin Data", "err", err)
		return nil, err
	} else {
		return res, nil
	}
}

func (s *service) ViewAll() ([]*schema.Admin, error) {
	if res, err := s.repository.ViewAll(); err != nil {
		log.Println("Failed To Call ViewAll Admin Data", "err", err)
		return nil, err
	} else {
		return res, nil
	}
}
