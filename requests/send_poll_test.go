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
				Question:               "test_question",
				AllowAddingOptions:     ptr(true),
				AllowPaidBroadcast:     ptr(true),
				AllowsMultipleAnswers:  ptr(true),
				AllowsRevoting:         ptr(true),
				BusinessConnectionId:   ptr("test_business_connection_id"),
				CloseDate:              ptr(int64(456)),
				Description:            ptr("test_description"),
				DescriptionParseMode:   ptr("test_description_parse_mode"),
				DisableNotification:    ptr(true),
				Explanation:            ptr("test_explanation"),
				ExplanationParseMode:   ptr("test_explanation_parse_mode"),
				HideResultsUntilCloses: ptr(true),
				IsAnonymous:            ptr(true),
				IsClosed:               ptr(true),
				MessageEffectId:        ptr("test_message_effect_id"),
				MessageThreadId:        ptr(int64(456)),
				OpenPeriod:             ptr(int64(456)),
				ProtectContent:         ptr(true),
				QuestionParseMode:      ptr("test_question_parse_mode"),
				ShuffleOptions:         ptr(true),
				Type:                   ptr("test_type"),
			},
			expected: map[string]string{
				"question":                  "test_question",
				"allow_adding_options":      "1",
				"allow_paid_broadcast":      "1",
				"allows_multiple_answers":   "1",
				"allows_revoting":           "1",
				"business_connection_id":    "test_business_connection_id",
				"close_date":                "456",
				"description":               "test_description",
				"description_parse_mode":    "test_description_parse_mode",
				"disable_notification":      "1",
				"explanation":               "test_explanation",
				"explanation_parse_mode":    "test_explanation_parse_mode",
				"hide_results_until_closes": "1",
				"is_anonymous":              "1",
				"is_closed":                 "1",
				"message_effect_id":         "test_message_effect_id",
				"message_thread_id":         "456",
				"open_period":               "456",
				"protect_content":           "1",
				"question_parse_mode":       "test_question_parse_mode",
				"shuffle_options":           "1",
				"type":                      "test_type",
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
