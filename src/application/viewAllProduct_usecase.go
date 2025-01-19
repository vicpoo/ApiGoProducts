package application

import "demob/src/domain"

type ViewAllProductsUseCase struct {
	Repo domain.ProductRepository
}

func (usecase *ViewAllProductsUseCase) Execute() ([]domain.Product, error) {
	return usecase.Repo.GetAll()
}
