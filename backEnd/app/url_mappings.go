package app

import (
	orderDetailController "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/controllers/orderDetail"
	productController "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/controllers/product"

	userController "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/controllers/user"

	orderController "github.com/milimansilla44/proyecto-arqiSoftware/tree/master/backEnd/controllers/order"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Products Mapping
	router.GET("/product/:id", productController.GetProductById)                       // TODO OK
	router.GET("/product", productController.GetProducts)                              // TODO OK
	router.GET("/productXpalabraClave/:clave", productController.GetProductsBYpalabra) //MANDAS LA CLAVE X JSON (BODY)
	router.GET("/productsByCategory/:id", productController.GetProductsByCategory)     // OK
	router.GET("/categories", productController.GetCategories)

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById) // TODO OK
	router.GET("/user", userController.GetUsers)        // TODO OK
	router.POST("/user", userController.NewUser)        //TODO OK

	// Login Mapping
	router.POST("/login", userController.LoginUser) //TODO OK

	//Order Mapping
	router.GET("/order/:id", orderController.GetOrderById)          // TODO OK
	router.GET("/order", orderController.GetOrders)                 // TODO OK
	router.GET("/orderUser/:id", orderController.GetOrdersByUserId) //TODO Ok

	router.GET("/ordersWithDetails", orderController.GetOrdersWithDetails)       // TODO OK
	router.GET("/orderWithDetails/:id", orderController.GetOrderWithDetailsById) // TODO OK
	router.GET("/ordersWithDetails/:id", orderController.GetOrdersWithDetailsByUserId)
	router.POST("/neworder", orderController.InsertOrder)            // TODO OK
	router.DELETE("/DeleteCarrito/:id", orderController.DeleteOrder) // TODO OK

	// NOTA: puedo agregar ESTADO en orden para avisar si esta comprado o en estado carrito

	// Detail Mapping
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById) // TODO OK
	router.GET("/orderDetail", orderDetailController.GetOrderDetails)        // TODO OK
	router.POST("/neworderDetail", orderDetailController.InsertOrderDetail)  // TODO OK

	log.Info("Finishing mappings configurations")
}
