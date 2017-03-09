package moneybird

import (
	"net/http"
	"os"
)

var testClient = &Client{
	Token:            os.Getenv("MONEYBIRD_TEST_TOKEN"),
	AdministrationID: os.Getenv("MONEYBIRD_TEST_ADMINISTRATION_ID"),
	HTTPClient:       &http.Client{},
}
