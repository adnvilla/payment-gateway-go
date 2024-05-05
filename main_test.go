package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/adnvilla/payment-gateway-go/src/pkg/gorm"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestInit(t *testing.T) {
	r := require.New(t)
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_HOST", "1")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PORT", "1")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_USER", "q")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD", "s")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_DBNAME", "t")

	gorm.ResetConnection()
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	r.Equal(200, w.Code)
	r.Equal("pong", w.Body.String())
}

func TestInitFail(t *testing.T) {
	r := require.New(t)
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_HOST", "")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PORT", "sda")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_USER", "")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD", "")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_DBNAME", "")

	gorm.ResetConnection()

	r.Panics(func() { initializeEndpoints(nil) })
}
