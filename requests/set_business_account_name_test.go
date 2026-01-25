package requests

import (
	"testing"
)

func TestSetBusinessAccountName_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SetBusinessAccountName
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SetBusinessAccountName{
				BusinessConnectionId: "test_business_connection_id",
				FirstName:            "test_first_name",
			},
			expected: map[string]string{
				"business_connection_id": "test_business_connection_id",
				"first_name":             "test_first_name",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SetBusinessAccountName{
				BusinessConnectionId: "test_business_connection_id",
				FirstName:            "test_first_name",
				LastName:             ptr("test_last_name"),
			},
			expected: map[string]string{
				"business_connection_id": "test_business_connection_id",
				"first_name":             "test_first_name",
				"last_name":              "test_last_name",
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

func TestSetBusinessAccountName_GetFiles(t *testing.T) {
	request := &SetBusinessAccountName{
		BusinessConnectionId: "test_business_connection_id",
		FirstName:            "test_first_name",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
