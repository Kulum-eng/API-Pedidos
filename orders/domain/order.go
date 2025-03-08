package domain

type Order struct {
    ID         int      `json:"id"`
    UserID     int      `json:"user_id"`
    Product_id  int    `json:"product_id"`
    TotalPrice float64  `json:"total_price"`
    Status     string   `json:"status"`
}
