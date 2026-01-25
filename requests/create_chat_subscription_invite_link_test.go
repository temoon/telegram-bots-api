package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestCreateChatSubscriptionInviteLink_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *CreateChatSubscriptionInviteLink
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &CreateChatSubscriptionInviteLink{
				ChatId:             telegram.NewChatId(123456, ""),
				SubscriptionPeriod: 123,
				SubscriptionPrice:  123,
			},
			expected: map[string]string{
				"chat_id":             "123456",
				"subscription_period": "123",
				"subscription_price":  "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &CreateChatSubscriptionInviteLink{
				SubscriptionPeriod: 123,
				SubscriptionPrice:  123,
				Name:               ptr("test_name"),
			},
			expected: map[string]string{
				"subscription_period": "123",
				"subscription_price":  "123",
				"name":                "test_name",
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

func TestCreateChatSubscriptionInviteLink_GetFiles(t *testing.T) {
	request := &CreateChatSubscriptionInviteLink{
		SubscriptionPeriod: 123,
		SubscriptionPrice:  123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
