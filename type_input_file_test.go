package telegram

import (
	"bytes"
	"encoding/json"
	"io"
	"strings"
	"testing"
)

func TestNewInputFile(t *testing.T) {
	t.Run("with file ID", func(t *testing.T) {
		inputFile := NewInputFile("file_id_123", nil, "")

		if inputFile.String() != "file_id_123" {
			t.Errorf("expected 'file_id_123', got '%s'", inputFile.String())
		}

		if inputFile.HasFile() {
			t.Error("expected HasFile to return false for file_id")
		}
	})

	t.Run("with file upload", func(t *testing.T) {
		reader := bytes.NewReader([]byte("test data"))
		inputFile := NewInputFile("", reader, "test.jpg")

		if !inputFile.HasFile() {
			t.Error("expected HasFile to return true for file upload")
		}

		if !strings.HasPrefix(inputFile.String(), "attach://") {
			t.Errorf("expected attach:// prefix, got '%s'", inputFile.String())
		}

		if inputFile.GetFile() != reader {
			t.Error("expected GetFile to return the same reader")
		}
	})

	t.Run("with file ID and reader prefers file ID", func(t *testing.T) {
		reader := bytes.NewReader([]byte("test data"))
		inputFile := NewInputFile("file_id_123", reader, "test.jpg")

		if inputFile.String() != "file_id_123" {
			t.Errorf("expected file ID to take precedence, got '%s'", inputFile.String())
		}
	})

	t.Run("with reader but no filename", func(t *testing.T) {
		reader := bytes.NewReader([]byte("test data"))
		inputFile := NewInputFile("", reader, "")

		if inputFile.HasFile() {
			t.Error("expected HasFile to return false when filename is empty")
		}

		if inputFile.String() != "" {
			t.Errorf("expected empty string, got '%s'", inputFile.String())
		}
	})

	t.Run("with filename but no reader", func(t *testing.T) {
		inputFile := NewInputFile("", nil, "test.jpg")

		if inputFile.HasFile() {
			t.Error("expected HasFile to return false when reader is nil")
		}

		if inputFile.String() != "" {
			t.Errorf("expected empty string, got '%s'", inputFile.String())
		}
	})

	t.Run("empty input", func(t *testing.T) {
		inputFile := NewInputFile("", nil, "")

		if inputFile.String() != "" {
			t.Errorf("expected empty string, got '%s'", inputFile.String())
		}

		if inputFile.HasFile() {
			t.Error("expected HasFile to return false")
		}
	})
}

func TestInputFile_String(t *testing.T) {
	tests := []struct {
		name           string
		fileId         string
		file           io.Reader
		fileName       string
		expectedPrefix string
	}{
		{
			name:           "file ID",
			fileId:         "AgACAgIAAxkBAAI...",
			file:           nil,
			fileName:       "",
			expectedPrefix: "AgACAgIAAxkBAAI...",
		},
		{
			name:           "HTTP URL",
			fileId:         "https://example.com/photo.jpg",
			file:           nil,
			fileName:       "",
			expectedPrefix: "https://example.com/photo.jpg",
		},
		{
			name:           "file upload",
			fileId:         "",
			file:           bytes.NewReader([]byte("data")),
			fileName:       "photo.jpg",
			expectedPrefix: "attach://file_photo.jpg",
		},
		{
			name:           "empty",
			fileId:         "",
			file:           nil,
			fileName:       "",
			expectedPrefix: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile := NewInputFile(tt.fileId, tt.file, tt.fileName)
			result := inputFile.String()

			if result != tt.expectedPrefix {
				t.Errorf("expected '%s', got '%s'", tt.expectedPrefix, result)
			}
		})
	}
}

func TestInputFile_GetFormFieldName(t *testing.T) {
	reader := bytes.NewReader([]byte("test data"))
	inputFile := NewInputFile("", reader, "photo.jpg")

	fieldName := inputFile.GetFormFieldName()

	if fieldName != "file_photo.jpg" {
		t.Errorf("expected 'file_photo.jpg', got '%s'", fieldName)
	}
}

