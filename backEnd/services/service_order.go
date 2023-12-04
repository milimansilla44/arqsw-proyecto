package services

import (
	orderCliente "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/order"
	orderDetailCliente "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/clients/orderDetail"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/dto"
	e "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/errors"
	"github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/model"
)

type orderService struct{}

type orderServiceInterface interface {
	GetOrderById(id int) (dto.OrderDto, e.ApiError)
	GetOrders() (dto.OrdersDto, e.ApiError)
	InsertOrder(orderwithdetailsDto dto.OrderWithDetailsDto) (dto.OrderWithDetailsDto, e.ApiError)
	GetOrdersWithDetails() (dto.OrdersWithDetailsDto, e.ApiError)
	GetOrderWithDetailsById(id int) (dto.OrderWithDetailsDto, e.ApiError)
	DeleteOrder(id int) (dto.OrderWithDetailsDto, e.ApiError)
	GetOrdersWithDetailsByUserId(order_id int) (dto.OrdersWithDetailsDto, e.ApiError)
	GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError)
}

var (
	OrderService orderServiceInterface
)

func init() {
	OrderService = &orderService{}
}

func (s *orderService) GetOrderById(id int) (dto.OrderDto, e.ApiError) {

	var order model.OrderTable = orderCliente.GetOrderById(id)
	var orderDto dto.OrderDto

	if order.Id == 0 {
		return orderDto, e.NewBadRequestApiError("order not found")
	}
	orderDto.Id = order.Id
	orderDto.Date = order.Date
	orderDto.TotalPrice = order.TotalPrice
	orderDto.UserId = order.UserId

	return orderDto, nil
}

func (s *orderService) GetOrders() (dto.OrdersDto, e.ApiError) {

	var orders model.Orders = orderCliente.GetOrders()
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto

		orderDto.Id = order.Id
		orderDto.Date = order.Date
		orderDto.TotalPrice = order.TotalPrice
		orderDto.UserId = order.UserId

		ordersDto = append(ordersDto, orderDto)
	}

	return ordersDto, nil
}

func (s *orderService) GetOrdersWithDetails() (dto.OrdersWithDetailsDto, e.ApiError) {
	var orders model.Orders = orderCliente.GetOrders()
	var details model.OrderDetails = orderDetailCliente.GetOrderDetails()
	var ordersWithDetailsDto dto.OrdersWithDetailsDto

	for _, order := range orders {
		var orderWithDetailsDto dto.OrderWithDetailsDto

		orderWithDetailsDto.Id = order.Id
		orderWithDetailsDto.Date = order.Date
		orderWithDetailsDto.TotalPrice = order.TotalPrice
		orderWithDetailsDto.UserId = order.UserId

		ordersWithDetailsDto = append(ordersWithDetailsDto, orderWithDetailsDto)
	}

	for j := 0; j < len(ordersWithDetailsDto); j++ {
		var control bool = false
		var detailsssDto dto.OrderDetailsDto

		for i := 0; i < len(details); i++ {
			if details[i].OrderId == ordersWithDetailsDto[j].Id {
				control = true
				var detailDto dto.OrderDetailDto

				detailDto.Id = details[i].Id
				detailDto.Quantity = details[i].Quantity
				detailDto.Price = details[i].Price
				detailDto.TotalPrice = details[i].TotalPrice
				detailDto.ProductId = details[i].ProductId
				detailDto.OrderId = details[i].OrderId

				detailsssDto = append(detailsssDto, detailDto)
			}
		}
		if control {
			ordersWithDetailsDto[j].Details = detailsssDto
		}
	}

	return ordersWithDetailsDto, nil
}

func (s *orderService) GetOrderWithDetailsById(id int) (dto.OrderWithDetailsDto, e.ApiError) {
	var order model.OrderTable = orderCliente.GetOrderById(id)
	var details model.OrderDetails = orderDetailCliente.GetOrderDetails()

	var orderwithdetailsDto dto.OrderWithDetailsDto
	var detailsssDto dto.OrderDetailsDto

	if order.Id == 0 {
		return orderwithdetailsDto, e.NewBadRequestApiError("order not found")
	}

	orderwithdetailsDto.Id = order.Id
	orderwithdetailsDto.Date = order.Date
	orderwithdetailsDto.TotalPrice = order.TotalPrice
	orderwithdetailsDto.UserId = order.UserId

	for _, detail := range details {

		if detail.OrderId == order.Id {
			var detailDto dto.OrderDetailDto

			detailDto.Id = detail.Id
			detailDto.OrderId = detail.OrderId
			detailDto.Price = detail.Price
			detailDto.ProductId = detail.ProductId
			detailDto.Quantity = detail.Quantity
			detailDto.TotalPrice = detail.TotalPrice

			detailsssDto = append(detailsssDto, detailDto)
		}
	}
	orderwithdetailsDto.Details = detailsssDto

	return orderwithdetailsDto, nil

}

