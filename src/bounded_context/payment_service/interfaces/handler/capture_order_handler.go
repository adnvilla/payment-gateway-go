package handler

import (
	"fmt"
	"net/http"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	errorshandle "github.com/adnvilla/payment-gateway-go/src/pkg/errors_handle"
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

	input := usecase.CaptureOrderInput{}

	result, err := handler.usecase.Handle(c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
