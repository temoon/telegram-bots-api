package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestPromoteChatMember_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *PromoteChatMember
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &PromoteChatMember{
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
			request: &PromoteChatMember{
				UserId:                  123,
				CanChangeInfo:           ptr(true),
				CanDeleteMessages:       ptr(true),
				CanDeleteStories:        ptr(true),
				CanEditMessages:         ptr(true),
				CanEditStories:          ptr(true),
				CanInviteUsers:          ptr(true),
				CanManageChat:           ptr(true),
				CanManageDirectMessages: ptr(true),
				CanManageTopics:         ptr(true),
				CanManageVideoChats:     ptr(true),
				CanPinMessages:          ptr(true),
				CanPostMessages:         ptr(true),
				CanPostStories:          ptr(true),
				CanPromoteMembers:       ptr(true),
				CanRestrictMembers:      ptr(true),
				IsAnonymous:             ptr(true),
			},
			expected: map[string]string{
				"user_id":                    "123",
				"can_change_info":            "1",
				"can_delete_messages":        "1",
				"can_delete_stories":         "1",
				"can_edit_messages":          "1",
				"can_edit_stories":           "1",
				"can_invite_users":           "1",
				"can_manage_chat":            "1",
				"can_manage_direct_messages": "1",
				"can_manage_topics":          "1",
				"can_manage_video_chats":     "1",
				"can_pin_messages":           "1",
				"can_post_messages":          "1",
				"can_post_stories":           "1",
				"can_promote_members":        "1",
				"can_restrict_members":       "1",
				"is_anonymous":               "1",
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

func TestPromoteChatMember_GetFiles(t *testing.T) {
	request := &PromoteChatMember{
		UserId: 123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
