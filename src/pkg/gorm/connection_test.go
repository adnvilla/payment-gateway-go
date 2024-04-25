package gorm

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeConnection(t *testing.T) {

	os.Setenv("PAYMENT_GATEWAY_POSTGRES_HOST", "1")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_USER", "q")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_PASSWORD", "s")
	os.Setenv("PAYMENT_GATEWAY_POSTGRES_DBNAME", "t")
	os.Setenv("PAYMENT_GATEWAY_ENV", "i")

	UseDefaultConnection()

	con := GetConnection()

	assert.NotNil(t, con)
}
