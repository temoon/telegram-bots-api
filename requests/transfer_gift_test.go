package requests

import (
	"testing"
)

func TestTransferGift_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *TransferGift
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &TransferGift{
				BusinessConnectionId: "test_business_connection_id",
				NewOwnerChatId:       123,
				OwnedGiftId:          "test_owned_gift_id",
			},
			expected: map[string]string{
				"business_connection_id": "test_business_connection_id",
				"new_owner_chat_id":      "123",
				"owned_gift_id":          "test_owned_gift_id",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &TransferGift{
				BusinessConnectionId: "test_business_connection_id",
				NewOwnerChatId:       123,
				OwnedGiftId:          "test_owned_gift_id",
				StarCount:            ptr(int64(456)),
			},
			expected: map[string]string{
				"business_connection_id": "test_business_connection_id",
				"new_owner_chat_id":      "123",
				"owned_gift_id":          "test_owned_gift_id",
				"star_count":             "456",
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

func TestTransferGift_GetFiles(t *testing.T) {
	request := &TransferGift{
		BusinessConnectionId: "test_business_connection_id",
		NewOwnerChatId:       123,
		OwnedGiftId:          "test_owned_gift_id",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
