package handler

import (
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
	"github.com/gin-gonic/gin"
)

type CaptureOrderHandler struct {
	usecase use_case.UseCase[usecase.CaptureOrderInput, usecase.CaptureOrderOutput]
}

func NewCaptureOrderHandler(usecase use_case.UseCase[usecase.CaptureOrderInput, usecase.CaptureOrderOutput]) CaptureOrderHandler {
	return CaptureOrderHandler{
		usecase: usecase,
	}
}

func (handler *CaptureOrderHandler) CaptureOrder(c *gin.Context) {

}
