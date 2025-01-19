package domain

type ProductRepository interface {
	Create(product Product) error
	GetAll() ([]Product, error)
	GetByID(id int) (Product, error)
	Update(product Product) error
	Delete(id int) error
}
