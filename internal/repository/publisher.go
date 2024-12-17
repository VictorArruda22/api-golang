package repository

import (
	"database/sql"
)

func NewPublisherMap(db *sql.DB) *PublisherMap {
	var defaultDb *sql.DB
	if db != nil {
		defaultDb = db
	}
	return &PublisherMap{db: defaultDb}
}

type PublisherMap struct {
	db *sql.DB
}