func (s *orderService) InsertOrder(orderwithdetailsDto dto.OrderWithDetailsDto) (dto.OrderWithDetailsDto, e.ApiError) {

	//recibe de order: date, user_id y de los detalles: quantity, price, product_id

	var order model.OrderTable
	var detailsssDto dto.OrderDetailsDto

	order.UserId = orderwithdetailsDto.UserId

	order = orderCliente.InsertOrder(order)
	orderwithdetailsDto.Id = order.Id
	orderwithdetailsDto.Date = order.Date

	for _, OrderDetailDto := range orderwithdetailsDto.Details {
		var detail model.OrderDetail

		detail.OrderId = order.Id
		OrderDetailDto.OrderId = order.Id
		detail.Price = 0
		detail.Quantity = OrderDetailDto.Quantity
		detail.TotalPrice = 0
		detail.ProductId = OrderDetailDto.ProductId

		detail = orderCliente.InsertOrderDetail(detail)
	}

	var detailsss model.OrderDetails = orderDetailCliente.GetOrderDetails()
	var Total_price float32 = 0

	for _, detail := range detailsss {

		if detail.OrderId == order.Id {
			var detailDto dto.OrderDetailDto

			detailDto.Id = detail.Id
			detailDto.OrderId = detail.OrderId
			detailDto.Price = detail.Price
			detailDto.ProductId = detail.ProductId
			detailDto.Quantity = detail.Quantity
			detailDto.TotalPrice = detail.TotalPrice

			detailsssDto = append(detailsssDto, detailDto)
			Total_price = Total_price + float32(detail.TotalPrice)
		}
	}

	orderCliente.UpdateMontoFinal(Total_price, orderwithdetailsDto.Id)
	orderwithdetailsDto.TotalPrice = Total_price

	orderwithdetailsDto.Details = detailsssDto
	print(detailsss)
	return orderwithdetailsDto, nil
}

func (s *orderService) DeleteOrder(id int) (dto.OrderWithDetailsDto, e.ApiError) {

	var order model.OrderTable = orderCliente.GetOrderById(id)
	var details model.OrderDetails = orderDetailCliente.GetOrderDetails()

	var orderwithdetailsDto dto.OrderWithDetailsDto
	var detailsssDto dto.OrderDetailsDto

	if order.Id == 0 {
		return orderwithdetailsDto, e.NewBadRequestApiError("order not found")
	}

	orderwithdetailsDto.Id = order.Id
	orderwithdetailsDto.Date = order.Date
	orderwithdetailsDto.TotalPrice = order.TotalPrice
	orderwithdetailsDto.UserId = order.UserId

	for _, detail := range details {

		if detail.OrderId == order.Id {
			var detailDto dto.OrderDetailDto

			detailDto.Id = detail.Id
			detailDto.OrderId = detail.OrderId
			detailDto.Price = detail.Price
			detailDto.ProductId = detail.ProductId
			detailDto.Quantity = detail.Quantity
			detailDto.TotalPrice = detail.TotalPrice

			detailsssDto = append(detailsssDto, detailDto)
		}
	}
	orderwithdetailsDto.Details = detailsssDto

	orderCliente.DeleteOrderById(id)
	orderCliente.DeleteDetailsByOrderId(id)

	return orderwithdetailsDto, nil

}

func (s *orderService) GetOrdersWithDetailsByUserId(id int) (dto.OrdersWithDetailsDto, e.ApiError) {
	var orders model.Orders = orderCliente.GetOrders()
	var details model.OrderDetails = orderDetailCliente.GetOrderDetails()
	var ordersWithDetailsDto dto.OrdersWithDetailsDto

	for _, order := range orders {
		if order.UserId == id {
			var orderWithDetailsDto dto.OrderWithDetailsDto

			orderWithDetailsDto.Id = order.Id
			orderWithDetailsDto.Date = order.Date
			orderWithDetailsDto.TotalPrice = order.TotalPrice
			orderWithDetailsDto.UserId = order.UserId

			ordersWithDetailsDto = append(ordersWithDetailsDto, orderWithDetailsDto)
		}
	}

	for j := 0; j < len(ordersWithDetailsDto); j++ {
		var control bool = false
		var detailsssDto dto.OrderDetailsDto

		for i := 0; i < len(details); i++ {
			if details[i].OrderId == ordersWithDetailsDto[j].Id {
				control = true
				var detailDto dto.OrderDetailDto

				detailDto.Id = details[i].Id
				detailDto.Quantity = details[i].Quantity
				detailDto.Price = details[i].Price
				detailDto.TotalPrice = details[i].TotalPrice
				detailDto.ProductId = details[i].ProductId
				detailDto.OrderId = details[i].OrderId

				detailsssDto = append(detailsssDto, detailDto)
			}
		}
		if control {
			ordersWithDetailsDto[j].Details = detailsssDto
		}
	}

	return ordersWithDetailsDto, nil
}

func (s *orderService) GetOrdersByUserId(id int) (dto.OrdersDto, e.ApiError) {
	var orders model.Orders = orderCliente.GetOrders()
	var ordersDto dto.OrdersDto

	for _, order := range orders {
		var orderDto dto.OrderDto

		if order.UserId == id {

			orderDto.Id = order.Id
			orderDto.Date = order.Date
			orderDto.TotalPrice = order.TotalPrice
			orderDto.UserId = order.UserId

			ordersDto = append(ordersDto, orderDto)
		}
	}

	return ordersDto, nil

}
