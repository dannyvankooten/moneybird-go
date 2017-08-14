package moneybird

import (
	"encoding/json"
	"net/http"
)

// Invoice contains all invoice details
type Invoice struct {
	ID                string            `json:"id,omitempty"`
	AdministrationID  string            `json:"administration_id,omitempty"`
	InvoiceID         string            `json:"invoice_id,omitempty"`
	Contact           Contact           `json:"contact,omitempty"`
	ContactID         string            `json:"contact_id,omitempty"`
	UpdateContact     bool              `json:"update_contact,omitempty"`
	WorkflowID        string            `json:"workflow_id,omitempty"`
	DocumentStyleID   string            `json:"document_style_id,omitempty"`
	IdentityID        string            `json:"identity_id,omitempty"`
	State             string            `json:"state,omitempty"`
	InvoiceDate       string            `json:"invoice_date,omitempty"`
	DueDate           string            `json:"due_date,omitempty"`
	PaymentConditions string            `json:"payment_conditions,omitempty"`
	Reference         string            `json:"reference,omitempty"`
	Language          string            `json:"language,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	Discount          string            `json:"discount,omitempty"`
	PaidAt            string            `json:"paid_at,omitempty"`
	SentAt            string            `json:"sent_at,omitempty"`
	CreatedAt         string            `json:"created_at,omitempty"`
	UpdatedAt         string            `json:"updated_at,omitempty"`
	Details           []*InvoiceDetails `json:"details_attributes,omitempty"`
	TotalPaid         string            `json:"total_paid,omitempty"`
	TotalUnpaid       string            `json:"total_unpaid,omitempty"`
	TotalUnpaidBase   string            `json:"total_unpaid_base,omitempty"`
	PricesAreInclTax  bool              `json:"prices_are_incl_tax,omitempty"`
	TotalPriceExclTax string            `json:"total_price_excl_tax,omitempty"`
	TotalPriceInclTax string            `json:"total_price_incl_tax,omitempty"`
	URL               string            `json:"url,omitempty"`
	Notes             []*InvoiceNote    `json:"notes,omitempty"`
}

// InvoiceDetails is a line on an invoice
type InvoiceDetails struct {
	ID              string `json:"id,omitempty"`
	Description     string `json:"description"`
	Price           string `json:"price"`
	Period          string `json:"period,omitempty"`
	Amount          string `json:"amount,omitempty"`
	TaxRateID       string `json:"tax_rate_id,omitempty"`
	LedgerAccountID int    `json:"ledger_account_id,omitempty"`
	ProductID       int    `json:"product_id,omitempty"`
	RowOrder        int    `json:"row_order,omitempty"`
}

// InvoiceGateway encapsulates all /sales_invoices related endpoints
type InvoiceGateway struct {
	*Client
}

// Invoice returns a new InvoiceGateway instance
func (c *Client) Invoice() *InvoiceGateway {
	return &InvoiceGateway{c}
}

// List returns all invoices
func (c *InvoiceGateway) List() ([]*Invoice, error) {
	var invoices []*Invoice
	var err error

	res, err := c.execute("GET", "sales_invoices", nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		err = json.NewDecoder(res.Body).Decode(&invoices)
		return invoices, err
	}

	return nil, res.error()
}

// Get returns the invoice with the specified id, or nil
func (c *InvoiceGateway) Get(ID string) (*Invoice, error) {
	res, err := c.execute("GET", "sales_invoices/"+ID, nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		return res.invoice()
	}

	return nil, res.error()
}

// Update updates the invoice in Moneybird
func (c *InvoiceGateway) Update(invoice *Invoice) (*Invoice, error) {
	res, err := c.execute("PATCH", "sales_invoices/"+invoice.ID, &envelope{Invoice: invoice})
	if err != nil {
		return invoice, err
	}

	switch res.StatusCode {
	case http.StatusOK:
		return res.invoice()
	}

	return nil, res.error()
}

// Create creates the invoice in Moneybird
func (c *InvoiceGateway) Create(invoice *Invoice) (*Invoice, error) {
	res, err := c.execute("POST", "sales_invoices", &envelope{Invoice: invoice})
	if err != nil {
		return invoice, err
	}

	switch res.StatusCode {
	case http.StatusCreated:
		return res.invoice()
	}

	return nil, res.error()
}

// Delete deletes the invoice in Moneybird
func (c *InvoiceGateway) Delete(invoice *Invoice) error {
	res, err := c.execute("DELETE", "sales_invoices/"+invoice.ID, nil)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		return nil
	}

	return res.error()
}
