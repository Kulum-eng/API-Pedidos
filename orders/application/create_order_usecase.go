package aplication

import (
	"ModaVane/orders/domain"
	"ModaVane/orders/domain/ports"
)

type CreateOrderUseCase struct {
	repo ports.OrderRepository
}

func NewCreateOrderUseCase(repo ports.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{repo: repo}
}

func (uc *CreateOrderUseCase) Execute(order domain.Order) (int, error) {
	return uc.repo.CreateOrder(order)
}
