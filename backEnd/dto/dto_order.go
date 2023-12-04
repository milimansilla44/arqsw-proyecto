package dto

type OrderDto struct {
	Id         int     `json:"order_id"`
	Date       string  `json:"date"`
	TotalPrice float32 `json:"total_price"`
	UserId     int     `json:"user_id"`
}

type OrdersDto []OrderDto

type OrderWithDetailsDto struct {
	Id         int             `json:"order_id"`
	Date       string          `json:"date"`
	TotalPrice float32         `json:"total_price"`
	UserId     int             `json:"user_id"`
	Details    OrderDetailsDto `json:"details"`
}

type OrdersWithDetailsDto []OrderWithDetailsDto
