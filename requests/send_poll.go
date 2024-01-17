package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendPoll struct {
	Options               []string
	OpenPeriod            *int64
	AllowsMultipleAnswers *bool
	CorrectOptionId       *int64
	ExplanationEntities   []telegram.MessageEntity
	MessageThreadId       *int64
	Question              string
	Explanation           *string
	ExplanationParseMode  *string
	CloseDate             *int64
	ReplyMarkup           interface{}
	IsAnonymous           *bool
	Type                  *string
	DisableNotification   *bool
	ProtectContent        *bool
	ReplyParameters       *telegram.ReplyParameters
	ChatId                telegram.ChatId
	IsClosed              *bool
}

func (r *SendPoll) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendPoll", r, response)
	return
}

func (r *SendPoll) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataOptions []byte
	if dataOptions, err = json.Marshal(r.Options); err != nil {
		return
	}

	values["options"] = string(dataOptions)

	if r.OpenPeriod != nil {
		values["open_period"] = strconv.FormatInt(*r.OpenPeriod, 10)
	}

	if r.AllowsMultipleAnswers != nil {
		if *r.AllowsMultipleAnswers {
			values["allows_multiple_answers"] = "1"
		} else {
			values["allows_multiple_answers"] = "0"
		}
	}

	if r.CorrectOptionId != nil {
		values["correct_option_id"] = strconv.FormatInt(*r.CorrectOptionId, 10)
	}

	if r.ExplanationEntities != nil {
		var dataExplanationEntities []byte
		if dataExplanationEntities, err = json.Marshal(r.ExplanationEntities); err != nil {
			return
		}

		values["explanation_entities"] = string(dataExplanationEntities)
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	values["question"] = r.Question

	if r.Explanation != nil {
		values["explanation"] = *r.Explanation
	}

	if r.ExplanationParseMode != nil {
		values["explanation_parse_mode"] = *r.ExplanationParseMode
	}

	if r.CloseDate != nil {
		values["close_date"] = strconv.FormatInt(*r.CloseDate, 10)
	}

	if r.ReplyMarkup != nil {
		switch value := r.ReplyMarkup.(type) {
		case telegram.InlineKeyboardMarkup, telegram.ReplyKeyboardMarkup, telegram.ReplyKeyboardRemove, telegram.ForceReply:
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		default:
			err = errors.New("unsupported reply_markup field type")
			return
		}
	}

	if r.IsAnonymous != nil {
		if *r.IsAnonymous {
			values["is_anonymous"] = "1"
		} else {
			values["is_anonymous"] = "0"
		}
	}

	if r.Type != nil {
		values["type"] = *r.Type
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	values["chat_id"] = r.ChatId.String()

	if r.IsClosed != nil {
		if *r.IsClosed {
			values["is_closed"] = "1"
		} else {
			values["is_closed"] = "0"
		}
	}

	return
}

func (r *SendPoll) GetFiles() (files map[string]io.Reader) {
	return
}
