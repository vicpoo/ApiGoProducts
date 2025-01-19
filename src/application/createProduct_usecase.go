package application

import "demob/src/domain"

type CreateProductUseCase struct {
	Repo domain.ProductRepository
}

func (usecase *CreateProductUseCase) Execute(product domain.Product) error {
	return usecase.Repo.Create(product)
}
