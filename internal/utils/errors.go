package utils

import "errors"

var (
	// ErrCostumerRepositoryNotFound is returned when the Costumer is not found
	ErrCostumerRepositoryNotFound = errors.New("repository: Costumer not found")
	// ErrCostumerRepositoryDuplicated is returned when the Costumer already exists
	ErrCostumerRepositoryDuplicated = errors.New("repository: Costumer already exists")
	//ErrCostumerRepositoryRequest is returned when the Costumer is wrong
	ErrCostumerRepositoryRequest = errors.New("repository: request Costumer is wrong")
	//ErrCostumerRepositoryInternalError is returned when there is an internal error
	ErrCostumerRepositoryInternalError = errors.New("repository: internal error")
	//ErrCostumerRepositoryBadField is returned when the Costumer has a bad field
	ErrCostumerRepositoryBadField = errors.New("repository: bad field")
	//ErrCostumerRepository is returned when the Costumer id is invalid
	ErrCostumerRepositoryInvalidID = errors.New("repository: invalid id")
	//ErrCostumerRepositoryNullValue is returned when the Costumer has a null value
	ErrCostumerRepositoryNullValue = errors.New("repository: null value")
)
