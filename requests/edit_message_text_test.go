package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestEditMessageText_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *EditMessageText
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &EditMessageText{
				Text: "test_text",
			},
			expected: map[string]string{
				"text": "test_text",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &EditMessageText{
				Text:                 "test_text",
				BusinessConnectionId: ptr("test_business_connection_id"),
				ChatId:               ptr(telegram.NewChatId(789, "")),
				InlineMessageId:      ptr("test_inline_message_id"),
				MessageId:            ptr(int64(456)),
				ParseMode:            ptr("test_parse_mode"),
			},
			expected: map[string]string{
				"text":                   "test_text",
				"business_connection_id": "test_business_connection_id",
				"chat_id":                "789",
				"inline_message_id":      "test_inline_message_id",
				"message_id":             "456",
				"parse_mode":             "test_parse_mode",
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

func TestEditMessageText_GetFiles(t *testing.T) {
	request := &EditMessageText{
		Text: "test_text",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
