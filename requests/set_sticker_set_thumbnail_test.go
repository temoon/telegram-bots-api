package requests

import (
	"bytes"
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSetStickerSetThumbnail_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SetStickerSetThumbnail
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SetStickerSetThumbnail{
				Format: "test_format",
				Name:   "test_name",
				UserId: 123,
			},
			expected: map[string]string{
				"format":  "test_format",
				"name":    "test_name",
				"user_id": "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SetStickerSetThumbnail{
				Format:    "test_format",
				Name:      "test_name",
				UserId:    123,
				Thumbnail: ptr(telegram.NewInputFile("file_id_456", nil, "")),
			},
			expected: map[string]string{
				"format":    "test_format",
				"name":      "test_name",
				"user_id":   "123",
				"thumbnail": "file_id_456",
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

func TestSetStickerSetThumbnail_GetFiles(t *testing.T) {
	t.Run("with file reader", func(t *testing.T) {
		reader := bytes.NewReader([]byte("fake file data"))
		request := &SetStickerSetThumbnail{
			Format:    "test_format",
			Name:      "test_name",
			UserId:    123,
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
		request := &SetStickerSetThumbnail{
			Format: "test_format",
			Name:   "test_name",
			UserId: 123,
		}

		files := request.GetFiles()

		if len(files) != 0 {
			t.Errorf("expected 0 files when using file_id, got %d", len(files))
		}
	})
}
