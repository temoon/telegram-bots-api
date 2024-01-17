package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
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

func (r *ApproveChatJoinRequest) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *ApproveChatJoinRequest) GetFiles() (files map[string]io.Reader) {
	return
}
