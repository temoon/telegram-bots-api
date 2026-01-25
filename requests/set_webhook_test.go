package requests

import (
	"bytes"
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSetWebhook_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SetWebhook
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SetWebhook{
				Url: "test_url",
			},
			expected: map[string]string{
				"url": "test_url",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SetWebhook{
				Url:                "test_url",
				Certificate:        ptr(telegram.NewInputFile("file_id_456", nil, "")),
				DropPendingUpdates: ptr(true),
				IpAddress:          ptr("test_ip_address"),
				MaxConnections:     ptr(int64(456)),
				SecretToken:        ptr("test_secret_token"),
			},
			expected: map[string]string{
				"url":                  "test_url",
				"certificate":          "file_id_456",
				"drop_pending_updates": "1",
				"ip_address":           "test_ip_address",
				"max_connections":      "456",
				"secret_token":         "test_secret_token",
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

func TestSetWebhook_GetFiles(t *testing.T) {
	t.Run("with file reader", func(t *testing.T) {
		reader := bytes.NewReader([]byte("fake file data"))
		request := &SetWebhook{
			Url:         "test_url",
			Certificate: &telegram.InputFile{},
		}
		// Set file for optional InputFile fields
		*request.Certificate = telegram.NewInputFile("", reader, "test.jpg")

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
		request := &SetWebhook{
			Url: "test_url",
		}

		files := request.GetFiles()

		if len(files) != 0 {
			t.Errorf("expected 0 files when using file_id, got %d", len(files))
		}
	})
}
