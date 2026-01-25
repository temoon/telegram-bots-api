package requests

import (
	"testing"
)

func TestSetGameScore_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SetGameScore
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SetGameScore{
				Score:  123,
				UserId: 123,
			},
			expected: map[string]string{
				"score":   "123",
				"user_id": "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SetGameScore{
				Score:              123,
				UserId:             123,
				ChatId:             ptr(int64(456)),
				DisableEditMessage: ptr(true),
				Force:              ptr(true),
				InlineMessageId:    ptr("test_inline_message_id"),
				MessageId:          ptr(int64(456)),
			},
			expected: map[string]string{
				"score":                "123",
				"user_id":              "123",
				"chat_id":              "456",
				"disable_edit_message": "1",
				"force":                "1",
				"inline_message_id":    "test_inline_message_id",
				"message_id":           "456",
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

func TestSetGameScore_GetFiles(t *testing.T) {
	request := &SetGameScore{
		Score:  123,
		UserId: 123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
