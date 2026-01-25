package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendChatAction_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendChatAction
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendChatAction{
				Action: "test_action",
				ChatId: telegram.NewChatId(123456, ""),
			},
			expected: map[string]string{
				"action":  "test_action",
				"chat_id": "123456",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendChatAction{
				Action:               "test_action",
				BusinessConnectionId: ptr("test_business_connection_id"),
				MessageThreadId:      ptr(int64(456)),
			},
			expected: map[string]string{
				"action":                 "test_action",
				"business_connection_id": "test_business_connection_id",
				"message_thread_id":      "456",
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

func TestSendChatAction_GetFiles(t *testing.T) {
	request := &SendChatAction{
		Action: "test_action",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
