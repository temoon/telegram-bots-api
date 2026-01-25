package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestPinChatMessage_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *PinChatMessage
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &PinChatMessage{
				ChatId:    telegram.NewChatId(123456, ""),
				MessageId: 123,
			},
			expected: map[string]string{
				"chat_id":    "123456",
				"message_id": "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &PinChatMessage{
				MessageId:            123,
				BusinessConnectionId: ptr("test_business_connection_id"),
				DisableNotification:  ptr(true),
			},
			expected: map[string]string{
				"message_id":             "123",
				"business_connection_id": "test_business_connection_id",
				"disable_notification":   "1",
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

func TestPinChatMessage_GetFiles(t *testing.T) {
	request := &PinChatMessage{
		MessageId: 123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
