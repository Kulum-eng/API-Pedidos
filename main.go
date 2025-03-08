package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	o_application "ModaVane/orders/application"
	core "ModaVane/orders/core"
	o_adapters "ModaVane/orders/infraestructure/adapters"
	o_controllers "ModaVane/orders/infraestructure/http/controllers"
	o_routes "ModaVane/orders/infraestructure/http/routes"

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
		log.Println("Error al conectar a la base de datos:", err)
		return
	}

	rabbitBroker := o_adapters.NewRabbitMQBroker("ec2-3-83-91-51.compute-1.amazonaws.com", 5672, "ale", "ale123")

	err = rabbitBroker.Connect()
	if err != nil {
		log.Println("Error al conectar a RabbitMQ:", err)
		return
	}

	err = rabbitBroker.InitChannel("Pedido")
	if err != nil {
		log.Println("Error al inicializar el canal de RabbitMQ:", err)
		return
	}

	orderRepository := o_adapters.NewMySQLOrderRepository(db)

	httpSenderNotification := o_adapters.NewHTTPSenderNotification("localhost", 3000)

	createOrderUseCase := o_application.NewCreateOrderUseCase(orderRepository, rabbitBroker, httpSenderNotification)
	getOrderUseCase := o_application.NewGetOrderUseCase(orderRepository)
	updateOrderUseCase := o_application.NewUpdateOrderUseCase(orderRepository)
	deleteOrderUseCase := o_application.NewDeleteOrderUseCase(orderRepository)

	createOrderController := o_controllers.NewOrderController(createOrderUseCase, getOrderUseCase, updateOrderUseCase, deleteOrderUseCase)

	o_routes.SetupOrderRoutes(myGin, createOrderController)

	if err := myGin.Run(":8083"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
