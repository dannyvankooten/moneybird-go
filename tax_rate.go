package moneybird

import "encoding/json"

// TaxRate contains info about a tax rate stored in Moneybird
type TaxRate struct {
	ID               string `json:"id"`
	AdministrationID string `json:"administration_id,omitempty"`
	Name             string `json:"name"`
	Percentage       string `json:"percentage"`
	TaxRateType      string `json:"tax_rate_type"`
	ShowTax          bool   `json:"show_tax"`
	Active           bool   `json:"active"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
}

// TaxRateGateway encapsulates all /tax_rates related endpoints
type TaxRateGateway struct {
	*Client
}

// TaxRate returns a TaxRateGateway instance
func (c *Client) TaxRate() *TaxRateGateway {
	return &TaxRateGateway{c}
}

// All returns all tax rates stored in Moneybird
func (c *TaxRateGateway) All() ([]*TaxRate, error) {
	var taxrates []*TaxRate
	var err error

	res, err := c.execute("GET", "tax_rates", nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		err = json.NewDecoder(res.Body).Decode(&taxrates)
		return taxrates, err
	}

	return nil, res.error()
}
