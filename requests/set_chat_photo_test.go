package requests

import (
	"bytes"
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSetChatPhoto_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SetChatPhoto
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SetChatPhoto{
				ChatId: telegram.NewChatId(123456, ""),
				Photo:  telegram.NewInputFile("file_id_123", nil, ""),
			},
			expected: map[string]string{
				"chat_id": "123456",
				"photo":   "file_id_123",
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

func TestSetChatPhoto_GetFiles(t *testing.T) {
	t.Run("with file reader", func(t *testing.T) {
		reader := bytes.NewReader([]byte("fake file data"))
		request := &SetChatPhoto{
			ChatId: telegram.NewChatId(123456, ""),
			Photo:  telegram.NewInputFile("", reader, "test.jpg"),
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
		request := &SetChatPhoto{}

		files := request.GetFiles()

		if len(files) != 0 {
			t.Errorf("expected 0 files when using file_id, got %d", len(files))
		}
	})
}
