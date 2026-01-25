package requests

import (
	"testing"
)

func TestSendGame_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendGame
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendGame{
				ChatId:        123,
				GameShortName: "test_game_short_name",
			},
			expected: map[string]string{
				"chat_id":         "123",
				"game_short_name": "test_game_short_name",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendGame{
				ChatId:               123,
				GameShortName:        "test_game_short_name",
				AllowPaidBroadcast:   ptr(true),
				BusinessConnectionId: ptr("test_business_connection_id"),
				DisableNotification:  ptr(true),
				MessageEffectId:      ptr("test_message_effect_id"),
				MessageThreadId:      ptr(int64(456)),
				ProtectContent:       ptr(true),
			},
			expected: map[string]string{
				"chat_id":                "123",
				"game_short_name":        "test_game_short_name",
				"allow_paid_broadcast":   "1",
				"business_connection_id": "test_business_connection_id",
				"disable_notification":   "1",
				"message_effect_id":      "test_message_effect_id",
				"message_thread_id":      "456",
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

func TestSendGame_GetFiles(t *testing.T) {
	request := &SendGame{
		ChatId:        123,
		GameShortName: "test_game_short_name",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
