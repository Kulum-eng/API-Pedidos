package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	o_application "ModaVane/orders/application"
	o_adapters "ModaVane/orders/infraestructure/adapters"
	o_controllers "ModaVane/orders/infraestructure/http/controllers"
	o_routes "ModaVane/orders/infraestructure/http/routes"

	core "ModaVane/orders/core"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CORS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	myGin := gin.New()
	myGin.RedirectTrailingSlash = false

	myGin.Use(CORS())

	db, err := core.InitDB()
	if err != nil {
		log.Println(err)
		return
	}
	orderRepository := o_adapters.NewMySQLOrderRepository(db)
	createOrderUseCase := o_application.NewCreateOrderUseCase(orderRepository)
	getOrderUseCase := o_application.NewGetOrderUseCase(orderRepository)
	
	updateOrderUseCase := o_application.NewUpdateOrderUseCase(orderRepository)
	deleteOrderUseCase := o_application.NewDeleteOrderUseCase(orderRepository)

	createOrderController := o_controllers.NewOrderController(createOrderUseCase , getOrderUseCase, updateOrderUseCase, deleteOrderUseCase)

	myGin.Run(":8080")
}
