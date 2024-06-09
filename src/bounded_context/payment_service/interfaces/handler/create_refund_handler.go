package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	"github.com/adnvilla/payment-gateway-go/src/pkg/dispatcher"
	errorshandle "github.com/adnvilla/payment-gateway-go/src/pkg/errors_handle"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type CreateRefundHandler struct{}

func NewCreateRefundHandler() CreateRefundHandler {
	return CreateRefundHandler{}
}

func (handler *CreateRefundHandler) CreateRefund(c *gin.Context) {
	orderId := c.Param("id")
	if orderId == "" {
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(errors.New("please check the request")),
		})
		return
	}

	id, err := uuid.FromString(orderId)
	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	input := usecase.CreateRefundInput{
		CaptureOrderId: id,
	}

	result, err := dispatcher.Send[usecase.CreateRefundInput, usecase.CreateRefundOutput](c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, result)

}
