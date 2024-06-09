package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/interfaces/dto"
	"github.com/adnvilla/payment-gateway-go/src/pkg/dispatcher"
	errorshandle "github.com/adnvilla/payment-gateway-go/src/pkg/errors_handle"
	"github.com/adnvilla/payment-gateway-go/src/pkg/shared_domain"
	"github.com/gin-gonic/gin"
)

type CreateOrderHandler struct {
}

func NewCreateOrderHandler() CreateOrderHandler {
	return CreateOrderHandler{}
}

func (handler *CreateOrderHandler) CreateOrder(c *gin.Context) {
	var body dto.CreateOrderRequest
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(errors.New("please check the request")),
		})
		return
	}

	input := usecase.CreateOrderInput{
		Amount:       body.Amount,
		Currency:     body.Currency,
		ProviderType: shared_domain.ProviderType(body.ProviderType),
	}

	result, err := dispatcher.Send[usecase.CreateOrderInput, usecase.CreateOrderOutput](c, input)

	if err != nil {
		// Errs it will be customize with handle errors
		c.JSON(http.StatusBadRequest, errorshandle.ErrorCustomize{
			Error: fmt.Sprint(err),
		})
		return
	}

	c.JSON(http.StatusOK, result)
}
