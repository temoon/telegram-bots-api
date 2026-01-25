package requests

import (
	"bytes"
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestUploadStickerFile_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *UploadStickerFile
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &UploadStickerFile{
				Sticker:       telegram.NewInputFile("file_id_123", nil, ""),
				StickerFormat: "test_sticker_format",
				UserId:        123,
			},
			expected: map[string]string{
				"sticker":        "file_id_123",
				"sticker_format": "test_sticker_format",
				"user_id":        "123",
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

func TestUploadStickerFile_GetFiles(t *testing.T) {
	t.Run("with file reader", func(t *testing.T) {
		reader := bytes.NewReader([]byte("fake file data"))
		request := &UploadStickerFile{
			Sticker:       telegram.NewInputFile("", reader, "test.jpg"),
			StickerFormat: "test_sticker_format",
			UserId:        123,
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
		request := &UploadStickerFile{
			StickerFormat: "test_sticker_format",
			UserId:        123,
		}

		files := request.GetFiles()

		if len(files) != 0 {
			t.Errorf("expected 0 files when using file_id, got %d", len(files))
		}
	})
}
