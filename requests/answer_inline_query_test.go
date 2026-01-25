package requests

import (
	"testing"
)

func TestAnswerInlineQuery_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *AnswerInlineQuery
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &AnswerInlineQuery{
				InlineQueryId: "test_inline_query_id",
			},
			expected: map[string]string{
				"inline_query_id": "test_inline_query_id",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &AnswerInlineQuery{
				InlineQueryId: "test_inline_query_id",
				CacheTime:     ptr(int64(456)),
				IsPersonal:    ptr(true),
				NextOffset:    ptr("test_next_offset"),
			},
			expected: map[string]string{
				"inline_query_id": "test_inline_query_id",
				"cache_time":      "456",
				"is_personal":     "1",
				"next_offset":     "test_next_offset",
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

func TestAnswerInlineQuery_GetFiles(t *testing.T) {
	request := &AnswerInlineQuery{
		InlineQueryId: "test_inline_query_id",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
