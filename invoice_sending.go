package moneybird

// InvoiceSending contains info on how the invoice is sent
type InvoiceSending struct {
	DeliveryMethod   string `json:"delivery_method"`
	EmailAddress     string `json:"email_address,omitempty"`
	EmailMessage     string `json:"email_message,omitempty"`
	SendingScheduled bool   `json:"sending_scheduled,omitempty"`
	DeliverUBL       bool   `json:"deliver_ubl,omitempty"`
	Mergeable        bool   `json:"mergeable,omitempty"`
	InvoiceDate      string `json:"invoice_date,omitempty"`
}

// InvoiceSendingGateway encapsulates all /invoices related endpoints
type InvoiceSendingGateway struct {
	*Client
}

// InvoiceSending returns a new gateway instance
func (c *Client) InvoiceSending() *InvoiceSendingGateway {
	return &InvoiceSendingGateway{c}
}

// Create marks the invoice as sent in Moneybird
func (c *InvoiceSendingGateway) Create(invoice *Invoice, sending *InvoiceSending) error {
	res, err := c.execute("PATCH", "sales_invoices/"+invoice.ID+"/send_invoice", &envelope{InvoiceSending: sending})
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 200:
		return nil
	}

	return err
}
