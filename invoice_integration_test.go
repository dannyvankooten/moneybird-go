package moneybird

import (
	"testing"
	"time"
)

// Because Moneybird schedules a background job when you create a new invoice, this test will fail when running too soon after a previous run.
func TestInvoiceGatewayListAndDelete(t *testing.T) {
	invoices, err := testClient.Invoice().List()
	if err != nil {
		t.Error(err)
	}

	for _, invoice := range invoices {
		err := testClient.Invoice().Delete(invoice)
		if err != nil {
			// let's ignore this error for now... (see func doc)
			if err.Error() == "Sales invoice cannot be destroyed" {
				continue
			}

			t.Error(err)
		}
	}

}

func TestInvoiceGatewayCRUD(t *testing.T) {
	var err error
	// create contact
	contact := &Contact{
		Email:     "johndoe@email.com",
		FirstName: "John",
		LastName:  "Doe",
	}
	contact, err = testClient.Contact().Create(contact)
	if err != nil {
		t.Fatal(err)
	}

	// delete contact (deferred)
	defer func() {
		err = testClient.Contact().Delete(contact)
		if err != nil {
			t.Error(err)
		}
	}()

	gateway := testClient.Invoice()
	// create invoice
	invoice := &Invoice{
		ContactID:   contact.ID,
		InvoiceDate: time.Now().Format("2006-01-02"),
		Details: []*InvoiceDetails{
			&InvoiceDetails{
				Amount:      "1",
				Price:       "10.00",
				Description: "Test Service",
			},
		},
	}
	invoice, err = gateway.Create(invoice)
	if err != nil {
		t.Fatal(err) // abandon test if invoice creation fails
	}

	// update invoice
	invoice.Reference = "my-reference"
	invoice, err = gateway.Update(invoice)
	if err != nil {
		t.Error(err)
	}

	if invoice.Reference != "my-reference" {
		t.Error("Invoice.Reference was not properly updated")
	}

	//  create invoice sending (send invoice)
	err = testClient.InvoiceSending().Create(invoice, &InvoiceSending{
		DeliveryMethod: "Manual",
	})
	if err != nil {
		t.Fatal(err)
	}

	// create invoice payment (mark invoice as paid)
	err = testClient.InvoicePayment().Create(invoice, &InvoicePayment{
		Price:       invoice.TotalUnpaid,
		PaymentDate: time.Now().Format("2006-01-02"),
	})
	if err != nil {
		t.Fatal(err)
	}

	// create invoice note
	note, err := testClient.InvoiceNote().Create(invoice, &InvoiceNote{
		Note: "my note",
	})
	if err != nil {
		t.Fatal(err)
	}

	if note.Note != "my note" {
		t.Errorf("Note.Note does not match input string. Got %#v", note.Note)
	}

	// delete invoice note
	err = testClient.InvoiceNote().Delete(invoice, note)
	if err != nil {
		t.Error(err)
	}

}
