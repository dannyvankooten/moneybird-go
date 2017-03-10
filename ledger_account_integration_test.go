package moneybird

import (
	"testing"
)

func TestLedgerAccountGatewayAll(t *testing.T) {
	_, err := testClient.LedgerAccount().List()

	if err != nil {
		t.Error(err)
	}
}

func TestLedgerAccountGatewayCRUD(t *testing.T) {
	gateway := testClient.LedgerAccount()

	// create
	acct, err := gateway.Create(&LedgerAccount{
		Name:        "Server costs",
		AccountType: "expenses",
	})

	if err != nil {
		t.Fatal(err)
	}

	if acct.AccountType != "expenses" {
		t.Errorf("LedgerAccount.Create(): invalid account type %#v", acct.AccountType)
	}

	// read
	acct, err = gateway.Get(acct.ID)
	if err != nil {
		t.Error(err)
	}

	if acct.Name != "Server costs" {
		t.Errorf("LedgerAccount.Get(): invalid account name %#v", acct.Name)
	}

	// update
	acct.Name = "Support costs"
	acct, err = gateway.Update(acct)
	if err != nil {
		t.Error(err)
	}

	if acct.Name != "Support costs" {
		t.Errorf("LedgerAccount.Update(): invalid account name %#v", acct.Name)
	}

	// delete
	err = gateway.Delete(acct)
	if err != nil {
		t.Error(err)
	}
}
