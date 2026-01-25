package requests

import (
	"testing"
)

func TestAnswerShippingQuery_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *AnswerShippingQuery
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &AnswerShippingQuery{
				Ok:              true,
				ShippingQueryId: "test_shipping_query_id",
			},
			expected: map[string]string{
				"ok":                "1",
				"shipping_query_id": "test_shipping_query_id",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &AnswerShippingQuery{
				Ok:              true,
				ShippingQueryId: "test_shipping_query_id",
				ErrorMessage:    ptr("test_error_message"),
			},
			expected: map[string]string{
				"ok":                "1",
				"shipping_query_id": "test_shipping_query_id",
				"error_message":     "test_error_message",
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

func TestAnswerShippingQuery_GetFiles(t *testing.T) {
	request := &AnswerShippingQuery{
		Ok:              true,
		ShippingQueryId: "test_shipping_query_id",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
