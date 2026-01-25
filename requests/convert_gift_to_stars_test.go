package requests

import (
	"testing"
)

func TestConvertGiftToStars_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *ConvertGiftToStars
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &ConvertGiftToStars{
				BusinessConnectionId: "test_business_connection_id",
				OwnedGiftId:          "test_owned_gift_id",
			},
			expected: map[string]string{
				"business_connection_id": "test_business_connection_id",
				"owned_gift_id":          "test_owned_gift_id",
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

func TestConvertGiftToStars_GetFiles(t *testing.T) {
	request := &ConvertGiftToStars{
		BusinessConnectionId: "test_business_connection_id",
		OwnedGiftId:          "test_owned_gift_id",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
