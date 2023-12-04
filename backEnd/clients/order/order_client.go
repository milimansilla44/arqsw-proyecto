package order

import (
	"time"

	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetOrderById(id int) model.OrderTable {
	var order model.OrderTable

	Db.Where("id = ?", id).First(&order)
	log.Debug("Order: ", order)

	return order
}

func GetOrders() model.Orders {
	var orders model.Orders
	Db.Find(&orders)

	log.Debug("Orders: ", orders)

	return orders
}
func UpdateMontoFinal(monto float32, id_Order int) {
	result := Db.Model(&model.OrderTable{}).Where("id = ?", id_Order).Update("total_price", monto)

	if result.Error != nil {
		log.Error("Order no encontrada")
	}

	return
}

func InsertOrder(order model.OrderTable) model.OrderTable {
	var ahora time.Time
	ahora = time.Now()
	order.Date = ahora.Format(time.RFC1123)

	result := Db.Create(&order)

	if result.Error != nil {
		log.Error("ERROR al crear la orden")
	}
	log.Debug("Order Created: ", order.Id)

	println("\n\n DATE: " + order.Date + "\n\n")
	return order
}

func InsertOrderDetail(order_detail model.OrderDetail) model.OrderDetail {
	result := Db.Create(&order_detail)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Detail Created: ", order_detail.Id)
	return order_detail
}
func DeleteOrderById(id int) {
	var order model.OrderTable
	result := Db.Where("id = ?", id).Delete(&order)

	if result.Error != nil {
		log.Error("Orden no encontrada")
		return
	}

	log.Debug("order DELETED")
	return

}
func DeleteDetailsByOrderId(orderId int) {
	var details model.OrderDetails

	result := Db.Where("order_id", orderId).Delete(&details)

	if result.Error != nil {
		log.Error("error al eliminar los detalles")
		return
	}

	log.Debug("DELETED details:", details)
	return

}
