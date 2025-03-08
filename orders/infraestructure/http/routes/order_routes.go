package routes

import (
    "github.com/gin-gonic/gin"

    "ModaVane/orders/infraestructure/http/controllers"

)

func SetupOrderRoutes(router *gin.Engine, controller *controllers.OrderController) {
    orderRoutes := router.Group("/orders")
    {
        orderRoutes.POST("/", controller.Create)
        orderRoutes.GET("/", controller.GetAll)
        orderRoutes.GET("/:id", controller.GetByID)
        orderRoutes.PUT("/:id", controller.Update)
        orderRoutes.DELETE("/:id", controller.Delete)
    }
}
