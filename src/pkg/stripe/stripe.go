package stripe

import (
	"github.com/stripe/stripe-go/v78/client"
)

func GetStripeClient(key string) *client.API {
	sc := client.New(key, nil) // the second parameter overrides the backends used if needed for mocking
	return sc
}
