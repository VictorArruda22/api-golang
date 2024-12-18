package repository

import "database/sql"

func NewRepoMap( db *sql.DB) *RepoMap {
	var defaultDb RepoMap

	if db != nil {
		defaultDb = RepoMap{

		}
	}

	return &defaultDb
}

type RepoMap struct {

}