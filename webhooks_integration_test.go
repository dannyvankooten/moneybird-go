package moneybird

import (
	"testing"
)

func TestWebhookGatewayList(t *testing.T) {
	_, err := testClient.Webhook().List()

	if err != nil {
		t.Error(err)
	}
}
