package moneybird

import "encoding/json"

//"time"

// Webhook is a MoneyBird webhook
type Webhook struct {
	ID               string `json:"id,omitempty"`
	AdministrationID string `json:"administration_id,omitempty"`
	URL              string `json:"url"`
	LastHTTPStatus   string `json:"last_http_status,omitempty"`
	LastHTTPBody     string `json:"last_http_body,omitempty"`
}

// WebhookGateway encapsulates all /webhooks related endpoints
type WebhookGateway struct {
	*Client
}

// Webhook returns a new gateway instance
func (c *Client) Webhook() *WebhookGateway {
	return &WebhookGateway{c}
}

// List returns all webhooks in Moneybird
func (c *WebhookGateway) List() ([]*Webhook, error) {
	var webhooks []*Webhook
	var err error

	res, err := c.execute("GET", "webhooks", nil)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case 200:
		err = json.NewDecoder(res.Body).Decode(webhooks)
		return webhooks, err
	}

	return nil, res.error()
}

// Create adds a webhook to MoneyBird
func (c *WebhookGateway) Create(webhook *Webhook) (*Webhook, error) {
	data, err := json.Marshal(webhook)
	if err != nil {
		return nil, err
	}

	// webhooks do not use an envelope so we can't use c.execute here... bummer
	req, err := c.newRequest("POST", "webhooks", data)
	if err != nil {
		return nil, err
	}

	httpRes, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	res := &Response{httpRes}
	switch res.StatusCode {
	case 201:
		return res.webhook()
	}

	return nil, res.error()
}

// Delete the given webhook
func (c *WebhookGateway) Delete(webhook *Webhook) error {
	res, err := c.execute("DELETE", "webhooks/"+webhook.ID, nil)
	if err != nil {
		return err
	}

	switch res.StatusCode {
	case 204:
		return nil
	}

	return res.error()
}
