package model

type Login struct {
	Username string `gorm:"type:varchar(50); not null"`
	Password string `gorm:"type:varchar(255); not null"`
}

type Logins []Login
