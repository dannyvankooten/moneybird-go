package moneybird

import (
	"net/http"
	"os"
	"fmt"
)

var testClient *Client

func init() {
	if os.Getenv("MONEYBIRD_TEST_TOKEN") == "" {
		fmt.Printf("Environment value MONEYBIRD_TEST_TOKEN not set\n")
		os.Exit(1)
	}

	if os.Getenv("MONEYBIRD_TEST_ADMINISTRATION_ID") == "" {
		fmt.Printf("Environment value MONEYBIRD_TEST_ADMINISTRATION_ID not set\n")
		os.Exit(1)
	}

	testClient = &Client{
		Token:            os.Getenv("MONEYBIRD_TEST_TOKEN"),
		AdministrationID: os.Getenv("MONEYBIRD_TEST_ADMINISTRATION_ID"),
		HTTPClient:       &http.Client{},
	}
}