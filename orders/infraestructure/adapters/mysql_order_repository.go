package adapters

import (
	"database/sql"
	"errors"

	"ModaVane/orders/domain"
)

type MySQLOrderRepository struct {
	DB *sql.DB
}

func NewMySQLOrderRepository(db *sql.DB) *MySQLOrderRepository {
	return &MySQLOrderRepository{DB: db}
}

func (repo *MySQLOrderRepository) CreateOrder(order domain.Order) (int, error) {
	res, err := repo.DB.Exec(
		"INSERT INTO orders (user_id, products, total_price, status) VALUES (?, ?, ?, ?)",
		order.UserID, order.Products, order.TotalPrice, order.Status,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (repo *MySQLOrderRepository) GetOrderByID(id int) (*domain.Order, error) {
	var order domain.Order
	err := repo.DB.QueryRow(
		"SELECT id, user_id, products, total_price, status FROM orders WHERE id = ?",
		id,
	).Scan(&order.ID, &order.UserID, &order.Products, &order.TotalPrice, &order.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &order, nil
}

func (repo *MySQLOrderRepository) GetAllOrders() ([]domain.Order, error) {
	rows, err := repo.DB.Query("SELECT id, user_id, products, total_price, status FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []domain.Order{}
	for rows.Next() {
		var o domain.Order
		if err := rows.Scan(&o.ID, &o.UserID, &o.Products, &o.TotalPrice, &o.Status); err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

func (repo *MySQLOrderRepository) UpdateOrder(order domain.Order) error {
	_, err := repo.DB.Exec(
		"UPDATE orders SET user_id=?, products=?, total_price=?, status=? WHERE id=?",
		order.UserID, order.Products, order.TotalPrice, order.Status, order.ID,
	)
	return err
}

func (repo *MySQLOrderRepository) DeleteOrder(id int) error {
	res, err := repo.DB.Exec("DELETE FROM orders WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no se eliminó ningún registro")
	}

	return nil
}
