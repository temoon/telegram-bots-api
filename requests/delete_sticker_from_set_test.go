package requests

import (
	"testing"
)

func TestDeleteStickerFromSet_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *DeleteStickerFromSet
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &DeleteStickerFromSet{
				Sticker: "test_sticker",
			},
			expected: map[string]string{
				"sticker": "test_sticker",
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

func TestDeleteStickerFromSet_GetFiles(t *testing.T) {
	request := &DeleteStickerFromSet{
		Sticker: "test_sticker",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
