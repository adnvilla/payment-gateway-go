package main

import (
	paymentusecases "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	paymenthandlers "github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/interfaces/handler"
	refundusecases "github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/application/usecase"
	refundhandlers "github.com/adnvilla/payment-gateway-go/src/bounded_context/refund_service/interfaces/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(cors.Default())

	routerV1 := engine.Group("/api/v1")

	initializeEndpoints(routerV1)

	if err := engine.Run(); err != nil {
		panic(err)
	}
}

func initializeEndpoints(routerV1 *gin.RouterGroup) {

	createOrderUsecase := paymentusecases.NewCreateOrderUseCase()
	captureOrderUsecase := paymentusecases.NewCaptureOrderUseCase()

	createRefundUsecase := refundusecases.NewCreateRefundUseCase()
	getRefundUseCase := refundusecases.NewGetRefundUseCase()

	paymentCreateOrderHandler := paymenthandlers.NewCreateOrderHandler(createOrderUsecase)
	paymentCaptureOrderHandler := paymenthandlers.NewCaptureOrderHandler(captureOrderUsecase)

	refundCreateRefundHandler := refundhandlers.NewCreateRefundHandler(createRefundUsecase)
	refundGetRefundHandler := refundhandlers.NewGetRefundHandler(getRefundUseCase)

	// Payments
	routerV1.POST("/payments", paymentCreateOrderHandler.CreateOrder)
	routerV1.GET("/payments/:id")
	routerV1.POST("/payments/:id/capture", paymentCaptureOrderHandler.CaptureOrder)

	// Refunds
	routerV1.POST("/refunds", refundCreateRefundHandler.CreateRefund)
	routerV1.GET("/refunds/:id", refundGetRefundHandler.GetRefund)

}
