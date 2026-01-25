package requests

import (
	"testing"
)

func TestGetMyStarBalance_GetValues(t *testing.T) {
	request := &GetMyStarBalance{}

	values, err := request.GetValues()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if values != nil && len(values) != 0 {
		t.Errorf("expected nil or empty values for request without fields")
	}
}

func TestGetMyStarBalance_GetFiles(t *testing.T) {
	request := &GetMyStarBalance{}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
