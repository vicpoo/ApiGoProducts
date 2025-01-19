package application

import "demob/src/domain"

type UpdateProductUseCase struct {
	Repo domain.ProductRepository
}

func (usecase *UpdateProductUseCase) Execute(product domain.Product) error {
	return usecase.Repo.Update(product)
}
