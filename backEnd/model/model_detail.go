package model

type OrderDetail struct {
	Id         int     `gorm:"primaryKey; int; not null; auto_increment"`
	Quantity   int     `gorm:"type:int; not null"`
	Price      float32 `gorm:"type:float; not null"`
	TotalPrice float32 `gorm:"type:float; not null"`
	ProductId  int     `gorm:"type:int; unsigned; not null"`
	OrderId    int     `gorm:"type:int; unsigned; not null"`
}

type OrderDetails []OrderDetail
