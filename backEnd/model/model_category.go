package model

type Category struct {
	Id   int    `gorm:"primaryKey; int; not null; unsigned; auto_increment"`
	Name string `gorm:"type:varchar(350); not null"`
}
