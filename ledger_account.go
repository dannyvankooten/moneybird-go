package moneybird

import "encoding/json"

// LedgerAccount represends a ledger account in Moneybird
type LedgerAccount struct {
	ID               string `json:"id,omitempty"`
	AdministrationID string `json:"administration_id,omitempty"`
	Name             string `json:"name"`
	AccountType      string `json:"account_type"`
	AccountID        string `json:"account_id,omitempty"`
	ParentID         string `json:"parent_id,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

// LedgerAccountGateway encapsulates all /ledger_accounts related endpoints
type LedgerAccountGateway struct {
	*Client
}

// LedgerAccount returns a new gateway instance
func (c *Client) LedgerAccount() *LedgerAccountGateway {
	return &LedgerAccountGateway{c}
}

// List returns all ledger accounts in Moneybird
func (c *LedgerAccountGateway) List() ([]*LedgerAccount, error) {
	var ledgerAccounts []*LedgerAccount
	var err error

	res, err := c.execute("GET", "ledger_accounts", nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		err = json.NewDecoder(res.Body).Decode(&ledgerAccounts)
		return ledgerAccounts, err
	}

	return nil, res.error()
}

// Get returns the ledger account with the specified id, or nil
func (c *LedgerAccountGateway) Get(id string) (*LedgerAccount, error) {
	var err error

	res, err := c.execute("GET", "ledger_accounts/"+id, nil)
	if err != nil {
		return nil, err
	}

	// TODO: Check status code here.
	switch res.StatusCode {
	case 200:
		return res.ledgerAccount()
	}

	return nil, err
}

// Create adds a ledger account to MoneyBird
func (c *LedgerAccountGateway) Create(ledgerAccount *LedgerAccount) (*LedgerAccount, error) {
	res, err := c.execute("POST", "ledger_accounts", &envelope{LedgerAccount: ledgerAccount})
	if err != nil {
		return ledgerAccount, err
	}

	switch res.StatusCode {
	case 201:
		return res.ledgerAccount()
	}

	return nil, err
}

// Update updates an existing ledger account in Moneybird
func (c *LedgerAccountGateway) Update(ledgerAccount *LedgerAccount) (*LedgerAccount, error) {
	res, err := c.execute("PATCH", "ledger_accounts/"+ledgerAccount.ID, &envelope{LedgerAccount: ledgerAccount})
	if err != nil {
		return ledgerAccount, err
	}

	switch res.StatusCode {
	case 200:
		return res.ledgerAccount()
	}

	return nil, err
}

// Delete the given ledger account
func (c *LedgerAccountGateway) Delete(ledgerAccount *LedgerAccount) error {
	res, err := c.execute("DELETE", "ledger_accounts/"+ledgerAccount.ID, nil)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		return nil
	}

	return err
}
