package aplication

import (
	"ModaVane/orders/domain/ports"

)

type DeleteOrderUseCase struct {
	repo ports.OrderRepository
}

func NewDeleteOrderUseCase(repo ports.OrderRepository) *DeleteOrderUseCase {
	return &DeleteOrderUseCase{repo: repo}
}

func (uc *DeleteOrderUseCase) Execute(id int) error {
	return uc.repo.DeleteOrder(id)
}
