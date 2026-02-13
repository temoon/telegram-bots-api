package requests

import (
	"testing"
)

func TestRemoveMyProfilePhoto_GetValues(t *testing.T) {
	request := &RemoveMyProfilePhoto{}

	values, err := request.GetValues()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if values != nil && len(values) != 0 {
		t.Errorf("expected nil or empty values for request without fields")
	}
}

func TestRemoveMyProfilePhoto_GetFiles(t *testing.T) {
	request := &RemoveMyProfilePhoto{}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
