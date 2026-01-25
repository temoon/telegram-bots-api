package requests

import (
	"testing"

	"github.com/temoon/telegram-bots-api"
)

func TestSendPoll_GetValues(t *testing.T) {
	tests := []struct {
		name     string
		request  *SendPoll
		expected map[string]string
		wantErr  bool
	}{
		{
			name: "required fields only",
			request: &SendPoll{
				ChatId:   telegram.NewChatId(123456, ""),
				Question: "test_question",
			},
			expected: map[string]string{
				"chat_id":  "123456",
				"question": "test_question",
			},
			wantErr: false,
		},
		{
			name: "with optional fields",
			request: &SendPoll{
				Question:              "test_question",
				AllowPaidBroadcast:    ptr(true),
				AllowsMultipleAnswers: ptr(true),
				BusinessConnectionId:  ptr("test_business_connection_id"),
				CloseDate:             ptr(int64(456)),
				CorrectOptionId:       ptr(int64(456)),
				DisableNotification:   ptr(true),
				Explanation:           ptr("test_explanation"),
				ExplanationParseMode:  ptr("test_explanation_parse_mode"),
				IsAnonymous:           ptr(true),
				IsClosed:              ptr(true),
				MessageEffectId:       ptr("test_message_effect_id"),
				MessageThreadId:       ptr(int64(456)),
				OpenPeriod:            ptr(int64(456)),
				ProtectContent:        ptr(true),
				QuestionParseMode:     ptr("test_question_parse_mode"),
				Type:                  ptr("test_type"),
			},
			expected: map[string]string{
				"question":                "test_question",
				"allow_paid_broadcast":    "1",
				"allows_multiple_answers": "1",
				"business_connection_id":  "test_business_connection_id",
				"close_date":              "456",
				"correct_option_id":       "456",
				"disable_notification":    "1",
				"explanation":             "test_explanation",
				"explanation_parse_mode":  "test_explanation_parse_mode",
				"is_anonymous":            "1",
				"is_closed":               "1",
				"message_effect_id":       "test_message_effect_id",
				"message_thread_id":       "456",
				"open_period":             "456",
				"protect_content":         "1",
				"question_parse_mode":     "test_question_parse_mode",
				"type":                    "test_type",
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

func TestSendPoll_GetFiles(t *testing.T) {
	request := &SendPoll{
		Question: "test_question",
	}

	files := request.GetFiles()

	if files != nil && len(files) != 0 {
		t.Errorf("expected nil or empty files for request without file fields")
	}
}
