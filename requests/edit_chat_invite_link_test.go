package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestEditChatInviteLink_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *EditChatInviteLink
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &EditChatInviteLink{
				ChatId:     telegram.NewChatId(123456, ""),
				InviteLink: "test_invite_link",
			},
			expected: map[string]string{
				"chat_id":     "123456",
				"invite_link": "test_invite_link",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &EditChatInviteLink{
				InviteLink:         "test_invite_link",
				CreatesJoinRequest: ptr(true),
				ExpireDate:         ptr(int64(456)),
				MemberLimit:        ptr(int64(456)),
				Name:               ptr("test_name"),
			},
			expected: map[string]string{
				"invite_link":          "test_invite_link",
				"creates_join_request": "1",
				"expire_date":          "456",
				"member_limit":         "456",
				"name":                 "test_name",
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

func TestEditChatInviteLink_GetFiles(t *testing.T) {
	request := &EditChatInviteLink{
		InviteLink: "test_invite_link",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
