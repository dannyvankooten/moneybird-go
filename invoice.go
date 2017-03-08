package moneybird

import (
	"encoding/json"
	"time"
)

// Invoice contains all invoice details
type Invoice struct {
	ID                string           `json:"id"`
	InvoiceID         string           `json:"invoice_id,omitempty"`
	Contact           Contact          `json:"contact"`
	InvoiceDate       string           `json:"invoice_date"`
	State             string           `json:"state,omitempty"`
	Language          string           `json:"language,omitempty"`
	Currency          string           `json:"currency,omitempty"`
	Discount          string           `json:"discount,omitempty"`
	Details           []InvoiceDetails `json:"details"`
	TotalPriceExclTax string           `json:"total_price_excl_tax,omitempty"`
	TotalPriceInclTax string           `json:"total_price_incl_tax,omitempty"`
	PricesAreInclTax  bool             `json:"prices_are_incl_tax,omitempty"`
	URL               string           `json:"url,omitempty"`
	WorkflowID        string           `json:"workflow_id,omitempty"`
	DocumentStyleID   string           `json:"document_style_id,omitempty"`
	IdentityID        string           `json:"identity_id,omitempty"`
	PaymentConditions string           `json:"payment_conditions,omitempty"`
	SentAt            time.Time        `json:"sent_at,omitempty"`
	Reference         string           `json:"reference,omitempty"`
	CreatedAt         time.Time        `json:"created_at,omitempty"`
	UpdatedAt         time.Time        `json:"updated_at,omitempty"`
}

// InvoiceDetails is a line on an invoice
type InvoiceDetails struct {
	ID                 string   `json:"id"`
	TaxRateID          string   `json:"tax_rate_id"`
	Amount             string   `json:"amount"`
	Description        string   `json:"description"`
	Price              string   `json:"price"`
	TaxReportReference []string `json:"tax_report_reference,omitempty"`
}

// InvoiceGateway encapsulates all /sales_invoices related endpoints
type InvoiceGateway struct {
	*Client
}

// Invoice returns a new InvoiceGateway instance
func (c *Client) Invoice() *InvoiceGateway {
	return &InvoiceGateway{c}
}

// All returns all invoices
func (c *InvoiceGateway) All() ([]*Invoice, error) {
	var invoices []*Invoice
	var err error

	res, err := c.execute("GET", "sales_invoices", nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		err = json.NewDecoder(res.Body).Decode(&invoices)
	}

	return invoices, err
}

// Get returns the invoice with the specified id, or nil
func (c *InvoiceGateway) Get(ID string) (*Invoice, error) {
	res, err := c.execute("GET", "sales_invoices/"+ID, nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		return res.invoice()
	}

	// TODO: return better error here.
	return nil, err
}

// Update updates the invoice in Moneybird
func (c *InvoiceGateway) Update(invoice *Invoice) (*Invoice, error) {
	res, err := c.execute("PATCH", "sales_invoices/"+invoice.ID, &envelope{Invoice: invoice})
	if err != nil {
		return invoice, err
	}

	switch res.StatusCode {
	case 200:
		return res.invoice()
	}

	return nil, err
}

// Create creates the invoice in Moneybird
func (c *InvoiceGateway) Create(invoice *Invoice) (*Invoice, error) {
	res, err := c.execute("POST", "sales_invoices", &envelope{Invoice: invoice})
	if err != nil {
		return invoice, err
	}

	switch res.StatusCode {
	case 201:
		return res.invoice()
	}

	return nil, err
}

// Delete deletes the invoice in Moneybird
func (c *InvoiceGateway) Delete(invoice *Invoice) error {
	res, err := c.execute("DELETE", "sales_invoices/"+invoice.ID, nil)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		return nil
	}

	// TODO: Return the actual error here.
	return err
}
