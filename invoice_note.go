package moneybird

// InvoiceNote is a note for an invoice
type InvoiceNote struct {
	ID               string `json:"id,omitempty"`
	AdministrationID string `json:"administration_id,omitempty"`
	Note             string `json:"note"`
	Todo             bool   `json:"todo,omitempty"`
	AssigneeID       int    `json:"assignee_id,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

// InvoiceNoteGateway encapsulates all /invoices related endpoints
type InvoiceNoteGateway struct {
	*Client
}

// InvoiceNote returns a new gateway instance
func (c *Client) InvoiceNote() *InvoiceNoteGateway {
	return &InvoiceNoteGateway{c}
}

// Create adds a note to an existing invoice
func (c *InvoiceNoteGateway) Create(invoice *Invoice, note *InvoiceNote) (*InvoiceNote, error) {
	res, err := c.execute("POST", "sales_invoices/"+invoice.ID+"/notes", &envelope{InvoiceNote: note})
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 201:
		return res.note()
	}

	return nil, res.error()
}

// Delete destroys the invoice note
func (c *InvoiceNoteGateway) Delete(invoice *Invoice, note *InvoiceNote) error {
	res, err := c.execute("DELETE", "sales_invoices/"+invoice.ID+"/notes/"+note.ID, nil)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		return nil
	}

	return res.error()
}
