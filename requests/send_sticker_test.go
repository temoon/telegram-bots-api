package requests

import (
	"bytes"
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendSticker_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendSticker
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendSticker{
				ChatId:  telegram.NewChatId(123456, ""),
				Sticker: telegram.NewInputFile("file_id_123", nil, ""),
			},
			expected: map[string]string{
				"chat_id": "123456",
				"sticker": "file_id_123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendSticker{
				AllowPaidBroadcast:    ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				Emoji:                 ptr("test_emoji"),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				ProtectContent:        ptr(true),
			},
			expected: map[string]string{
				"allow_paid_broadcast":     "1",
				"business_connection_id":   "test_business_connection_id",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"emoji":                    "test_emoji",
				"message_effect_id":        "test_message_effect_id",
				"message_thread_id":        "456",
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

func TestSendSticker_GetFiles(t *testing.T) {
	t.Run("with file reader", func(t *testing.T) {
		reader := bytes.NewReader([]byte("fake file data"))
		request := &SendSticker{
			ChatId:  telegram.NewChatId(123456, ""),
			Sticker: telegram.NewInputFile("", reader, "test.jpg"),
		}
		// Set file for optional InputFile fields

		files := request.GetFiles()

		if len(files) == 0 {
			t.Error("expected files to be present")
		}

		// Verify that the reader is in the files map
		found := false
		for _, r := range files {
			if r == reader {
				found = true
				break
			}
		}

		if !found {
			t.Error("expected file reader to be in files map")
		}
	})

	t.Run("with file_id", func(t *testing.T) {
		request := &SendSticker{}

		files := request.GetFiles()

		if len(files) != 0 {
			t.Errorf("expected 0 files when using file_id, got %d", len(files))
		}
	})
}
