package handler

import (
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/application/usecase"
	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
	"github.com/gin-gonic/gin"
)

type CreateRefundHandler struct {
	usecase use_case.UseCase[usecase.CreateRefundInput, usecase.CreateRefundOutput]
}

func NewCreateRefundHandler(usecase use_case.UseCase[usecase.CreateRefundInput, usecase.CreateRefundOutput]) CreateRefundHandler {
	return CreateRefundHandler{
		usecase: usecase,
	}
}

func (handler *CreateRefundHandler) CreateRefund(c *gin.Context) {

}
