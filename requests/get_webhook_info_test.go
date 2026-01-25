package requests

import (
	"testing"
)

func TestGetWebhookInfo_GetValues(t *testing.T) {
	request := &GetWebhookInfo{}

	values, err := request.GetValues()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if values != nil && len(values) != 0 {
		t.Errorf("expected nil or empty values for request without fields")
	}
}

func TestGetWebhookInfo_GetFiles(t *testing.T) {
	request := &GetWebhookInfo{}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
