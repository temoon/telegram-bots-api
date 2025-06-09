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
	ChatId                telegram.ChatId
	Options               []telegram.InputPollOption
	Question              string
	AllowPaidBroadcast    *bool
	AllowsMultipleAnswers *bool
	BusinessConnectionId  *string
	CloseDate             *int64
	CorrectOptionId       *int64
	DisableNotification   *bool
	Explanation           *string
	ExplanationEntities   []telegram.MessageEntity
	ExplanationParseMode  *string
	IsAnonymous           *bool
	IsClosed              *bool
	MessageEffectId       *string
	MessageThreadId       *int64
	OpenPeriod            *int64
	ProtectContent        *bool
	QuestionEntities      []telegram.MessageEntity
	QuestionParseMode     *string
	ReplyMarkup           interface{}
	ReplyParameters       *telegram.ReplyParameters
	Type                  *string
}

func (r *SendPoll) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendPoll", r, response)
	return
}

func (r *SendPoll) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	var dataOptions []byte
	if dataOptions, err = json.Marshal(r.Options); err != nil {
		return
	}

	values["options"] = string(dataOptions)

	values["question"] = r.Question

	if r.AllowPaidBroadcast != nil {
		if *r.AllowPaidBroadcast {
			values["allow_paid_broadcast"] = "1"
		} else {
			values["allow_paid_broadcast"] = "0"
		}
	}

	if r.AllowsMultipleAnswers != nil {
		if *r.AllowsMultipleAnswers {
			values["allows_multiple_answers"] = "1"
		} else {
			values["allows_multiple_answers"] = "0"
		}
	}

	if r.BusinessConnectionId != nil {
		values["business_connection_id"] = *r.BusinessConnectionId
	}

	if r.CloseDate != nil {
		values["close_date"] = strconv.FormatInt(*r.CloseDate, 10)
	}

	if r.CorrectOptionId != nil {
		values["correct_option_id"] = strconv.FormatInt(*r.CorrectOptionId, 10)
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

	if r.MessageEffectId != nil {
		values["message_effect_id"] = *r.MessageEffectId
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.OpenPeriod != nil {
		values["open_period"] = strconv.FormatInt(*r.OpenPeriod, 10)
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.QuestionEntities != nil {
		var dataQuestionEntities []byte
		if dataQuestionEntities, err = json.Marshal(r.QuestionEntities); err != nil {
			return
		}

		values["question_entities"] = string(dataQuestionEntities)
	}

	if r.QuestionParseMode != nil {
		values["question_parse_mode"] = *r.QuestionParseMode
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

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	if r.Type != nil {
		values["type"] = *r.Type
	}

	return
}

func (r *SendPoll) GetFiles() (files map[string]io.Reader) {
	return
}
