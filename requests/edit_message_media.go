package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type EditMessageMedia struct {
	ChatId          interface{}
	InlineMessageId *string
	Media           interface{}
	MessageId       *int64
	ReplyMarkup     *telegram.InlineKeyboardMarkup
}

func (r *EditMessageMedia) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageMedia", r, response)
	return
}

func (r *EditMessageMedia) IsMultipart() bool {
	return false
}

func (r *EditMessageMedia) GetValues() (values map[string]interface{}, err error) {
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

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	var dataMedia []byte
	if dataMedia, err = json.Marshal(r.Media); err != nil {
		return
	}

	values["media"] = string(dataMedia)

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
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
