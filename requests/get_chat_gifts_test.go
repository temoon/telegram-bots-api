package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestGetChatGifts_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *GetChatGifts
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &GetChatGifts{
				ChatId: telegram.NewChatId(123456, ""),
			},
			expected: map[string]string{
				"chat_id": "123456",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &GetChatGifts{
				ExcludeFromBlockchain:       ptr(true),
				ExcludeLimitedNonUpgradable: ptr(true),
				ExcludeLimitedUpgradable:    ptr(true),
				ExcludeSaved:                ptr(true),
				ExcludeUnique:               ptr(true),
				ExcludeUnlimited:            ptr(true),
				ExcludeUnsaved:              ptr(true),
				Limit:                       ptr(int64(456)),
				Offset:                      ptr("test_offset"),
				SortByPrice:                 ptr(true),
			},
			expected: map[string]string{
				"exclude_from_blockchain":        "1",
				"exclude_limited_non_upgradable": "1",
				"exclude_limited_upgradable":     "1",
				"exclude_saved":                  "1",
				"exclude_unique":                 "1",
				"exclude_unlimited":              "1",
				"exclude_unsaved":                "1",
				"limit":                          "456",
				"offset":                         "test_offset",
				"sort_by_price":                  "1",
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

func TestGetChatGifts_GetFiles(t *testing.T) {
	request := &GetChatGifts{}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
