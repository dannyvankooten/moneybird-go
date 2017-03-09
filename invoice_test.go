package moneybird

import (
	"testing"
	"time"
)

func TestInvoiceGatewayCreateAndDelete(t *testing.T) {
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

	// TODO: delete invoice
	// Moneybird processes a queued background task after an invoice is created, so this needs a few minutes delay.
	// err = testClient.Invoice().Delete(invoice)
	// if err != nil {
	// 	t.Error(err)
	// }

}
