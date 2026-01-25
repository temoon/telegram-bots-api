package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendVenue_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendVenue
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendVenue{
				Address:   "test_address",
				ChatId:    telegram.NewChatId(123456, ""),
				Latitude:  123.45,
				Longitude: 123.45,
				Title:     "test_title",
			},
			expected: map[string]string{
				"address":   "test_address",
				"chat_id":   "123456",
				"latitude":  "123.45",
				"longitude": "123.45",
				"title":     "test_title",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendVenue{
				Address:               "test_address",
				Latitude:              123.45,
				Longitude:             123.45,
				Title:                 "test_title",
				AllowPaidBroadcast:    ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				DirectMessagesTopicId: ptr(int64(456)),
				DisableNotification:   ptr(true),
				FoursquareId:          ptr("test_foursquare_id"),
				FoursquareType:        ptr("test_foursquare_type"),
				GooglePlaceId:         ptr("test_google_place_id"),
				GooglePlaceType:       ptr("test_google_place_type"),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				ProtectContent:        ptr(true),
			},
			expected: map[string]string{
				"address":                  "test_address",
				"latitude":                 "123.45",
				"longitude":                "123.45",
				"title":                    "test_title",
				"allow_paid_broadcast":     "1",
				"business_connection_id":   "test_business_connection_id",
				"direct_messages_topic_id": "456",
				"disable_notification":     "1",
				"foursquare_id":            "test_foursquare_id",
				"foursquare_type":          "test_foursquare_type",
				"google_place_id":          "test_google_place_id",
				"google_place_type":        "test_google_place_type",
				"message_effect_id":        "test_message_effect_id",
				"message_thread_id":        "456",
				"protect_content":          "1",
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

func TestSendVenue_GetFiles(t *testing.T) {
	request := &SendVenue{
		Address:   "test_address",
		Latitude:  123.45,
		Longitude: 123.45,
		Title:     "test_title",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
