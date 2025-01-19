package infrastructure

import (
	"database/sql"
	"demob/src/domain"
	"errors"
)

type ProductRepositoryMySQL struct {
	DB *sql.DB
}

func NewProductRepository(db *sql.DB) domain.ProductRepository {
	return &ProductRepositoryMySQL{DB: db}
}

func (repo *ProductRepositoryMySQL) Create(product domain.Product) error {
	query := "INSERT INTO producto (nombre, precio, cantidad) VALUES (?, ?, ?)"
	_, err := repo.DB.Exec(query, product.Nombre, product.Precio, product.Cantidad)
	return err
}

func (repo *ProductRepositoryMySQL) GetAll() ([]domain.Product, error) {
	query := "SELECT id, nombre, precio, cantidad FROM producto"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product
	for rows.Next() {
		var product domain.Product
		if err := rows.Scan(&product.ID, &product.Nombre, &product.Precio, &product.Cantidad); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (repo *ProductRepositoryMySQL) GetByID(id int) (domain.Product, error) {
	query := "SELECT id, nombre, precio, cantidad FROM producto WHERE id = ?"
	row := repo.DB.QueryRow(query, id)

	var product domain.Product
	err := row.Scan(&product.ID, &product.Nombre, &product.Precio, &product.Cantidad)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Product{}, errors.New("producto no encontrado")
	}
	return product, err
}

func (repo *ProductRepositoryMySQL) Update(product domain.Product) error {
	query := "UPDATE producto SET nombre = ?, precio = ?, cantidad = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, product.Nombre, product.Precio, product.Cantidad, product.ID)
	return err
}

func (repo *ProductRepositoryMySQL) Delete(id int) error {
	query := "DELETE FROM producto WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	return err
}
