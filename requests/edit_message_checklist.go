package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type EditMessageChecklist struct {
	BusinessConnectionId string
	ChatId               int64
	Checklist            telegram.InputChecklist
	MessageId            int64
	ReplyMarkup          *telegram.InlineKeyboardMarkup
}

func (r *EditMessageChecklist) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageChecklist", r, response)
	return
}

func (r *EditMessageChecklist) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["business_connection_id"] = r.BusinessConnectionId

	values["chat_id"] = strconv.FormatInt(r.ChatId, 10)

	var dataChecklist []byte
	if dataChecklist, err = json.Marshal(r.Checklist); err != nil {
		return
	}

	values["checklist"] = string(dataChecklist)

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

func (r *EditMessageChecklist) GetFiles() (files map[string]io.Reader) {
	return
}
