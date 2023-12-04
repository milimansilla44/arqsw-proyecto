package orderDetailController

import (
	"net/http"
	"strconv"

	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/dto"

	service "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/services"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetOrderDetailById(c *gin.Context) {
	log.Debug("OrderDetail id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var orderDetailDto dto.OrderDetailDto

	orderDetailDto, err := service.OrderDetailService.GetOrderDetailById(id)

	if err != nil {
		c.JSON(err.Status(), err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, orderDetailDto)
}

func GetOrderDetails(c *gin.Context) {
	var orderDetailsDto dto.OrderDetailsDto
	orderDetailsDto, err := service.OrderDetailService.GetOrderDetails() //delega

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, orderDetailsDto)
}

func InsertOrderDetail(c *gin.Context) {
	var orderDetailDto dto.OrderDetailDto
	err := c.BindJSON(&orderDetailDto)

	// error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	orderDetailDto, er := service.OrderDetailService.InsertOrderDetail(orderDetailDto)
	// error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, orderDetailDto)
}
