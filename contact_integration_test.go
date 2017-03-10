package moneybird

import "testing"

func TestContactGatewayCRUD(t *testing.T) {
	var err error
	var testContact *Contact
	gateway := testClient.Contact()

	// 1. Create
	testContact, err = gateway.Create(&Contact{
		Email:     "johndoe@email.com",
		FirstName: "John",
		LastName:  "Doe",
	})
	if err != nil {
		t.Fatal(err)
	}

	// 2. Scheduled Delete
	defer func() {
		err = gateway.Delete(testContact)
		if err != nil {
			t.Error(err)
		}
	}()

	// 3. Update
	testContact.FirstName = "Peter"
	testContact, err = gateway.Update(testContact)
	if err != nil {
		t.Error(err)
	}

	if testContact.FirstName != "Peter" {
		t.Errorf("Contact was not properly updated.")
	}

	// 4. Get
	testContact, err = gateway.Get(testContact.ID)
	if err != nil {
		t.Error(err)
	}

	if testContact.LastName != "Doe" {
		t.Errorf("Invalid Contact.LastName: %#v", testContact.LastName)
	}

}

func TestContactGatewayList(t *testing.T) {
	_, err := testClient.Contact().List()
	if err != nil {
		t.Error(err)
	}
}
