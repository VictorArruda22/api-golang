package utils

import "errors"

var (
	// ErrAuthorRepositoryNotFound is returned when the Author is not found
	ErrAuthorRepositoryNotFound = errors.New("repository: author not found")
	// ErrAuthorRepositoryDuplicated is returned when the Author already exists
	ErrAuthorRepositoryDuplicated = errors.New("repository: Author already exists")
	//ErrAuthorRepositoryForeignKey is returned when the foreign key constraint fails
	ErrAuthorRepositoryForeignKey = errors.New("repository: foreign key constraint fails")
	//ErrAuthorRepositoryRequest is returned when the Author is wrong
	ErrAuthorrRepositoryRequest = errors.New("repository: request Author is wrong")
	//ErrAuthorRepositoryInternalError is returned when there is an internal error
	ErrAuthorRepositoryInternalError = errors.New("repository: internal error")
	//ErrAuthorRepositoryBadField is returned when the Author has a bad field
	ErrAuthorRepositoryBadField = errors.New("repository: bad field")
	//ErrAuthorRepository is returned when the Author id is invalid
	ErrAuthorRepositoryInvalidID = errors.New("repository: invalid id")
	//ErrAuthorRepositoryNullValue is returned when the Author has a null value
	ErrAuthorRepositoryNullValue = errors.New("repository: null value")
)
