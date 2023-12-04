package controllers

import (
	"net/http"
	"strconv"

	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/dto"

	service "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts()

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, productsDto)
}
func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	productDto, err := service.ProductService.GetProductById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProductsBYpalabra(c *gin.Context) {

	log.Debug("Product id to load: " + c.Param("clave"))
	clave := c.Param("clave")

	productsDto, er := service.ProductService.GetProductsByPalabrasClaves(clave)

	if er != nil {
		c.JSON(er.Status(), er)
		return
	}
	c.JSON(http.StatusOK, productsDto)

}

func GetProductsByCategory(c *gin.Context) {
	log.Debug("category id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))

	productsDto, err := service.ProductService.GetProductsByCategory(id)

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		log.Error(err.Error())
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func GetCategories(c *gin.Context) {
	var categories []dto.CategoryDto
	categories, err := service.ProductService.GetCategories()

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, categories)
}
