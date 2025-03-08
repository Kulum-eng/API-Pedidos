package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	aplication "ModaVane/orders/application"
	"ModaVane/orders/domain"
	"ModaVane/orders/infraestructure/http/responses"
)

type OrderController struct {
	createOrderUseCase *aplication.CreateOrderUseCase
	getOrderUseCase    *aplication.GetOrderUseCase
	updateOrderUseCase *aplication.UpdateOrderUseCase
	deleteOrderUseCase *aplication.DeleteOrderUseCase
}

func NewOrderController(createUC *aplication.CreateOrderUseCase, getUC *aplication.GetOrderUseCase, updateUC *aplication.UpdateOrderUseCase, deleteUC *aplication.DeleteOrderUseCase) *OrderController {
	return &OrderController{
		createOrderUseCase: createUC,
		getOrderUseCase:    getUC,
		updateOrderUseCase: updateUC,
		deleteOrderUseCase: deleteUC,
	}
}

func (ctrl *OrderController) Create(ctx *gin.Context) {
	var order domain.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("los datos son inválidos", err.Error()))
		return
	}

	idOrder, err := ctrl.createOrderUseCase.Execute(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear orden", err.Error()))
		return
	}

	order.ID = idOrder
	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Orden creada exitosamente", order))
}

func (ctrl *OrderController) GetAll(ctx *gin.Context) {
	orders, err := ctrl.getOrderUseCase.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener órdenes", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Órdenes obtenidas exitosamente", orders))
}

func (ctrl *OrderController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	order, err := ctrl.getOrderUseCase.ExecuteByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener orden", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Orden obtenida exitosamente", order))
}

func (ctrl *OrderController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	var order domain.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	order.ID = id
	if err := ctrl.updateOrderUseCase.Execute(order); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al actualizar orden", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Orden actualizada exitosamente", order))
}

func (ctrl *OrderController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	if err := ctrl.deleteOrderUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al eliminar orden", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Orden eliminada exitosamente", nil))
}
