package aplication

import (
	"ModaVane/orders/domain"
	"ModaVane/orders/domain/ports"
)

type GetOrderUseCase struct {
	repo ports.OrderRepository
}

func NewGetOrderUseCase(repo ports.OrderRepository) *GetOrderUseCase {
	return &GetOrderUseCase{repo: repo}
}

func (uc *GetOrderUseCase) ExecuteByID(id int) (*domain.Order, error) {
	return uc.repo.GetOrderByID(id)
}

func (uc *GetOrderUseCase) ExecuteAll() ([]domain.Order, error) {
	return uc.repo.GetAllOrders()
}
