package requests

import (
	"testing"
)

func TestSetMyShortDescription_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SetMyShortDescription
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "with optional fields",
			request: &SetMyShortDescription{
				LanguageCode:     ptr("test_language_code"),
				ShortDescription: ptr("test_short_description"),
			},
			expected: map[string]string{
				"language_code":     "test_language_code",
				"short_description": "test_short_description",
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

func TestSetMyShortDescription_GetFiles(t *testing.T) {
	request := &SetMyShortDescription{}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
