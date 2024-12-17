package repository

import "database/sql"

func NewRepoMap( db *sql.DB) *RepoMap {
	if db != nil {
		defaultDb = RepoMap{
			Publishers:

		}
	}
}