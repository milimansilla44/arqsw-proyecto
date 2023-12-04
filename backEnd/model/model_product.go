package model

type Product struct {
	Id          int     `gorm:"primaryKey; int; not null; unsigned; auto_increment"`
	Name        string  `gorm:"type:varchar(350); not null; unique"`
	Description string  `gorm:"type:varchar(350); not null"`
	Picture     string  `gorm:"type:varchar(350); not null"`
	Price       float32 `gorm:"type:float; not null"`
	Stock       int     `gorm:"type:int; unsigned; not null"`
	CategoryId  int     `gorm:"type:int; unsigned; not null"`
}

type Products []Product
