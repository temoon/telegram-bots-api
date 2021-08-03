package requests

import (
	"encoding/json"
	"strconv"
)

type SendPoll struct {
	AllowSendingWithoutReply bool
	AllowsMultipleAnswers    bool
	ChatId                   interface{}
	CloseDate                uint64
	CorrectOptionId          uint64
	DisableNotification      bool
	Explanation              string
	ExplanationEntities      []interface{}
	ExplanationParseMode     string
	IsAnonymous              bool
	IsClosed                 bool
	OpenPeriod               uint64
	Options                  []string
	Question                 string
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
	Type                     string
}

func (r *SendPoll) IsMultipart() bool {
	return false
}

func (r *SendPoll) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	if r.AllowsMultipleAnswers {
		values["allows_multiple_answers"] = "1"
	}

	switch value := r.ChatId.(type) {
	case uint64:
		values["chat_id"] = strconv.FormatUint(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.CloseDate != 0 {
		values["close_date"] = strconv.FormatUint(r.CloseDate, 10)
	}

	if r.CorrectOptionId != 0 {
		values["correct_option_id"] = strconv.FormatUint(r.CorrectOptionId, 10)
	}

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	if r.Explanation != "" {
		values["explanation"] = r.Explanation
	}

	if r.ExplanationEntities != nil {
		var data []byte
		if data, err = json.Marshal(r.ExplanationEntities); err != nil {
			return
		}

		values["explanation_entities"] = string(data)
	}

	if r.ExplanationParseMode != "" {
		values["explanation_parse_mode"] = r.ExplanationParseMode
	}

	if r.IsAnonymous {
		values["is_anonymous"] = "1"
	}

	if r.IsClosed {
		values["is_closed"] = "1"
	}

	if r.OpenPeriod != 0 {
		values["open_period"] = strconv.FormatUint(r.OpenPeriod, 10)
	}

	if r.Options != nil {
		var data []byte
		if data, err = json.Marshal(r.Options); err != nil {
			return
		}

		values["options"] = string(data)
	}

	values["question"] = r.Question

	switch value := r.ReplyMarkup.(type) {
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	default:
		if value != nil {
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		}
	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatUint(r.ReplyToMessageId, 10)
	}

	if r.Type != "" {
		values["type"] = r.Type
	}

	return
}
