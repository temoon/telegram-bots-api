package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendPaidMedia_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendPaidMedia
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendPaidMedia{
				ChatId:    telegram.NewChatId(123456, ""),
				StarCount: 123,
			},
			expected: map[string]string{
				"chat_id":    "123456",
				"star_count": "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendPaidMedia{
				StarCount:             123,
				AllowPaidBroadcast:    ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				Caption:               ptr("test_caption"),
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				MessageThreadId:       ptr(int64(456)),
				ParseMode:             ptr("test_parse_mode"),
				Payload:               ptr("test_payload"),
				ProtectContent:        ptr(true),
				ShowCaptionAboveMedia: ptr(true),
			},
			expected: map[string]string{
				"star_count":               "123",
				"allow_paid_broadcast":     "1",
				"business_connection_id":   "test_business_connection_id",
				"caption":                  "test_caption",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"message_thread_id":        "456",
				"parse_mode":               "test_parse_mode",
				"payload":                  "test_payload",
				"protect_content":          "1",
				"show_caption_above_media": "1",
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

func TestSendPaidMedia_GetFiles(t *testing.T) {
	request := &SendPaidMedia{
		StarCount: 123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
