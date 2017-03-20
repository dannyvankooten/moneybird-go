package moneybird

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Response wraps a Moneybird API response
type Response struct {
	*http.Response
}

// APIError holds data for a MoneyBird API error
type APIError struct {
	response *Response
	data     map[string]interface{}
}

func (e *APIError) Error() string {
	if v, ok := e.data["error"]; ok {
		return v.(string)
	}

	return e.response.Status
}

func (res *Response) error() error {
	apiErr := &APIError{
		response: res,
	}

	//body, _ := ioutil.ReadAll(res.Body)

	// try to decode into APIError struct
	//err := json.Unmarshal(body, apiErr.data)
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

	// fixes an inconsistency with MoneyBird using `details_attributes` for outgoing JSON requests, but `details` for responses.
	body, _ := ioutil.ReadAll(res.Body)
	body = bytes.Replace(body, []byte(`"details"`), []byte(`"details_attributes"`), -1)

	err := json.Unmarshal(body, &invoice)
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

func (res *Response) note() (*InvoiceNote, error) {
	var note *InvoiceNote
	err := json.NewDecoder(res.Body).Decode(&note)
	return note, err
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
