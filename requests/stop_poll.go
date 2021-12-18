package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type StopPoll struct {
	ChatId      interface{}
	MessageId   int32
	ReplyMarkup *telegram.InlineKeyboardMarkup
}

func (r *StopPoll) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Poll)
	err = b.CallMethod(ctx, "stopPoll", r, response)
	return
}

func (r *StopPoll) IsMultipart() (multipart bool) {
	return false
}

func (r *StopPoll) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	return
}
