package repository

import (
	"database/sql"
	"errors"

	"github.com/VictorArruda22/api-golang/internal/entities"
	"github.com/VictorArruda22/api-golang/internal/utils"
	"github.com/go-sql-driver/mysql"
)

type customerRepositoryImpl struct {
	db *sql.DB
}

func NewCustomerRepository(db *sql.DB) entities.CustomerRepository {
	return &customerRepositoryImpl{db: db}
}

// GetAll implements entities.CustomerRepository
func (r *customerRepositoryImpl) GetAll() ([]entities.Customer, error) {
	rows, err := r.db.Query("SELECT id, name, email, phone FROM Customer")
	if err != nil {
		return nil, handlerError(err)
	}
	defer rows.Close()

	var customers []entities.Customer
	for rows.Next() {
		var customer entities.Customer
		if err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Phone); err != nil {
			return nil, handlerError(err)
		}

		customers = append(customers, customer)
	}
	return customers, nil
}

// GetByID implements entities.CustomerRepository
func (r *customerRepositoryImpl) GetByID(id int) ([]entities.Customer, error) {
	var customer entities.Customer
	err := r.db.QueryRow("SELECT id, name, email, phone FROM Customer WHERE id = ?", id).
		Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Phone)
	if err != nil {
		return nil, handlerError(err)
	}
	return []entities.Customer{customer}, nil
}

// Create implements entities.CustomerRepository
func (r *customerRepositoryImpl) Create(customer entities.Customer) ([]entities.Customer, error) {
	result, err := r.db.Exec("INSERT INTO Customer (name, email, phone) VALUES (?, ?, ?)",
		customer.Name, customer.Email, customer.Phone)
	if err != nil {
		return nil, handlerError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, handlerError(err)
	}
	customer.Id = int(id)
	return []entities.Customer{customer}, nil
}

// Update implements entities.CustomerRepository
func (r *customerRepositoryImpl) Update(customer entities.Customer) ([]entities.Customer, error) {
	_, err := r.db.Exec("UPDATE Customer SET name = ?, email = ?, phone = ? WHERE id = ?",
		customer.Name, customer.Email, customer.Phone, customer.Id)
	return nil, handlerError(err)
}

// Delete implements entities.Customerrepository
func (r *customerRepositoryImpl) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM Customer WHERE id = ?", id)
	return handlerError(err)
}

func handlerError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		switch mysqlErr.Number {
		case 1062:
			if mysqlErr.Message == "Duplicate entry" {
				return utils.ErrCostumerRepositoryInvalidID
			}
			err = utils.ErrCostumerRepositoryDuplicated
		case 1054:
			err = utils.ErrCostumerRepositoryBadField
		case 1048:
			err = utils.ErrCostumerRepositoryNullValue
		default:
			return utils.ErrCostumerRepositoryInternalError
		}
	}
	if errors.Is(err, sql.ErrNoRows) {
		return utils.ErrCostumerRepositoryNotFound
	}

	return err
}
