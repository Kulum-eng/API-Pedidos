package aplication

import (
	"strconv"

	"ModaVane/orders/domain"
	"ModaVane/orders/domain/ports"

)

type CreateOrderUseCase struct {
	repo               ports.OrderRepository
	broker             ports.Broker
	senderNotification ports.SenderNotification
}

func NewCreateOrderUseCase(repo ports.OrderRepository, broker ports.Broker, senderNotification ports.SenderNotification) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		repo:               repo,
		broker:             broker,
		senderNotification: senderNotification,
	}
}

func (uc *CreateOrderUseCase) Execute(order domain.Order) (int, error) {
	idOrder, err := uc.repo.CreateOrder(order)
	if err != nil {
		return 0, err
	}

	idOrderStr := strconv.Itoa(idOrder)

	err = uc.broker.Publish(idOrderStr)
	if err != nil {
		return idOrder, err
	}

	err = uc.senderNotification.SendNotification(map[string]interface{}{
		"event": "new-order",
		"data":  idOrderStr,
	})
	if err != nil {
		return idOrder, err
	}

	return idOrder, nil
}
