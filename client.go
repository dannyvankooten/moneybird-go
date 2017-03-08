package moneybird

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

// Client is the MoneyBird API consumer
type Client struct {
	Token            string
	AdministrationID string

	HTTPClient *http.Client
}

type envelope struct {
	Contact        *Contact        `json:"contact,omitempty"`
	Invoice        *Invoice        `json:"sales_invoice,omitempty"`
	InvoiceSending *InvoiceSending `json:"sales_invoice_sending,omitempty"`
	InvoicePayment *InvoicePayment `json:"payment,omitempty"`
}

func (c *Client) resourceURL(path string) string {
	return "https://moneybird.com/api/v2/" + c.AdministrationID + "/" + path
}

func (c *Client) newRequest(method string, path string, data *envelope) (*http.Request, error) {
	var body []byte
	var err error

	if data != nil {
		body, err = json.Marshal(data)

		if err != nil {
			return nil, err
		}
	}

	url := c.resourceURL(path)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req, nil
}

func (c *Client) execute(method string, path string, data *envelope) (*Response, error) {
	req, err := c.newRequest(method, path, data)
	if err != nil {
		return nil, err
	}

	log.Printf("Moneybird: %s %s\n", req.Method, req.URL)
	res, err := c.HTTPClient.Do(req)
	log.Printf("Moneybird: %d %s", res.StatusCode, res.Status)

	if err != nil {
		return nil, err
	}

	// TODO: Move to Response class & improve.
	if res.StatusCode > 399 {
		var data map[string]string
		err = json.NewDecoder(res.Body).Decode(&data)
		log.Printf("Moneybird: error data %#v\n", data)
		return nil, errors.New(data["error"])
	}

	return &Response{res}, err
}
