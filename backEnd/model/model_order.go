package model

type OrderTable struct {
	Id         int     `gorm:"primaryKey; int; unsigned; not null; auto_increment"`
	Date       string  `gorm:"type:varchar(50); not null"`
	TotalPrice float32 `gorm:"type:float; not null"`
	UserId     int     `gorm:"type:int; not null"`
}

type Orders []OrderTable
