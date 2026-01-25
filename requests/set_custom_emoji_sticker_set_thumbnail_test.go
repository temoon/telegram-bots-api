package requests

import (
	"testing"
)

func TestSetCustomEmojiStickerSetThumbnail_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SetCustomEmojiStickerSetThumbnail
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SetCustomEmojiStickerSetThumbnail{
				Name: "test_name",
			},
			expected: map[string]string{
				"name": "test_name",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SetCustomEmojiStickerSetThumbnail{
				Name:          "test_name",
				CustomEmojiId: ptr("test_custom_emoji_id"),
			},
			expected: map[string]string{
				"name":            "test_name",
				"custom_emoji_id": "test_custom_emoji_id",
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

func TestSetCustomEmojiStickerSetThumbnail_GetFiles(t *testing.T) {
	request := &SetCustomEmojiStickerSetThumbnail{
		Name: "test_name",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
