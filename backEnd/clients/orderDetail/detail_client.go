package orderDetail

import (
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetOrderDetailById(id int) model.OrderDetail {
	var orderDetail model.OrderDetail

	Db.Where("id = ?", id).First(&orderDetail)
	log.Debug("OrderDetail: ", orderDetail)

	return orderDetail
}

func GetOrderDetails() model.OrderDetails {
	var orderDetails model.OrderDetails
	Db.Find(&orderDetails)

	log.Debug("OrderDetails: ", orderDetails)

	return orderDetails
}

func InsertOrderDetail(orderDetail model.OrderDetail) model.OrderDetail {
	result := Db.Create(&orderDetail)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("OrderDetail Created: ", orderDetail.Id)
	return orderDetail
}

func InsertOrdersDetail(ordersDetail model.OrderDetails) model.OrderDetails {

	for _, orderDetail := range ordersDetail {

		result := Db.Create(&orderDetail)

		log.Debug("Order_Detail Created: ", orderDetail.Id)

		if result.Error != nil {
			log.Error("")
		}
	}

	return ordersDetail
}
