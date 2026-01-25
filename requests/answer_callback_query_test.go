package requests

import (
	"testing"
)

func TestAnswerCallbackQuery_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *AnswerCallbackQuery
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &AnswerCallbackQuery{
				CallbackQueryId: "test_callback_query_id",
			},
			expected: map[string]string{
				"callback_query_id": "test_callback_query_id",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &AnswerCallbackQuery{
				CallbackQueryId: "test_callback_query_id",
				CacheTime:       ptr(int64(456)),
				ShowAlert:       ptr(true),
				Text:            ptr("test_text"),
				Url:             ptr("test_url"),
			},
			expected: map[string]string{
				"callback_query_id": "test_callback_query_id",
				"cache_time":        "456",
				"show_alert":        "1",
				"text":              "test_text",
				"url":               "test_url",
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

func TestAnswerCallbackQuery_GetFiles(t *testing.T) {
	request := &AnswerCallbackQuery{
		CallbackQueryId: "test_callback_query_id",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
