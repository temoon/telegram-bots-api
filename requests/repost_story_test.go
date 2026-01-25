package requests

import (
	"testing"
)

func TestRepostStory_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *RepostStory
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &RepostStory{
				ActivePeriod:         123,
				BusinessConnectionId: "test_business_connection_id",
				FromChatId:           123,
				FromStoryId:          123,
			},
			expected: map[string]string{
				"active_period":          "123",
				"business_connection_id": "test_business_connection_id",
				"from_chat_id":           "123",
				"from_story_id":          "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &RepostStory{
				ActivePeriod:         123,
				BusinessConnectionId: "test_business_connection_id",
				FromChatId:           123,
				FromStoryId:          123,
				PostToChatPage:       ptr(true),
				ProtectContent:       ptr(true),
			},
			expected: map[string]string{
				"active_period":          "123",
				"business_connection_id": "test_business_connection_id",
				"from_chat_id":           "123",
				"from_story_id":          "123",
				"post_to_chat_page":      "1",
				"protect_content":        "1",
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

func TestRepostStory_GetFiles(t *testing.T) {
	request := &RepostStory{
		ActivePeriod:         123,
		BusinessConnectionId: "test_business_connection_id",
		FromChatId:           123,
		FromStoryId:          123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
