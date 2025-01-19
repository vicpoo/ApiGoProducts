package application

import "demob/src/domain"

type DeleteProductUseCase struct {
	Repo domain.ProductRepository
}

func (usecase *DeleteProductUseCase) Execute(id int) error {
	return usecase.Repo.Delete(id)
}
