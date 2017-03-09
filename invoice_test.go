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
			t.Error(err)
		}
	}

}

func TestInvoiceGatewayCreate(t *testing.T) {
	var err error
	// create contact
	contact := &Contact{
		Email:     "johndoe@email.com",
		FirstName: "John",
		LastName:  "Doe",
	}
	contact, err = testClient.Contact().Create(contact)
	if err != nil {
		t.Error(err)
	}

	// delete contact
	defer func() {
		err = testClient.Contact().Delete(contact)
		if err != nil {
			t.Error(err)
		}
	}()

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
	_, err = testClient.Invoice().Create(invoice)
	if err != nil {
		t.Error(err)
	}

}
