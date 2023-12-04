package services

import (
	orderDetailCliente "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/orderDetail"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/dto"
	e "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/errors"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"
)

type orderDetailService struct{}

type orderDetailServiceInterface interface {
	GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError)
	GetOrderDetails() (dto.OrderDetailsDto, e.ApiError)
	InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, e.ApiError)
}

var (
	OrderDetailService orderDetailServiceInterface
)

func init() {
	OrderDetailService = &orderDetailService{}
}

func (s *orderDetailService) GetOrderDetailById(id int) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail = orderDetailCliente.GetOrderDetailById(id)
	var orderDetailDto dto.OrderDetailDto

	if orderDetail.Id == 0 {
		return orderDetailDto, e.NewBadRequestApiError("orderDetail not found")
	}
	orderDetailDto.Id = orderDetail.Id
	orderDetailDto.Quantity = orderDetail.Quantity
	orderDetailDto.Price = orderDetail.Price
	orderDetailDto.TotalPrice = orderDetail.TotalPrice
	orderDetailDto.OrderId = orderDetail.OrderId
	orderDetailDto.ProductId = orderDetail.ProductId
	return orderDetailDto, nil
}

// get array
func (s *orderDetailService) GetOrderDetails() (dto.OrderDetailsDto, e.ApiError) {

	var orderDetails model.OrderDetails = orderDetailCliente.GetOrderDetails()
	var orderDetailsDto dto.OrderDetailsDto

	for _, orderDetail := range orderDetails {
		var orderDetailDto dto.OrderDetailDto
		orderDetailDto.Id = orderDetail.Id
		orderDetailDto.Quantity = orderDetail.Quantity
		orderDetailDto.Price = orderDetail.Price
		orderDetailDto.TotalPrice = float32(orderDetail.TotalPrice)
		orderDetailDto.OrderId = orderDetail.OrderId
		orderDetailDto.ProductId = orderDetail.ProductId

		orderDetailsDto = append(orderDetailsDto, orderDetailDto)
	}

	return orderDetailsDto, nil
}

func (s *orderDetailService) InsertOrderDetail(orderDetailDto dto.OrderDetailDto) (dto.OrderDetailDto, e.ApiError) {

	var orderDetail model.OrderDetail
	orderDetail.Quantity = orderDetailDto.Quantity
	orderDetail.Price = orderDetailDto.Price
	orderDetail.TotalPrice = orderDetailDto.Price * float32(orderDetailDto.Quantity)
	orderDetail.OrderId = orderDetailDto.OrderId
	orderDetail.ProductId = orderDetailDto.ProductId

	orderDetail = orderDetailCliente.InsertOrderDetail(orderDetail)

	orderDetailDto.Id = orderDetail.Id
	orderDetailDto.TotalPrice = orderDetail.TotalPrice

	return orderDetailDto, nil
}
