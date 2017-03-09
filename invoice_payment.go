package moneybird

// InvoicePayment contains info on how the invoice is paid
type InvoicePayment struct {
	PaymentDate         string  `json:"payment_date"`
	Price               float64 `json:"price"`
	PriceBase           float64 `json:"price_base,omitempty"`
	FinancialAccountID  int64   `json:"financial_account_id,omitempty"`
	FinancialMutationID int64   `json:"financial_mutation_id,omitempty"`
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
	res, err := c.execute("PATCH", "sales_invoices/"+invoice.ID+"/register_payment", &envelope{InvoicePayment: payment})
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		return nil
	}

	return res.error()
}
