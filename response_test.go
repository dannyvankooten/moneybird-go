package moneybird

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestResponseContact(t *testing.T) {
	contact := &Contact{
		ID: "id",
	}

	recorder := httptest.NewRecorder()
	json.NewEncoder(recorder.Body).Encode(contact)
	res := &Response{recorder.Result()}
	c, err := res.contact()
	if err != nil {
		t.Error(err)
	}
	if c.ID != contact.ID {
		t.Errorf("decoded output ID %#v does not match encoded input ID %#v", c.ID, contact.ID)
	}
}

func TestResponseInvoice(t *testing.T) {
	invoice := &Invoice{
		ID: "id",
	}

	recorder := httptest.NewRecorder()
	json.NewEncoder(recorder.Body).Encode(invoice)
	res := &Response{recorder.Result()}

	c, err := res.invoice()
	if err != nil {
		t.Error(err)
	}
	if c.ID != invoice.ID {
		t.Errorf("decoded output ID %#v does not match encoded input ID %#v", c.ID, invoice.ID)
	}
}

func TestResponseInvoicePayment(t *testing.T) {
	payment := &InvoicePayment{
		PaymentDate: "2017-03-08",
	}

	recorder := httptest.NewRecorder()
	json.NewEncoder(recorder.Body).Encode(payment)
	res := &Response{recorder.Result()}

	c, err := res.invoicePayment()
	if err != nil {
		t.Error(err)
	}
	if c.PaymentDate != payment.PaymentDate {
		t.Errorf("decoded output %#v does not match encoded input %#v", c.PaymentDate, payment.PaymentDate)
	}
}

func TestResponseInvoiceSending(t *testing.T) {
	sending := &InvoiceSending{
		DeliveryMethod: "Email",
	}

	recorder := httptest.NewRecorder()
	json.NewEncoder(recorder.Body).Encode(sending)
	res := &Response{recorder.Result()}

	c, err := res.invoiceSending()
	if err != nil {
		t.Error(err)
	}

	if c.DeliveryMethod != sending.DeliveryMethod {
		t.Errorf("decoded output %#v does not match encoded input %#v", c.DeliveryMethod, sending.DeliveryMethod)
	}
}
