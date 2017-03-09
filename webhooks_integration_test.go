package moneybird

import (
	"testing"
)

func TestWebhookGatewayCreate(t *testing.T) {
	webhook, err := testClient.Webhook().Create(&Webhook{
		URL: "http://mockbin.org/bin/bbe7f656-12d6-4877-9fa8-5cd61f9522a9/view",
	})

	if err != nil {
		t.Fatal(err)
	}

	if webhook.ID == "" {
		t.Fatal("Webhook ID was not set")
	}

	if webhook.URL != "http://mockbin.org/bin/bbe7f656-12d6-4877-9fa8-5cd61f9522a9/view" {
		t.Fatalf("Invalid webhook URL: %#v", webhook.URL)
	}
}

func TestWebhookGatewayListAndDelete(t *testing.T) {
	gateway := testClient.Webhook()
	webhooks, err := gateway.List()

	if err != nil {
		t.Fatal(err)
	}

	for _, w := range webhooks {
		err := gateway.Delete(w)
		if err != nil {
			t.Fatal(err)
		}
	}
}
