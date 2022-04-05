package moneybird

// InvoicePayment contains info on how the invoice is paid
type InvoicePayment struct {
	ID                  string `json:"id,omitempty"`
	PaymentDate         string `json:"payment_date"`
	Price               string `json:"price"`
	PriceBase           string `json:"price_base,omitempty"`
	FinancialAccountID  int64  `json:"financial_account_id,omitempty"`
	FinancialMutationID int64  `json:"financial_mutation_id,omitempty"`
}

// InvoicePaymentGateway encapsulates all /invoices related endpoints
type InvoicePaymentGateway struct {
	*Client
}

// InvoicePayment returns a new gateway instance
func (c *Client) InvoicePayment() *InvoicePaymentGateway {
	return &InvoicePaymentGateway{c}
}

// Create marks the invoice as paid in Moneybird
func (c *InvoicePaymentGateway) Create(invoice *Invoice, payment *InvoicePayment) error {
	res, err := c.execute("POST", "sales_invoices/"+invoice.ID+"/payments", &envelope{InvoicePayment: payment})
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 201:
		return nil
	}

	return res.error()
}

// Delete payment of invoice in Moneybird
func (c *InvoicePaymentGateway) Delete(invoice *Invoice, payment *InvoicePayment) error {
	res, err := c.execute("DELETE", "sales_invoices/"+invoice.ID+"/payments/"+payment.ID, nil)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		return nil
	}

	return res.error()
}
