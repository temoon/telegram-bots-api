package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditMessageText struct {
	InlineMessageId    *string
	Text               string
	ParseMode          *string
	Entities           []telegram.MessageEntity
	LinkPreviewOptions *telegram.LinkPreviewOptions
	ReplyMarkup        *telegram.InlineKeyboardMarkup
	ChatId             *telegram.ChatId
	MessageId          *int64
}

func (r *EditMessageText) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageText", r, response)
	return
}

func (r *EditMessageText) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	values["text"] = r.Text

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.Entities != nil {
		var dataEntities []byte
		if dataEntities, err = json.Marshal(r.Entities); err != nil {
			return
		}

		values["entities"] = string(dataEntities)
	}

	if r.LinkPreviewOptions != nil {
		var dataLinkPreviewOptions []byte
		if dataLinkPreviewOptions, err = json.Marshal(r.LinkPreviewOptions); err != nil {
			return
		}

		values["link_preview_options"] = string(dataLinkPreviewOptions)
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

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	return
}

func (r *EditMessageText) GetFiles() (files map[string]io.Reader) {
	return
}
