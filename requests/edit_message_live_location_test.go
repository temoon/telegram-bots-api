package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestEditMessageLiveLocation_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *EditMessageLiveLocation
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &EditMessageLiveLocation{
				Latitude:  123.45,
				Longitude: 123.45,
			},
			expected: map[string]string{
				"latitude":  "123.45",
				"longitude": "123.45",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &EditMessageLiveLocation{
				Latitude:             123.45,
				Longitude:            123.45,
				BusinessConnectionId: ptr("test_business_connection_id"),
				ChatId:               ptr(telegram.NewChatId(789, "")),
				Heading:              ptr(int64(456)),
				HorizontalAccuracy:   ptr(456.78),
				InlineMessageId:      ptr("test_inline_message_id"),
				LivePeriod:           ptr(int64(456)),
				MessageId:            ptr(int64(456)),
				ProximityAlertRadius: ptr(int64(456)),
			},
			expected: map[string]string{
				"latitude":               "123.45",
				"longitude":              "123.45",
				"business_connection_id": "test_business_connection_id",
				"chat_id":                "789",
				"heading":                "456",
				"horizontal_accuracy":    "456.78",
				"inline_message_id":      "test_inline_message_id",
				"live_period":            "456",
				"message_id":             "456",
				"proximity_alert_radius": "456",
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

func TestEditMessageLiveLocation_GetFiles(t *testing.T) {
	request := &EditMessageLiveLocation{
		Latitude:  123.45,
		Longitude: 123.45,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
