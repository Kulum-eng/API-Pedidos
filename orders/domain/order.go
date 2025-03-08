package domain

type Order struct {
    ID         int      `json:"id"`
    UserID     int      `json:"user_id"`
    Products   []int    `json:"products"`
    TotalPrice float64  `json:"total_price"`
    Status     string   `json:"status"`
}
