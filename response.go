package moneybird

import (
	"encoding/json"
	"net/http"
)

// Response wraps a Moneybird API response
type Response struct {
	*http.Response
}

func (res *Response) contact() (*Contact, error) {
	var contact *Contact
	err := json.NewDecoder(res.Body).Decode(&contact)
	return contact, err
}

func (res *Response) invoice() (*Invoice, error) {
	var invoice *Invoice
	err := json.NewDecoder(res.Body).Decode(&invoice)
	return invoice, err
}

func (res *Response) invoiceSending() (*InvoiceSending, error) {
	var invoiceSending *InvoiceSending
	err := json.NewDecoder(res.Body).Decode(&invoiceSending)
	return invoiceSending, err
}

func (res *Response) invoicePayment() (*InvoicePayment, error) {
	var invoicePayment *InvoicePayment
	err := json.NewDecoder(res.Body).Decode(&invoicePayment)
	return invoicePayment, err
}
