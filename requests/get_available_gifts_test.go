package requests

import (
	"testing"
)

func TestGetAvailableGifts_GetValues(t *testing.T) {
	request := &GetAvailableGifts{}

	values, err := request.GetValues()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if values != nil && len(values) != 0 {
		t.Errorf("expected nil or empty values for request without fields")
	}
}

func TestGetAvailableGifts_GetFiles(t *testing.T) {
	request := &GetAvailableGifts{}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
