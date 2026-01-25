package requests

import (
	"bytes"
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendDocument_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendDocument
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendDocument{
				ChatId:   telegram.NewChatId(123456, ""),
				Document: telegram.NewInputFile("file_id_123", nil, ""),
			},
			expected: map[string]string{
				"chat_id":  "123456",
				"document": "file_id_123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendDocument{
				AllowPaidBroadcast:          ptr(true),
				BusinessConnectionId:        ptr("test_business_connection_id"),
				Caption:                     ptr("test_caption"),
				DirectMessagesTopicId:       ptr(int64(456)),
				DisableContentTypeDetection: ptr(true),
				DisableNotification:         ptr(true),
				MessageEffectId:             ptr("test_message_effect_id"),
				MessageThreadId:             ptr(int64(456)),
				ParseMode:                   ptr("test_parse_mode"),
				ProtectContent:              ptr(true),
				Thumbnail:                   ptr(telegram.NewInputFile("file_id_456", nil, "")),
			},
			expected: map[string]string{
				"allow_paid_broadcast":           "1",
				"business_connection_id":         "test_business_connection_id",
				"caption":                        "test_caption",
				"direct_messages_topic_id":       "456",
				"disable_content_type_detection": "1",
				"disable_notification":           "1",
				"message_effect_id":              "test_message_effect_id",
				"message_thread_id":              "456",
				"parse_mode":                     "test_parse_mode",
				"protect_content":                "1",
				"thumbnail":                      "file_id_456",
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

func TestSendDocument_GetFiles(t *testing.T) {
	t.Run("with file reader", func(t *testing.T) {
		reader := bytes.NewReader([]byte("fake file data"))
		request := &SendDocument{
			ChatId:    telegram.NewChatId(123456, ""),
			Document:  telegram.NewInputFile("", reader, "test.jpg"),
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
		request := &SendDocument{}

		files := request.GetFiles()

		if len(files) != 0 {
			t.Errorf("expected 0 files when using file_id, got %d", len(files))
		}
	})
}
