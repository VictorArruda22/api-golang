package entities

type Customer struct {
	Id    int
	Name  string
	Email string
	Phone string
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetByID(id int) ([]Customer, error)
	Create(customer Customer) ([]Customer, error)
	Update(customer Customer) ([]Customer, error)
	Delete(id int) error
}

type CustomerService interface {
	GetAll() ([]Customer, error)
	GetByID(id int) ([]Customer, error)
	Create(customer []Customer) ([]Customer, error)
	Update(customer []Customer) ([]Customer, error)
	Delete(id int) error
}
