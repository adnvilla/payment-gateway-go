package handler

import (
	"fmt"
	"net/http"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/application/usecase"
	errorshandle "github.com/adnvilla/payment-gateway-go/src/pkg/errors_handle"
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

	input := usecase.GetRefundInput{}

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
