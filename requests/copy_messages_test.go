package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestCopyMessages_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *CopyMessages
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &CopyMessages{
				ChatId:     telegram.NewChatId(123456, ""),
				FromChatId: telegram.NewChatId(123456, ""),
			},
			expected: map[string]string{
				"chat_id":      "123456",
				"from_chat_id": "123456",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &CopyMessages{
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				MessageThreadId:       ptr(int64(456)),
				ProtectContent:        ptr(true),
				RemoveCaption:         ptr(true),
			},
			expected: map[string]string{
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"message_thread_id":        "456",
				"protect_content":          "1",
				"remove_caption":           "1",
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

func TestCopyMessages_GetFiles(t *testing.T) {
	request := &CopyMessages{}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
