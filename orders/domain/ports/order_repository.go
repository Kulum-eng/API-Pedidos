package ports

import "ModaVane/orders/domain"

type OrderRepository interface {
    CreateOrder(order domain.Order) (int, error)
    GetOrderByID(id int) (*domain.Order, error)
    GetAllOrders() ([]domain.Order, error)
    UpdateOrder(order domain.Order) error
    DeleteOrder(id int) error
}
