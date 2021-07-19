package main

import (
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"testing"
)

func TestCapture(t *testing.T) {
	stripe.Key = "sk_test_rQfae3eAoCgwQLTlYVyz7HkZ00xAMR4aHA"

	res, err := paymentintent.Confirm(
		"pi_1JDqW7DNFpZPwqTE4IcLGZbv",
		&stripe.PaymentIntentConfirmParams{
			PaymentMethod: stripe.String("pm_1JDqVPDNFpZPwqTEnFZCSdnR"),
		},
	)

	println(res, err)
}

func TestAuth(t *testing.T) {
	stripe.Key = "sk_test_rQfae3eAoCgwQLTlYVyz7HkZ00xAMR4aHA"

	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(1099),
		PaymentMethod: stripe.String("pm_1JDqVPDNFpZPwqTEnFZCSdnR"),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}
	pi, _ := paymentintent.New(params)

	println(pi)
}
