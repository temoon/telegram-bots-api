package requests

import (
	"testing"
)

func TestGetForumTopicIconStickers_GetValues(t *testing.T) {
	request := &GetForumTopicIconStickers{}

	values, err := request.GetValues()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if values != nil && len(values) != 0 {
		t.Errorf("expected nil or empty values for request without fields")
	}
}

func TestGetForumTopicIconStickers_GetFiles(t *testing.T) {
	request := &GetForumTopicIconStickers{}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
