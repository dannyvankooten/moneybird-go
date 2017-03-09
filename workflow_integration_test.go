package moneybird

import (
	"testing"
)

func TestWorkflowGatewayList(t *testing.T) {
	_, err := testClient.Workflow().List()

	if err != nil {
		t.Error(err)
	}
}
