package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestEditForumTopic_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *EditForumTopic
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &EditForumTopic{
				ChatId:          telegram.NewChatId(123456, ""),
				MessageThreadId: 123,
			},
			expected: map[string]string{
				"chat_id":           "123456",
				"message_thread_id": "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &EditForumTopic{
				MessageThreadId:   123,
				IconCustomEmojiId: ptr("test_icon_custom_emoji_id"),
				Name:              ptr("test_name"),
			},
			expected: map[string]string{
				"message_thread_id":    "123",
				"icon_custom_emoji_id": "test_icon_custom_emoji_id",
				"name":                 "test_name",
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

func TestEditForumTopic_GetFiles(t *testing.T) {
	request := &EditForumTopic{
		MessageThreadId: 123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
