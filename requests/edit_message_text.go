package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type EditMessageText struct {
	ChatId                interface{}
	DisableWebPagePreview *bool
	Entities              []telegram.MessageEntity
	InlineMessageId       *string
	MessageId             *int32
	ParseMode             *string
	ReplyMarkup           *telegram.InlineKeyboardMarkup
	Text                  string
}

func (r *EditMessageText) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(interface{})
	err = b.CallMethod(ctx, "editMessageText", r, response)
	return
}

func (r *EditMessageText) IsMultipart() (multipart bool) {
	return false
}

func (r *EditMessageText) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableWebPagePreview != nil {
		if *r.DisableWebPagePreview {
			values["disable_web_page_preview"] = "1"
		} else {
			values["disable_web_page_preview"] = "0"
		}
	}

	if r.Entities != nil {
		var dataEntities []byte
		if dataEntities, err = json.Marshal(r.Entities); err != nil {
			return
		}

		values["entities"] = string(dataEntities)
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(int64(*r.MessageId), 10)
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

	values["text"] = r.Text

	return
}