func TestInputFile_MarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		fileId   string
		file     io.Reader
		fileName string
		expected string
	}{
		{
			name:     "file ID",
			fileId:   "file_id_123",
			file:     nil,
			fileName: "",
			expected: `"file_id_123"`,
		},
		{
			name:     "file upload",
			fileId:   "",
			file:     bytes.NewReader([]byte("data")),
			fileName: "test.jpg",
			expected: `"attach://file_test.jpg"`,
		},
		{
			name:     "empty",
			fileId:   "",
			file:     nil,
			fileName: "",
			expected: `""`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputFile := NewInputFile(tt.fileId, tt.file, tt.fileName)
			data, err := json.Marshal(inputFile)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if string(data) != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, string(data))
			}
		})
	}
}

func TestInputFile_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		expected string
	}{
		{
			name:     "file ID",
			json:     `"file_id_123"`,
			expected: "file_id_123",
		},
		{
			name:     "HTTP URL",
			json:     `"https://example.com/file.jpg"`,
			expected: "https://example.com/file.jpg",
		},
		{
			name:     "empty string",
			json:     `""`,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var inputFile InputFile
			err := json.Unmarshal([]byte(tt.json), &inputFile)

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if inputFile.String() != tt.expected {
				t.Errorf("expected '%s', got '%s'", tt.expected, inputFile.String())
			}
		})
	}
}

func TestInputFile_UnmarshalJSON_Invalid(t *testing.T) {
	tests := []struct {
		name string
		json string
	}{
		{
			name: "object",
			json: `{"file_id": "123"}`,
		},
		{
			name: "array",
			json: `["file_id"]`,
		},
		{
			name: "number",
			json: `123`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var inputFile InputFile
			err := json.Unmarshal([]byte(tt.json), &inputFile)

			if err == nil {
				t.Error("expected error for invalid JSON, got nil")
			}
		})
	}
}

func TestInputFile_RoundTrip(t *testing.T) {
	t.Run("file ID", func(t *testing.T) {
		original := NewInputFile("file_id_123", nil, "")

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("marshal error: %v", err)
		}

		var decoded InputFile
		err = json.Unmarshal(data, &decoded)
		if err != nil {
			t.Fatalf("unmarshal error: %v", err)
		}

		if decoded.String() != original.String() {
			t.Errorf("expected '%s', got '%s'", original.String(), decoded.String())
		}
	})
}

func TestInputFile_ReaderReuse(t *testing.T) {
	t.Run("reader can only be read once", func(t *testing.T) {
		data := []byte("test file content")
		reader := bytes.NewReader(data)
		inputFile := NewInputFile("", reader, "test.txt")

		// First read
		firstRead, err := io.ReadAll(inputFile.GetFile())
		if err != nil {
			t.Fatalf("first read error: %v", err)
		}

		if !bytes.Equal(firstRead, data) {
			t.Error("first read did not match original data")
		}

		// Second read should return empty (reader is exhausted)
		secondRead, err := io.ReadAll(inputFile.GetFile())
		if err != nil {
			t.Fatalf("second read error: %v", err)
		}

		if len(secondRead) != 0 {
			t.Errorf("expected second read to be empty, got %d bytes", len(secondRead))
		}

		// This test documents the current behavior (bug)
		// In a fixed version, this test should verify that both reads return the same data
	})
}

func TestInputFile_GetFile(t *testing.T) {
	t.Run("returns nil for file ID", func(t *testing.T) {
		inputFile := NewInputFile("file_id_123", nil, "")

		if inputFile.GetFile() != nil {
			t.Error("expected GetFile to return nil for file_id")
		}
	})

	t.Run("returns reader for file upload", func(t *testing.T) {
		reader := bytes.NewReader([]byte("test"))
		inputFile := NewInputFile("", reader, "test.txt")

		if inputFile.GetFile() != reader {
			t.Error("expected GetFile to return the same reader")
		}
	})
}
