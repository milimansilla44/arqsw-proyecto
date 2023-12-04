package dto

type OrderDetailDto struct {
	Id         int     `json:"detail_id"`
	Quantity   int     `json:"quantity"`
	Price      float32 `json:"price"`
	TotalPrice float32 `json:"total_price"`
	ProductId  int     `json:"product_id"`
	OrderId    int     `json:"order_id"`
}

type OrderDetailsDto []OrderDetailDto
