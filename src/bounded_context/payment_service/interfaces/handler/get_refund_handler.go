package handler

import (
	"fmt"
	"net/http"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	"github.com/adnvilla/payment-gateway-go/src/pkg/dispatcher"
	errorshandle "github.com/adnvilla/payment-gateway-go/src/pkg/errors_handle"
	"github.com/gin-gonic/gin"
)

type GetRefundHandler struct{}

func NewGetRefundHandler() GetRefundHandler {
	return GetRefundHandler{}
}

func (handler *GetRefundHandler) GetRefund(c *gin.Context) {

	input := usecase.GetRefundInput{}

	result, err := dispatcher.Send[usecase.GetRefundInput, usecase.GetRefundOutput](c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, result)

}
