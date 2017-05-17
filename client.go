package moneybird

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Client is the MoneyBird API consumer
type Client struct {
	Token            string
	AdministrationID string

	Logger     *log.Logger
	HTTPClient *http.Client
}

type envelope struct {
	Contact        *Contact        `json:"contact,omitempty"`
	Invoice        *Invoice        `json:"sales_invoice,omitempty"`
	InvoiceSending *InvoiceSending `json:"sales_invoice_sending,omitempty"`
	InvoicePayment *InvoicePayment `json:"payment,omitempty"`
	InvoiceNote    *InvoiceNote    `json:"note,omitempty"`
	LedgerAccount  *LedgerAccount  `json:"ledger_account,omitempty"`
}

func (c *Client) resourceURL(path string) string {
	return "https://moneybird.com/api/v2/" + c.AdministrationID + "/" + path
}

func (c *Client) newRequest(method string, path string, data []byte) (*http.Request, error) {

	var err error

	url := c.resourceURL(path)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	return req, nil
}

func (c *Client) execute(method string, path string, env *envelope) (*Response, error) {
	var data []byte
	var err error

	if env != nil {
		data, err = json.Marshal(env)

		if err != nil {
			return nil, err
		}
	}

	req, err := c.newRequest(method, path, data)
	if err != nil {
		return nil, err
	}

	if c.Logger != nil {
		c.Logger.Printf("Moneybird: %s %s\n", req.Method, req.URL)
	}
	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return nil, err
	}

	if c.Logger != nil {
		body, _ := ioutil.ReadAll(res.Body)
		c.Logger.Printf("Moneybird: %s", res.Status)
		c.Logger.Printf("Moneybird: %s", body)
		res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}

	return &Response{res}, nil
}
