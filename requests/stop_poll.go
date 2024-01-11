package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type StopPoll struct {
	ChatId      interface{}
	MessageId   int64
	ReplyMarkup *telegram.InlineKeyboardMarkup
}

func (r *StopPoll) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Poll)
	err = b.CallMethod(ctx, "stopPoll", r, response)
	return
}

func (r *StopPoll) IsMultipart() bool {
	return false
}

func (r *StopPoll) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	return
}
