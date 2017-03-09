package moneybird

import (
	"testing"
)

func TestContactGatewayCreateAndDelete(t *testing.T) {
	i := &Contact{
		Email:     "johndoe@email.com",
		FirstName: "John",
		LastName:  "Doe",
	}

	// create contact
	o, err := testClient.Contact().Create(i)
	if err != nil {
		t.Error(err)
	}

	if i.Email != o.Email {
		t.Errorf("Output field %#v does not match input field %#v.", o.Email, i.Email)
	}

	// delete contact
	err = testClient.Contact().Delete(o)
	if err != nil {
		t.Error(err)
	}
}
