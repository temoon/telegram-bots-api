package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type StopPoll struct {
	ChatId      telegram.ChatId
	MessageId   int64
	ReplyMarkup *telegram.InlineKeyboardMarkup
}

func (r *StopPoll) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Poll)
	err = b.CallMethod(ctx, "stopPoll", r, response)
	return
}

func (r *StopPoll) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

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

func (r *StopPoll) GetFiles() (files map[string]io.Reader) {
	return
}
