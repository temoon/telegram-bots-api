package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type DeclineChatJoinRequest struct {
	ChatId interface{}
	UserId int64
}

func (r *DeclineChatJoinRequest) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "declineChatJoinRequest", r, response)
	return
}

func (r *DeclineChatJoinRequest) IsMultipart() (multipart bool) {
	return false
}

func (r *DeclineChatJoinRequest) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
