package main

import (
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/application/usecase"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/domain/service"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/infrastructure/provider"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/infrastructure/sql/postgresql"
	"github.com/adnvilla/payment-gateway-go/src/bounded_context/payment_service/interfaces/handler"
	"github.com/adnvilla/payment-gateway-go/src/pkg/gorm"

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

	db := gorm.GetConnection()
	factory := provider.NewGetProviderFactory()
	service := service.NewCreateOrderService(factory)
	repository := postgresql.NewOrderRepository(db)

	createOrderUsecase := usecase.NewCreateOrderUseCase(service, repository)
	captureOrderUsecase := usecase.NewCaptureOrderUseCase(service, repository)

	createRefundUsecase := usecase.NewCreateRefundUseCase()
	getRefundUseCase := usecase.NewGetRefundUseCase()

	paymentCreateOrderHandler := handler.NewCreateOrderHandler(createOrderUsecase)
	paymentCaptureOrderHandler := handler.NewCaptureOrderHandler(captureOrderUsecase)

	refundCreateRefundHandler := handler.NewCreateRefundHandler(createRefundUsecase)
	refundGetRefundHandler := handler.NewGetRefundHandler(getRefundUseCase)

	// Payments
	routerV1.POST("/payments", paymentCreateOrderHandler.CreateOrder)
	routerV1.GET("/payments/:id")
	routerV1.POST("/payments/:id/capture", paymentCaptureOrderHandler.CaptureOrder)

	// Refunds
	routerV1.POST("/refunds", refundCreateRefundHandler.CreateRefund)
	routerV1.GET("/refunds/:id", refundGetRefundHandler.GetRefund)

}
