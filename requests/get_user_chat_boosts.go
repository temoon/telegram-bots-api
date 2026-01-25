package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GetUserChatBoosts struct {
	ChatId telegram.ChatId
	UserId int64
}

func (r *GetUserChatBoosts) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.UserChatBoosts)
	err = b.CallMethod(ctx, "getUserChatBoosts", r, response)
	return
}

func (r *GetUserChatBoosts) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *GetUserChatBoosts) GetFiles() (files map[string]io.Reader) {
	return
}
