package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type EditMessageText struct {
	ChatId             interface{}
	Entities           []telegram.MessageEntity
	InlineMessageId    *string
	LinkPreviewOptions *telegram.LinkPreviewOptions
	MessageId          *int64
	ParseMode          *string
	ReplyMarkup        *telegram.InlineKeyboardMarkup
	Text               string
}

func (r *EditMessageText) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageText", r, response)
	return
}

func (r *EditMessageText) IsMultipart() bool {
	return false
}

func (r *EditMessageText) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

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

	if r.LinkPreviewOptions != nil {
		var dataLinkPreviewOptions []byte
		if dataLinkPreviewOptions, err = json.Marshal(r.LinkPreviewOptions); err != nil {
			return
		}

		values["link_preview_options"] = string(dataLinkPreviewOptions)
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

	values["text"] = r.Text

	return
}
