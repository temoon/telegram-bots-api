package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditMessageCaption struct {
	MessageId       *int64
	InlineMessageId *string
	Caption         *string
	ParseMode       *string
	CaptionEntities []telegram.MessageEntity
	ReplyMarkup     *telegram.InlineKeyboardMarkup
	ChatId          *telegram.ChatId
}

func (r *EditMessageCaption) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageCaption", r, response)
	return
}

func (r *EditMessageCaption) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	if r.Caption != nil {
		values["caption"] = *r.Caption
	}

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.CaptionEntities != nil {
		var dataCaptionEntities []byte
		if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(dataCaptionEntities)
	}

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	if r.ChatId != nil {
		values["chat_id"] = r.ChatId.String()
	}

	return
}

func (r *EditMessageCaption) GetFiles() (files map[string]io.Reader) {
	return
}
