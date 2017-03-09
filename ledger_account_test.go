package moneybird

import (
	"testing"
)

func TestLedgerAccountGatewayAll(t *testing.T) {
	_, err := testClient.LedgerAccount().List()

	if err != nil {
		t.Error(err)
	}
}
