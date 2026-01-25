package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendInvoice_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendInvoice
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendInvoice{
				ChatId:      telegram.NewChatId(123456, ""),
				Currency:    "test_currency",
				Description: "test_description",
				Payload:     "test_payload",
				Title:       "test_title",
			},
			expected: map[string]string{
				"chat_id":     "123456",
				"currency":    "test_currency",
				"description": "test_description",
				"payload":     "test_payload",
				"title":       "test_title",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendInvoice{
				Currency:                  "test_currency",
				Description:               "test_description",
				Payload:                   "test_payload",
				Title:                     "test_title",
				AllowPaidBroadcast:        ptr(true),
				DirectMessagesTopicId:     ptr(int64(456)),
				DisableNotification:       ptr(true),
				IsFlexible:                ptr(true),
				MaxTipAmount:              ptr(int64(456)),
				MessageEffectId:           ptr("test_message_effect_id"),
				MessageThreadId:           ptr(int64(456)),
				NeedEmail:                 ptr(true),
				NeedName:                  ptr(true),
				NeedPhoneNumber:           ptr(true),
				NeedShippingAddress:       ptr(true),
				PhotoHeight:               ptr(int64(456)),
				PhotoSize:                 ptr(int64(456)),
				PhotoUrl:                  ptr("test_photo_url"),
				PhotoWidth:                ptr(int64(456)),
				ProtectContent:            ptr(true),
				ProviderData:              ptr("test_provider_data"),
				ProviderToken:             ptr("test_provider_token"),
				SendEmailToProvider:       ptr(true),
				SendPhoneNumberToProvider: ptr(true),
				StartParameter:            ptr("test_start_parameter"),
			},
			expected: map[string]string{
				"currency":                      "test_currency",
				"description":                   "test_description",
				"payload":                       "test_payload",
				"title":                         "test_title",
				"allow_paid_broadcast":          "1",
				"direct_messages_topic_id":      "456",
				"disable_notification":          "1",
				"is_flexible":                   "1",
				"max_tip_amount":                "456",
				"message_effect_id":             "test_message_effect_id",
				"message_thread_id":             "456",
				"need_email":                    "1",
				"need_name":                     "1",
				"need_phone_number":             "1",
				"need_shipping_address":         "1",
				"photo_height":                  "456",
				"photo_size":                    "456",
				"photo_url":                     "test_photo_url",
				"photo_width":                   "456",
				"protect_content":               "1",
				"provider_data":                 "test_provider_data",
				"provider_token":                "test_provider_token",
				"send_email_to_provider":        "1",
				"send_phone_number_to_provider": "1",
				"start_parameter":               "test_start_parameter",
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

func TestSendInvoice_GetFiles(t *testing.T) {
	request := &SendInvoice{
		Currency:    "test_currency",
		Description: "test_description",
		Payload:     "test_payload",
		Title:       "test_title",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
