package paypal_sdk

import (
	"log"

	"github.com/plutov/paypal/v4"
)

func GetPaypalClient(clientId, secretId string) *paypal.Client {
	// Create a client instance
	c, err := paypal.NewClient(clientId, secretId, paypal.APIBaseSandBox)

	if err != nil {
		log.Panicf("paypal client error: %v", err)
	}
	return c
}
