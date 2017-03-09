package moneybird

import (
	"testing"
)

func TestTaxRateGatewayAll(t *testing.T) {
	taxrates, err := testClient.TaxRate().All()

	if err != nil {
		t.Error(err)
	}

	if len(taxrates) == 0 {
		t.Error("Client returned 0 tax rates but sandbox account should have multiple.")
	}
}
