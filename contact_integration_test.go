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
		t.Fatalf("ContactGateway.Create: %s", err)
	}

	// 2. Scheduled Delete
	defer func() {
		err = gateway.Delete(testContact)
		if err != nil {
			t.Errorf("ContactGateway.Delete: %s", err)
		}
	}()

	// 3. Update
	testContact.FirstName = "Peter"
	testContact, err = gateway.Update(testContact)
	if err != nil {
		t.Errorf("ContactGateway.Update: %s", err)
	}

	if testContact.FirstName != "Peter" {
		t.Errorf("ContactGateway.Update: first name was not properly updated")
	}

	// 4. Get
	testContact, err = gateway.Get(testContact.ID)
	if err != nil {
		t.Errorf("ContactGateway.Get: %s", err)
	}

	if testContact.LastName != "Doe" {
		t.Errorf("ContactGateway.Get: invalid last name %#v", testContact.LastName)
	}

}

func TestContactGatewayList(t *testing.T) {
	_, err := testClient.Contact().List()
	if err != nil {
		t.Error(err)
	}
}
