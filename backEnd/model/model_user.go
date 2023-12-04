package model

type User struct {
	Id       int    `gorm:"primaryKey; not null; auto_increment"`
	UserName string `gorm:"column:username; type:varchar(55); not null"`
	Password string `gorm:"type:varchar(255); not null"`
	City     string `gorm:"type:varchar(50); not null"`
	Street   string `gorm:"type:varchar(50); not null"`
	Number   int    `gorm:"type:int; not null"`
	Name     string `gorm:"type:varchar(55); not null"`
}

type Users []User
