package product

import (
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetProductById(id int) model.Product {
	var product model.Product
	Db.Where("id = ?", id).First(&product)
	log.Debug("Product: ", product)
	return product
}

func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products)
	log.Debug("Products: ", products)
	return products
}

func GetCategories() []model.Category {
	var categories []model.Category
	Db.Find(&categories)
	log.Debug("Categories: ", categories)
	return categories
}
