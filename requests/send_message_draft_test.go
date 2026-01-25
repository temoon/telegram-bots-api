package requests

import (
	"testing"
)

func TestSendMessageDraft_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendMessageDraft
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendMessageDraft{
				ChatId:  123,
				DraftId: 123,
				Text:    "test_text",
			},
			expected: map[string]string{
				"chat_id":  "123",
				"draft_id": "123",
				"text":     "test_text",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendMessageDraft{
				ChatId:          123,
				DraftId:         123,
				Text:            "test_text",
				MessageThreadId: ptr(int64(456)),
				ParseMode:       ptr("test_parse_mode"),
			},
			expected: map[string]string{
				"chat_id":           "123",
				"draft_id":          "123",
				"text":              "test_text",
				"message_thread_id": "456",
				"parse_mode":        "test_parse_mode",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			values, err := tt.request.GetValues()

			if (err != nil) != tt.wantErr {
				t.Errorf("GetValues() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Check that all expected fields are present
			for key, want := range tt.expected {
				if got, ok := values[key]; !ok {
					t.Errorf("missing field %q", key)
				} else if got != want {
					t.Errorf("field %q = %q, want %q", key, got, want)
				}
			}
		})
	}
}

func TestSendMessageDraft_GetFiles(t *testing.T) {
	request := &SendMessageDraft{
		ChatId:  123,
		DraftId: 123,
		Text:    "test_text",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
