package requests

import (
	"testing"
)

func TestSendChecklist_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendChecklist
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendChecklist{
				BusinessConnectionId: "test_business_connection_id",
				ChatId:               123,
			},
			expected: map[string]string{
				"business_connection_id": "test_business_connection_id",
				"chat_id":                "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendChecklist{
				BusinessConnectionId: "test_business_connection_id",
				ChatId:               123,
				DisableNotification:  ptr(true),
				MessageEffectId:      ptr("test_message_effect_id"),
				ProtectContent:       ptr(true),
			},
			expected: map[string]string{
				"business_connection_id": "test_business_connection_id",
				"chat_id":                "123",
				"disable_notification":   "1",
				"message_effect_id":      "test_message_effect_id",
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

func TestSendChecklist_GetFiles(t *testing.T) {
	request := &SendChecklist{
		BusinessConnectionId: "test_business_connection_id",
		ChatId:               123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
