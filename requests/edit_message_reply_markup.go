package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditMessageReplyMarkup struct {
	ChatId          *telegram.ChatId
	MessageId       *int64
	InlineMessageId *string
	ReplyMarkup     *telegram.InlineKeyboardMarkup
}

func (r *EditMessageReplyMarkup) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageReplyMarkup", r, response)
	return
}

func (r *EditMessageReplyMarkup) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ChatId != nil {
		values["chat_id"] = r.ChatId.String()
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
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

func (r *EditMessageReplyMarkup) GetFiles() (files map[string]io.Reader) {
	return
}
