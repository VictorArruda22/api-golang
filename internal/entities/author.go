package entities

type Author struct {
	ID            int
	Name          string
	Nascimento    string
	Nacionalidade string
}

type AuthorRepository interface {
	GetAll() ([]Author, error)
	GetByID(id int) ([]Author, error)
	Create(author Author) ([]Author, error)
	Update(author Author) ([]Author, error)
	Delete(id int) error
}
