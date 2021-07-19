package main

import (
	"errors"
	"fmt"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"net/http"
)

type PaymentGateway interface {
	auth(interface{}) (*Transaction, error)
	capture(interface{}) (*Transaction, error)
}

type transactionStatus string

var authorized transactionStatus = "authorized"
var captured transactionStatus = "captured"

type Transaction struct {
	status transactionStatus
	amount int64
}

type Stripe struct {
}

func (s Stripe) auth(paymentMethod interface{}) (*Transaction, error) {
	stripe.Key = "sk_test_rQfae3eAoCgwQLTlYVyz7HkZ00xAMR4aHA"

	amount := int64(1099)

	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount),
		PaymentMethod: stripe.String("pm_1JDqVPDNFpZPwqTEnFZCSdnR"),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return &Transaction{
		status: authorized,
		amount: pi.Amount,
	}, nil
}

func (s Stripe) capture(paymentMethod interface{}) (*Transaction, error) {
	stripe.Key = "sk_test_rQfae3eAoCgwQLTlYVyz7HkZ00xAMR4aHA"

	res, err := paymentintent.Confirm(
		"pi_1JDqW7DNFpZPwqTE4IcLGZbv",
		&stripe.PaymentIntentConfirmParams{
			PaymentMethod: stripe.String("pm_1JDqVPDNFpZPwqTEnFZCSdnR"),
		},
	)

	if err != nil {
		return nil, err
	}

	return &Transaction{
		status: captured,
		amount: res.Amount,
	}, nil
}

type Paypal struct {
}

func (s Paypal) auth(paymentMethod interface{}) (*Transaction, error) {
	// do something
	return &Transaction{}, nil
}

func (s Paypal) capture(paymentMethod interface{}) (*Transaction, error) {
	return &Transaction{}, nil
}

type Factory struct {
}

func (f Factory) makePaymentGateway(gateway string) (PaymentGateway, error) {

	switch gateway {
	case "stripe":
		return Stripe{}, nil

	case "paypal":
		return Paypal{}, nil
	}

	return nil, errors.New(fmt.Sprintf("payment %s method not support", gateway))
}

func authorize(w http.ResponseWriter, req *http.Request) {

	gateway, err := Factory{}.makePaymentGateway("")

	if err != nil {

		return
	}

	//
	gateway.auth(nil)

}

func capture(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/authorize/:paymentMethod", authorize)
	http.HandleFunc("/capture", capture)

	http.ListenAndServe(":8090", nil)
}
