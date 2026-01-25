package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendMessage_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendMessage
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendMessage{
				ChatId: telegram.NewChatId(123456, ""),
				Text:   "test_text",
			},
			expected: map[string]string{
				"chat_id": "123456",
				"text":    "test_text",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendMessage{
				Text:                  "test_text",
				AllowPaidBroadcast:    ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				ParseMode:             ptr("test_parse_mode"),
				ProtectContent:        ptr(true),
			},
			expected: map[string]string{
				"text":                     "test_text",
				"allow_paid_broadcast":     "1",
				"business_connection_id":   "test_business_connection_id",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"message_effect_id":        "test_message_effect_id",
				"message_thread_id":        "456",
				"parse_mode":               "test_parse_mode",
				"protect_content":          "1",
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

func TestSendMessage_GetFiles(t *testing.T) {
	request := &SendMessage{
		Text: "test_text",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
