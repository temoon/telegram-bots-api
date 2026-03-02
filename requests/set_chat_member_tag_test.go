package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSetChatMemberTag_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SetChatMemberTag
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SetChatMemberTag{
				ChatId: telegram.NewChatId(123456, ""),
				UserId: 123,
			},
			expected: map[string]string{
				"chat_id": "123456",
				"user_id": "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SetChatMemberTag{
				UserId: 123,
				Tag:    ptr("test_tag"),
			},
			expected: map[string]string{
				"user_id": "123",
				"tag":     "test_tag",
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

func TestSetChatMemberTag_GetFiles(t *testing.T) {
	request := &SetChatMemberTag{
		UserId: 123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
