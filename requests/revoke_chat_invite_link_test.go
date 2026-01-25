package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestRevokeChatInviteLink_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *RevokeChatInviteLink
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &RevokeChatInviteLink{
				ChatId:     telegram.NewChatId(123456, ""),
				InviteLink: "test_invite_link",
			},
			expected: map[string]string{
				"chat_id":     "123456",
				"invite_link": "test_invite_link",
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

func TestRevokeChatInviteLink_GetFiles(t *testing.T) {
	request := &RevokeChatInviteLink{
		InviteLink: "test_invite_link",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
