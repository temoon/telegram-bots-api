package requests

import (
	"testing"
)

func TestGiftPremiumSubscription_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *GiftPremiumSubscription
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &GiftPremiumSubscription{
				MonthCount: 123,
				StarCount:  123,
				UserId:     123,
			},
			expected: map[string]string{
				"month_count": "123",
				"star_count":  "123",
				"user_id":     "123",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &GiftPremiumSubscription{
				MonthCount:    123,
				StarCount:     123,
				UserId:        123,
				Text:          ptr("test_text"),
				TextParseMode: ptr("test_text_parse_mode"),
			},
			expected: map[string]string{
				"month_count":     "123",
				"star_count":      "123",
				"user_id":         "123",
				"text":            "test_text",
				"text_parse_mode": "test_text_parse_mode",
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

func TestGiftPremiumSubscription_GetFiles(t *testing.T) {
	request := &GiftPremiumSubscription{
		MonthCount: 123,
		StarCount:  123,
		UserId:     123,
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
