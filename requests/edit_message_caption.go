package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type EditMessageCaption struct {
	Caption         *string
	CaptionEntities []telegram.MessageEntity
	ChatId          interface{}
	InlineMessageId *string
	MessageId       *int64
	ParseMode       *string
	ReplyMarkup     *telegram.InlineKeyboardMarkup
}

func (r *EditMessageCaption) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageCaption", r, response)
	return
}

func (r *EditMessageCaption) IsMultipart() bool {
	return false
}

func (r *EditMessageCaption) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Caption != nil {
		values["caption"] = *r.Caption
	}

	if r.CaptionEntities != nil {
		var dataCaptionEntities []byte
		if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(dataCaptionEntities)
	}

	if r.ChatId != nil {
		switch value := r.ChatId.(type) {
		case int64:
			values["chat_id"] = strconv.FormatInt(value, 10)
		case string:
			values["chat_id"] = value
		default:
			err = errors.New("invalid chat_id field type")
			return
		}
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	return
}
