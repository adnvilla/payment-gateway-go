package handler

import (
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/application/usecase"
	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
	"github.com/gin-gonic/gin"
)

type GetRefundHandler struct {
	usecase use_case.UseCase[usecase.GetRefundInput, usecase.GetRefundOutput]
}

func NewGetRefundHandler(usecase use_case.UseCase[usecase.GetRefundInput, usecase.GetRefundOutput]) GetRefundHandler {
	return GetRefundHandler{
		usecase: usecase,
	}
}

func (handler *GetRefundHandler) GetRefund(c *gin.Context) {

}
