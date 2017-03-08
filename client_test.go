package moneybird

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestClientBaseURL(t *testing.T) {
	c := &Client{
		AdministrationID: "administration-id",
	}

	if u := c.resourceURL("contacts"); !strings.HasSuffix(u, "administration-id/contacts") {
		t.Errorf("Resource URL does not include administration ID. Got %#v", u)
	}
}

func TestNewRequest(t *testing.T) {
	var req *http.Request

	client := &Client{
		AdministrationID: "administration-id",
		Token:            "token",
	}

	data := &envelope{
		Contact: &Contact{
			ID: "id",
		}}

	req, _ = client.newRequest("POST", "contacts", data)
	if h := req.Header.Get("Authorization"); h != "Bearer token" {
		t.Errorf("Expected header %#v, got %#v", "Bearer token", h)
	}

	expected := "application/json"
	if h := req.Header.Get("Content-Type"); h != expected {
		t.Errorf("Expected header %#v, got %#v", expected, h)
	}

	// test marshaling of struct
	expectedBody, _ := json.Marshal(data)
	body, _ := ioutil.ReadAll(req.Body)

	if len(body) != len(expectedBody) {
		t.Errorf("Request body does not have expected size")
	}
}
