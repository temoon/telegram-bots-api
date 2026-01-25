package requests

import (
	"bytes"
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendAudio_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendAudio
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendAudio{
				Audio:  telegram.NewInputFile("file_id_123", nil, ""),
				ChatId: telegram.NewChatId(123456, ""),
			},
			expected: map[string]string{
				"audio":   "file_id_123",
				"chat_id": "123456",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendAudio{
				AllowPaidBroadcast:    ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				Caption:               ptr("test_caption"),
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				Duration:              ptr(int64(456)),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				ParseMode:             ptr("test_parse_mode"),
				Performer:             ptr("test_performer"),
				ProtectContent:        ptr(true),
				Thumbnail:             ptr(telegram.NewInputFile("file_id_456", nil, "")),
				Title:                 ptr("test_title"),
			},
			expected: map[string]string{
				"allow_paid_broadcast":     "1",
				"business_connection_id":   "test_business_connection_id",
				"caption":                  "test_caption",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"duration":                 "456",
				"message_effect_id":        "test_message_effect_id",
				"message_thread_id":        "456",
				"parse_mode":               "test_parse_mode",
				"performer":                "test_performer",
				"protect_content":          "1",
				"thumbnail":                "file_id_456",
				"title":                    "test_title",
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

func TestSendAudio_GetFiles(t *testing.T) {
	t.Run("with file reader", func(t *testing.T) {
		reader := bytes.NewReader([]byte("fake file data"))
		request := &SendAudio{
			Audio:     telegram.NewInputFile("", reader, "test.jpg"),
			ChatId:    telegram.NewChatId(123456, ""),
			Thumbnail: &telegram.InputFile{},
		}
		// Set file for optional InputFile fields
		*request.Thumbnail = telegram.NewInputFile("", reader, "test.jpg")

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
		request := &SendAudio{}

		files := request.GetFiles()

		if len(files) != 0 {
			t.Errorf("expected 0 files when using file_id, got %d", len(files))
		}
	})
}
