package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type UnbanChatMember struct {
	ChatId       telegram.ChatId
	UserId       int64
	OnlyIfBanned *bool
}

func (r *UnbanChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unbanChatMember", r, response)
	return
}

func (r *UnbanChatMember) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.OnlyIfBanned != nil {
		if *r.OnlyIfBanned {
			values["only_if_banned"] = "1"
		} else {
			values["only_if_banned"] = "0"
		}
	}

	return
}

func (r *UnbanChatMember) GetFiles() (files map[string]io.Reader) {
	return
}
