package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	errorshandle "github.com/adnvilla/payment-gateway-go/src/pkg/errors_handle"
	"github.com/adnvilla/payment-gateway-go/src/pkg/use_case"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
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
