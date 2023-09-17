package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type SendPoll struct {
	AllowSendingWithoutReply *bool
	AllowsMultipleAnswers    *bool
	ChatId                   interface{}
	CloseDate                *int32
	CorrectOptionId          *int32
	DisableNotification      *bool
	Explanation              *string
	ExplanationEntities      []telegram.MessageEntity
	ExplanationParseMode     *string
	IsAnonymous              *bool
	IsClosed                 *bool
	MessageThreadId          *int32
	OpenPeriod               *int32
	Options                  []string
	ProtectContent           *bool
	Question                 string
	ReplyMarkup              interface{}
	ReplyToMessageId         *int32
	Type                     *string
}

func (r *SendPoll) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendPoll", r, response)
	return
}

func (r *SendPoll) IsMultipart() (multipart bool) {
	return false
}

func (r *SendPoll) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply != nil {
		if *r.AllowSendingWithoutReply {
			values["allow_sending_without_reply"] = "1"
		} else {
			values["allow_sending_without_reply"] = "0"
		}
	}

	if r.AllowsMultipleAnswers != nil {
		if *r.AllowsMultipleAnswers {
			values["allows_multiple_answers"] = "1"
		} else {
			values["allows_multiple_answers"] = "0"
		}
	}

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.CloseDate != nil {
		values["close_date"] = strconv.FormatInt(int64(*r.CloseDate), 10)
	}

	if r.CorrectOptionId != nil {
		values["correct_option_id"] = strconv.FormatInt(int64(*r.CorrectOptionId), 10)
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.Explanation != nil {
		values["explanation"] = *r.Explanation
	}

	if r.ExplanationEntities != nil {
		var dataExplanationEntities []byte
		if dataExplanationEntities, err = json.Marshal(r.ExplanationEntities); err != nil {
			return
		}

		values["explanation_entities"] = string(dataExplanationEntities)
	}

	if r.ExplanationParseMode != nil {
		values["explanation_parse_mode"] = *r.ExplanationParseMode
	}

	if r.IsAnonymous != nil {
		if *r.IsAnonymous {
			values["is_anonymous"] = "1"
		} else {
			values["is_anonymous"] = "0"
		}
	}

	if r.IsClosed != nil {
		if *r.IsClosed {
			values["is_closed"] = "1"
		} else {
			values["is_closed"] = "0"
		}
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(int64(*r.MessageThreadId), 10)
	}

	if r.OpenPeriod != nil {
		values["open_period"] = strconv.FormatInt(int64(*r.OpenPeriod), 10)
	}

	var dataOptions []byte
	if dataOptions, err = json.Marshal(r.Options); err != nil {
		return
	}

	values["options"] = string(dataOptions)

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	values["question"] = r.Question

	switch value := r.ReplyMarkup.(type) {
	case *telegram.InlineKeyboardMarkup:
		if value != nil {
			var dataInlineKeyboardMarkup []byte
			if dataInlineKeyboardMarkup, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(dataInlineKeyboardMarkup)
		}
	case *telegram.ReplyKeyboardMarkup:
		if value != nil {
			var dataReplyKeyboardMarkup []byte
			if dataReplyKeyboardMarkup, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(dataReplyKeyboardMarkup)
		}
	case *telegram.ReplyKeyboardRemove:
		if value != nil {
			var dataReplyKeyboardRemove []byte
			if dataReplyKeyboardRemove, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(dataReplyKeyboardRemove)
		}
	case *telegram.ForceReply:
		if value != nil {
			var dataForceReply []byte
			if dataForceReply, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(dataForceReply)
		}
	}

	if r.ReplyToMessageId != nil {
		values["reply_to_message_id"] = strconv.FormatInt(int64(*r.ReplyToMessageId), 10)
	}

	if r.Type != nil {
		values["type"] = *r.Type
	}

	return
}
