package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendGift_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendGift
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendGift{
				GiftId: "test_gift_id",
			},
			expected: map[string]string{
				"gift_id": "test_gift_id",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendGift{
				GiftId:        "test_gift_id",
				ChatId:        ptr(telegram.NewChatId(789, "")),
				PayForUpgrade: ptr(true),
				Text:          ptr("test_text"),
				TextParseMode: ptr("test_text_parse_mode"),
				UserId:        ptr(int64(456)),
			},
			expected: map[string]string{
				"gift_id":         "test_gift_id",
				"chat_id":         "789",
				"pay_for_upgrade": "1",
				"text":            "test_text",
				"text_parse_mode": "test_text_parse_mode",
				"user_id":         "456",
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

func TestSendGift_GetFiles(t *testing.T) {
	request := &SendGift{
		GiftId: "test_gift_id",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
