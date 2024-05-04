package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_HOST", "1")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PORT", "1")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_USER", "q")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD", "s")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_DBNAME", "t")

	gin.SetMode(gin.TestMode)

	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestInitFail(t *testing.T) {
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_HOST", "")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PORT", "sda")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_USER", "")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD", "")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_DBNAME", "")

	assert.Panics(t, func() { initializeEndpoints(nil) })
}
