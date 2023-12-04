package db

import (
	orderClient "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/order"
	orderDetailClient "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/orderDetail"
	productClient "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/product"
	userClient "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/user"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// Actualiza la cadena de conexión para usar el nombre del servicio del contenedor MySQL
	dsn := "root:secret@tcp(mysql-db:3306)/nodelogin?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("Connection Failed to Open")
		log.Error(err)
		// Puedes decidir si quieres seguir ejecutando la aplicación o salir.
	}

	userClient.Db = db
	orderClient.Db = db
	productClient.Db = db
	orderDetailClient.Db = db
}

func StartDbEngine() {
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.OrderTable{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.OrderDetail{})

	log.Info("Finishing Migration Database Tables")
}
