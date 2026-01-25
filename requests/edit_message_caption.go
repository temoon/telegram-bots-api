package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type EditMessageCaption struct {
	BusinessConnectionId  *string
	Caption               *string
	CaptionEntities       []telegram.MessageEntity
	ChatId                *telegram.ChatId
	InlineMessageId       *string
	MessageId             *int64
	ParseMode             *string
	ReplyMarkup           *telegram.InlineKeyboardMarkup
	ShowCaptionAboveMedia *bool
}

func (r *EditMessageCaption) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageCaption", r, response)
	return
}

func (r *EditMessageCaption) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	if r.BusinessConnectionId != nil {
		values["business_connection_id"] = *r.BusinessConnectionId
	}

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
		values["chat_id"] = r.ChatId.String()
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

	if r.ShowCaptionAboveMedia != nil {
		if *r.ShowCaptionAboveMedia {
			values["show_caption_above_media"] = "1"
		} else {
			values["show_caption_above_media"] = "0"
		}
	}

	return
}

func (r *EditMessageCaption) GetFiles() (files map[string]io.Reader) {
	return
}
