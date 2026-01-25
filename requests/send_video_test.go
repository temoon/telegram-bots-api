package requests

import (
	"bytes"
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendVideo_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendVideo
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendVideo{
				ChatId: telegram.NewChatId(123456, ""),
				Video:  telegram.NewInputFile("file_id_123", nil, ""),
			},
			expected: map[string]string{
				"chat_id": "123456",
				"video":   "file_id_123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendVideo{
				AllowPaidBroadcast:    ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				Caption:               ptr("test_caption"),
				Cover:                 ptr(telegram.NewInputFile("file_id_456", nil, "")),
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				Duration:              ptr(int64(456)),
				HasSpoiler:            ptr(true),
				Height:                ptr(int64(456)),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				ParseMode:             ptr("test_parse_mode"),
				ProtectContent:        ptr(true),
				ShowCaptionAboveMedia: ptr(true),
				StartTimestamp:        ptr(int64(456)),
				SupportsStreaming:     ptr(true),
				Thumbnail:             ptr(telegram.NewInputFile("file_id_456", nil, "")),
				Width:                 ptr(int64(456)),
			},
			expected: map[string]string{
				"allow_paid_broadcast":     "1",
				"business_connection_id":   "test_business_connection_id",
				"caption":                  "test_caption",
				"cover":                    "file_id_456",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"duration":                 "456",
				"has_spoiler":              "1",
				"height":                   "456",
				"message_effect_id":        "test_message_effect_id",
				"message_thread_id":        "456",
				"parse_mode":               "test_parse_mode",
				"protect_content":          "1",
				"show_caption_above_media": "1",
				"start_timestamp":          "456",
				"supports_streaming":       "1",
				"thumbnail":                "file_id_456",
				"width":                    "456",
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

func TestSendVideo_GetFiles(t *testing.T) {
	t.Run("with file reader", func(t *testing.T) {
		reader := bytes.NewReader([]byte("fake file data"))
		request := &SendVideo{
			ChatId:    telegram.NewChatId(123456, ""),
			Video:     telegram.NewInputFile("", reader, "test.jpg"),
			Cover:     &telegram.InputFile{},
			Thumbnail: &telegram.InputFile{},
		}
		// Set file for optional InputFile fields
		*request.Cover = telegram.NewInputFile("", reader, "test.jpg")
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
		request := &SendVideo{}

		files := request.GetFiles()

		if len(files) != 0 {
			t.Errorf("expected 0 files when using file_id, got %d", len(files))
		}
	})
}
