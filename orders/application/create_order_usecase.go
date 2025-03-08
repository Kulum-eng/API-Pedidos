package aplication

import (
	"strconv"

	"ModaVane/orders/domain"
	"ModaVane/orders/domain/ports"

)

type CreateOrderUseCase struct {
    repo   ports.OrderRepository
    broker ports.Broker
}

func NewCreateOrderUseCase(repo ports.OrderRepository, broker ports.Broker) *CreateOrderUseCase {
    return &CreateOrderUseCase{repo: repo, broker: broker}
}

func (uc *CreateOrderUseCase) Execute(order domain.Order) (int, error) {
    // Crear la orden
    idOrder, err := uc.repo.CreateOrder(order)
    if err != nil {
        return 0, err
    }

    // Convertir el ID de la orden a string
    idOrderStr := strconv.Itoa(idOrder)

    // Publicar el ID de la orden en el broker
    err = uc.broker.Publish(idOrderStr)
    if err != nil {
        return idOrder, err
    }

    return idOrder, nil
}
