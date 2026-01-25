package telegram

import (
	"encoding/json"
	"testing"
)

func TestNewChatId(t *testing.T) {
	t.Run("with numeric ID", func(t *testing.T) {
		chatId := NewChatId(123456, "")

		if chatId.GetId() != 123456 {
			t.Errorf("expected ID 123456, got %d", chatId.GetId())
		}

		if chatId.GetName() != "" {
			t.Errorf("expected empty name, got '%s'", chatId.GetName())
		}
	})

	t.Run("with username", func(t *testing.T) {
		chatId := NewChatId(0, "@testuser")

		if chatId.GetId() != 0 {
			t.Errorf("expected ID 0, got %d", chatId.GetId())
		}

		if chatId.GetName() != "@testuser" {
			t.Errorf("expected name '@testuser', got '%s'", chatId.GetName())
		}
	})

	t.Run("with both ID and name prefers ID", func(t *testing.T) {
		chatId := NewChatId(123456, "@testuser")

		if chatId.GetId() != 123456 {
			t.Errorf("expected ID 123456, got %d", chatId.GetId())
		}

		// Name should not be set when ID is provided
		if chatId.GetName() != "" {
			t.Errorf("expected empty name when ID is provided, got '%s'", chatId.GetName())
		}
	})

	t.Run("with neither ID nor name", func(t *testing.T) {
		chatId := NewChatId(0, "")

		if chatId.GetId() != 0 {
			t.Errorf("expected ID 0, got %d", chatId.GetId())
		}

		if chatId.GetName() != "" {
			t.Errorf("expected empty name, got '%s'", chatId.GetName())
		}
	})

	t.Run("with negative ID", func(t *testing.T) {
		chatId := NewChatId(-123456, "")

		if chatId.GetId() != -123456 {
			t.Errorf("expected ID -123456, got %d", chatId.GetId())
		}
	})
}

func TestChatId_String(t *testing.T) {
	tests := []struct {
		name     string
		id       int64
		username string
		expected string
	}{
		{
			name:     "numeric ID",
			id:       123456,
			username: "",
			expected: "123456",
		},
		{
			name:     "negative ID",
			id:       -123456,
			username: "",
			expected: "-123456",
		},
		{
			name:     "username",
			id:       0,
			username: "@testuser",
			expected: "@testuser",
		},
		{
			name:     "empty",
			id:       0,
			username: "",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chatId := NewChatId(tt.id, tt.username)
			result := chatId.String()

			if result != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, result)
			}
		})
	}
}

func TestChatId_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		id       int64
		username string
		expected string
	}{
		{
			name:     "numeric ID",
			id:       123456,
			username: "",
			expected: `"123456"`,
		},
		{
			name:     "username",
			id:       0,
			username: "@testuser",
			expected: `"@testuser"`,
		},
		{
			name:     "empty",
			id:       0,
			username: "",
			expected: `""`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chatId := NewChatId(tt.id, tt.username)
			data, err := json.Marshal(chatId)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(data) != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, string(data))
			}
		})
	}
}

func TestChatId_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name         string
		json         string
		expectedId   int64
		expectedName string
	}{
		{
			name:         "numeric ID",
			json:         `"123456"`,
			expectedId:   123456,
			expectedName: "",
		},
		{
			name:         "numeric ID without quotes",
			json:         `123456`,
			expectedId:   123456,
			expectedName: "",
		},
		{
			name:         "negative ID",
			json:         `"-123456"`,
			expectedId:   -123456,
			expectedName: "",
		},
		{
			name:         "username",
			json:         `"@testuser"`,
			expectedId:   0,
			expectedName: "@testuser",
		},
		{
			name:         "empty string",
			json:         `""`,
			expectedId:   0,
			expectedName: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var chatId ChatId
			err := json.Unmarshal([]byte(tt.json), &chatId)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if chatId.GetId() != tt.expectedId {
				t.Errorf("expected ID %d, got %d", tt.expectedId, chatId.GetId())
			}

			if chatId.GetName() != tt.expectedName {
				t.Errorf("expected name '%s', got '%s'", tt.expectedName, chatId.GetName())
			}
		})
	}
}

func TestChatId_UnmarshalJSON_Invalid(t *testing.T) {
	tests := []struct {
		name string
		json string
	}{
		{
			name: "invalid number",
			json: `"abc"`,
		},
		{
			name: "object",
			json: `{"id": 123}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var chatId ChatId
			err := json.Unmarshal([]byte(tt.json), &chatId)

			if err == nil {
				t.Error("expected error for invalid JSON, got nil")
			}
		})
	}
}

func TestChatId_RoundTrip(t *testing.T) {
	tests := []struct {
		name     string
		id       int64
		username string
	}{
		{name: "numeric ID", id: 123456, username: ""},
		{name: "username", id: 0, username: "@testuser"},
		{name: "negative ID", id: -100, username: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			original := NewChatId(tt.id, tt.username)

			data, err := json.Marshal(original)
			if err != nil {
				t.Fatalf("marshal error: %v", err)
			}

			var decoded ChatId
			err = json.Unmarshal(data, &decoded)
			if err != nil {
				t.Fatalf("unmarshal error: %v", err)
			}

			if decoded.GetId() != original.GetId() {
				t.Errorf("ID mismatch: expected %d, got %d", original.GetId(), decoded.GetId())
			}

			if decoded.GetName() != original.GetName() {
				t.Errorf("Name mismatch: expected '%s', got '%s'", original.GetName(), decoded.GetName())
			}
		})
	}
}
