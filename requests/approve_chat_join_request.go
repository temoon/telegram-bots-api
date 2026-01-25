package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type ApproveChatJoinRequest struct {
	ChatId telegram.ChatId
	UserId int64
}

func (r *ApproveChatJoinRequest) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "approveChatJoinRequest", r, response)
	return
}

func (r *ApproveChatJoinRequest) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *ApproveChatJoinRequest) GetFiles() (files map[string]io.Reader) {
	return
}
