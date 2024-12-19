package utils

import "errors"

var (
	// ErrAuthorRepositoryNotFound is returned when the Author is not found
	ErrAuthorRepositoryNotFound = errors.New("repository: author not found")
	// ErrAuthorRepositoryDuplicated is returned when the Author already exists
	ErrAuthorRepositoryDuplicated = errors.New("repository: Author already exists")
	//ErrAuthorRepositoryRequest is returned when the Author is wrong
	ErrAuthorRepositoryRequest = errors.New("repository: request Author is wrong")
	//ErrAuthorRepositoryInternalError is returned when there is an internal error
	ErrAuthorRepositoryInternalError = errors.New("repository: internal error")
	//ErrAuthorRepositoryBadField is returned when the Author has a bad field
	ErrAuthorRepositoryBadField = errors.New("repository: bad field")
	//ErrAuthorRepository is returned when the Author id is invalid
	ErrAuthorRepositoryInvalidID = errors.New("repository: invalid id")
	//ErrAuthorRepositoryNullValue is returned when the Author has a null value
	ErrAuthorRepositoryNullValue = errors.New("repository: null value")
)
