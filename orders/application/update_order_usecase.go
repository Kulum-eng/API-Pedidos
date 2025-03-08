package aplication

import (
	"ModaVane/orders/domain"
	"ModaVane/orders/domain/ports"
)

type UpdateOrderUseCase struct {
	repo ports.OrderRepository
}

func NewUpdateOrderUseCase(repo ports.OrderRepository) *UpdateOrderUseCase {
	return &UpdateOrderUseCase{repo: repo}
}

func (uc *UpdateOrderUseCase) Execute(order domain.Order) error {
	return uc.repo.UpdateOrder(order)
}
