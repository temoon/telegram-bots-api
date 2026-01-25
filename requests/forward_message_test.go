package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestForwardMessage_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *ForwardMessage
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &ForwardMessage{
				ChatId:     telegram.NewChatId(123456, ""),
				FromChatId: telegram.NewChatId(123456, ""),
				MessageId:  123,
			},
			expected: map[string]string{
				"chat_id":      "123456",
				"from_chat_id": "123456",
				"message_id":   "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &ForwardMessage{
				MessageId:             123,
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				ProtectContent:        ptr(true),
				VideoStartTimestamp:   ptr(int64(456)),
			},
			expected: map[string]string{
				"message_id":               "123",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"message_effect_id":        "test_message_effect_id",
				"message_thread_id":        "456",
				"protect_content":          "1",
				"video_start_timestamp":    "456",
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

func TestForwardMessage_GetFiles(t *testing.T) {
	request := &ForwardMessage{
		MessageId: 123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
