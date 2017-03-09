package moneybird

import (
	"encoding/json"
	"time"
)

// Workflow contains info about a tax rate stored in Moneybird
type Workflow struct {
	ID               string    `json:"id"`
	AdministrationID string    `json:"administration_id,omitempty"`
	Type             string    `json:"type,omitempty"`
	Name             string    `json:"name,omitempty"`
	Default          bool      `json:"default,omitempty"`
	Currency         string    `json:"currency,omitempty"`
	Language         string    `json:"language,omitempty"`
	Active           bool      `json:"active,omitempty"`
	PricesAreInclTax bool      `json:"prices_are_incl_tax,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`
}

// WorkflowGateway encapsulates all /tax_rates related endpoints
type WorkflowGateway struct {
	*Client
}

// Workflow returns a WorkflowGateway instance
func (c *Client) Workflow() *WorkflowGateway {
	return &WorkflowGateway{c}
}

// List returns all workflows stored in Moneybird
func (c *WorkflowGateway) List() ([]*Workflow, error) {
	var err error
	var workflows []*Workflow

	res, err := c.execute("GET", "workflows", nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		err = json.NewDecoder(res.Body).Decode(&workflows)
		return workflows, err
	}
	return nil, res.error()
}
