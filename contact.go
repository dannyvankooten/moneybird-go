package moneybird

import "encoding/json"

// Contact is a MoneyBird contact
type Contact struct {
	ID                       string `json:"id,omitempty"`
	AdministrationID         string `json:"administration_id,omitempty"`
	CompanyName              string `json:"company_name,omitempty"`
	FirstName                string `json:"firstname,omitempty"`
	LastName                 string `json:"lastname,omitempty"`
	Address1                 string `json:"address1,omitempty"`
	Address2                 string `json:"address2,omitempty"`
	ZipCode                  string `json:"zipcode,omitempty"`
	City                     string `json:"city,omitempty"`
	Country                  string `json:"country,omitempty"`
	Phone                    string `json:"phone,omitempty"`
	DeliveryMethod           string `json:"delivery_method,omitempty"`
	CustomerID               string `json:"customer_id,omitempty"`
	TaxNumber                string `json:"tax_number,omitempty"`
	ChamberOfCommerce        string `json:"chamber_of_commerce,omitempty"`
	BankAccount              string `json:"bank_account,omitempty"`
	Attention                string `json:"attention,omitempty"`
	Email                    string `json:"email"`
	EmailUBL                 bool   `json:"email_ubl,omitempty"`
	SendInvoicesToAttention  string `json:"send_invoices_to_attention,omitempty"`
	SendInvoicesToEmail      string `json:"send_invoices_to_email,omitempty"`
	SendEstimatesToAttention string `json:"send_estimates_to_attention,omitempty"`
	SendEstimatesToEmail     string `json:"send_estimates_to_email,omitempty"`
	SEPAActive               bool   `json:"sepa_active,omitempty"`
	SEPAIBAN                 string `json:"sepa_iban,omitempty"`
	SEPAIBANAccountName      string `json:"sepa_iban_account_name,omitempty"`
	SEPABIC                  string `json:"sepa_bic,omitempty"`
	SEPAMandateID            string `json:"sepa_mandate_id,omitempty"`
	SEPAMandateDate          string `json:"sepa_mandate_date,omitempty"`
	SEPASequenceType         string `json:"sepa_sequence_type,omitempty"`
	CreditCardNumber         string `json:"credit_card_number,omitempty"`
	CreditCardReference      string `json:"credit_card_reference,omitempty"`
	CreditCardType           string `json:"credit_card_type,omitempty"`
	TaxNumberValidatedAt     string `json:"tax_number_validated_at,omitempty"`
	CreatedAt                string `json:"created_at,omitempty"`
	UpdatedAt                string `json:"updated_at,omitempty"`
	SalesInvoicesURL         string `json:"sales_invoices_url,omitempty"`
}

// ContactGateway encapsulates all /contacts related endpoints
type ContactGateway struct {
	*Client
}

// Contact returns a new gateway instance
func (c *Client) Contact() *ContactGateway {
	return &ContactGateway{c}
}

// List returns all contacts in Moneybird
func (c *ContactGateway) List() ([]*Contact, error) {
	var contacts []*Contact
	var err error

	res, err := c.execute("GET", "contacts", nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		err = json.NewDecoder(res.Body).Decode(contacts)
		return contacts, err
	}

	return nil, res.error()
}

// Get returns the contact with the specified id, or nil
func (c *ContactGateway) Get(id string) (*Contact, error) {
	var err error

	res, err := c.execute("GET", "contacts/"+id, nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		return res.contact()
	}

	return nil, res.error()
}

// Create adds a contact to MoneyBird
func (c *ContactGateway) Create(contact *Contact) (*Contact, error) {
	res, err := c.execute("POST", "contacts", &envelope{Contact: contact})
	if err != nil {
		return contact, err
	}

	switch res.StatusCode {
	case 201:
		return res.contact()
	}

	return nil, res.error()
}

// Update updates an existing contact in Moneybird
func (c *ContactGateway) Update(contact *Contact) (*Contact, error) {
	res, err := c.execute("PATCH", "contacts/"+contact.ID, &envelope{Contact: contact})
	if err != nil {
		return contact, err
	}

	switch res.StatusCode {
	case 200:
		return res.contact()
	}

	return nil, res.error()
}

// Delete the given contact
func (c *ContactGateway) Delete(contact *Contact) error {
	res, err := c.execute("DELETE", "contacts/"+contact.ID, nil)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		return nil
	}

	return res.error()
}
