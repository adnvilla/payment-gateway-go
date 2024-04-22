package handler

import (
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
	"github.com/gin-gonic/gin"
)

type CreateOrderHandler struct {
	usecase use_case.UseCase[usecase.CreateOrderInput, usecase.CreateOrderOutput]
}

func NewCreateOrderHandler(usecase use_case.UseCase[usecase.CreateOrderInput, usecase.CreateOrderOutput]) CreateOrderHandler {
	return CreateOrderHandler{
		usecase: usecase,
	}
}

func (handler *CreateOrderHandler) CreateOrder(c *gin.Context) {

}
