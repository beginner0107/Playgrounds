package repository

import (
	"database/sql"
	"demo-scrapping/config"
	"strings"
)

type repository struct {
	cfg *config.Config
	db  *sql.DB
}

type RepositoryImpl interface {
}

const (
	admin = "Scrapping.Admin"
)

func NewRepository(cfg *config.Config) (RepositoryImpl, error) {
	r := &repository{cfg: cfg}

	dbCfg := cfg.DB

	var err error

	if r.db, err = sql.Open(dbCfg.Database, dbCfg.URL); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

func (r *repository) Add(url, cardSelector, innerSelector string, tag []string) error {
	_, err := r.db.Exec("INSERT INTO Scrapping.Admin(url, tag, cardSelector, innerSelector VALUES(?, ?, ?, ?))",
		url, strings.Join(tag, " "), cardSelector, innerSelector)
	return err
}

func (r *repository) View(url string) (interface{}, error) {}

func (r *repository) ViewAll() ([]interface{}, error) {}

func (r *repository) Update(url, cardSelector, innerSelector string, tag []string) error {
	q := query([]string{"UPDATE", admin, "SET tag = ? cardSelector = ? innerSelector = ? WHERE url = ?"})
	_, err := r.db.Exec(q, strings.Join(tag, " "), cardSelector, innerSelector, url)
	return err
}

func (r *repository) Delete(url string) error {
	q := query([]string{"DELETE", admin, "WHERE url = ?"})
	_, err := r.db.Exec(q, url)
	return err
}

func query(qs []string) string {
	return strings.Join(qs, " ") + ";"
}
