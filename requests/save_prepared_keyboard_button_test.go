package requests

import (
	"testing"
)

func TestSavePreparedKeyboardButton_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SavePreparedKeyboardButton
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SavePreparedKeyboardButton{
				UserId: 123,
			},
			expected: map[string]string{
				"user_id": "123",
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

func TestSavePreparedKeyboardButton_GetFiles(t *testing.T) {
	request := &SavePreparedKeyboardButton{
		UserId: 123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
