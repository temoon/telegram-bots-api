package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendLocation_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendLocation
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendLocation{
				ChatId:    telegram.NewChatId(123456, ""),
				Latitude:  123.45,
				Longitude: 123.45,
			},
			expected: map[string]string{
				"chat_id":   "123456",
				"latitude":  "123.45",
				"longitude": "123.45",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendLocation{
				Latitude:              123.45,
				Longitude:             123.45,
				AllowPaidBroadcast:    ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				Heading:               ptr(int64(456)),
				HorizontalAccuracy:    ptr(456.78),
				LivePeriod:            ptr(int64(456)),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				ProtectContent:        ptr(true),
				ProximityAlertRadius:  ptr(int64(456)),
			},
			expected: map[string]string{
				"latitude":                 "123.45",
				"longitude":                "123.45",
				"allow_paid_broadcast":     "1",
				"business_connection_id":   "test_business_connection_id",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"heading":                  "456",
				"horizontal_accuracy":      "456.78",
				"live_period":              "456",
				"message_effect_id":        "test_message_effect_id",
				"message_thread_id":        "456",
				"protect_content":          "1",
				"proximity_alert_radius":   "456",
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

func TestSendLocation_GetFiles(t *testing.T) {
	request := &SendLocation{
		Latitude:  123.45,
		Longitude: 123.45,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
