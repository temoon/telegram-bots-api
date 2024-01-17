package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type DeclineChatJoinRequest struct {
	ChatId telegram.ChatId
	UserId int64
}

func (r *DeclineChatJoinRequest) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "declineChatJoinRequest", r, response)
	return
}

func (r *DeclineChatJoinRequest) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *DeclineChatJoinRequest) GetFiles() (files map[string]io.Reader) {
	return
}
