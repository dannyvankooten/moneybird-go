package moneybird

import (
	"encoding/json"
	"net/http"
)

// Response wraps a Moneybird API response
type Response struct {
	*http.Response
}

// APIError holds data for a MoneyBird API error
type APIError struct {
	Err      string              `json:"error,omitempty"`
	Errs     map[string][]string `json:"errors,omitempty"`
	Symbolic map[string]string   `json:"symbolic,omitempty"`
}

func (e *APIError) Error() string {
	return e.Err
}

func (res *Response) error() error {
	apiErr := &APIError{}

	// try to decode into APIError struct
	err := json.NewDecoder(res.Body).Decode(apiErr)
	if err != nil {
		return err
	}

	return apiErr
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

func (res *Response) ledgerAccount() (*LedgerAccount, error) {
	var ledgerAccount *LedgerAccount
	err := json.NewDecoder(res.Body).Decode(&ledgerAccount)
	return ledgerAccount, err
}

func (res *Response) webhook() (*Webhook, error) {
	var webhook *Webhook
	err := json.NewDecoder(res.Body).Decode(&webhook)
	return webhook, err
}
