package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendContact_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendContact
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendContact{
				ChatId:      telegram.NewChatId(123456, ""),
				FirstName:   "test_first_name",
				PhoneNumber: "test_phone_number",
			},
			expected: map[string]string{
				"chat_id":      "123456",
				"first_name":   "test_first_name",
				"phone_number": "test_phone_number",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendContact{
				FirstName:             "test_first_name",
				PhoneNumber:           "test_phone_number",
				AllowPaidBroadcast:    ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				LastName:              ptr("test_last_name"),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				ProtectContent:        ptr(true),
				Vcard:                 ptr("test_vcard"),
			},
			expected: map[string]string{
				"first_name":               "test_first_name",
				"phone_number":             "test_phone_number",
				"allow_paid_broadcast":     "1",
				"business_connection_id":   "test_business_connection_id",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"last_name":                "test_last_name",
				"message_effect_id":        "test_message_effect_id",
				"message_thread_id":        "456",
				"protect_content":          "1",
				"vcard":                    "test_vcard",
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

func TestSendContact_GetFiles(t *testing.T) {
	request := &SendContact{
		FirstName:   "test_first_name",
		PhoneNumber: "test_phone_number",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
