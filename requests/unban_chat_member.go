package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type UnbanChatMember struct {
	ChatId       interface{}
	OnlyIfBanned bool
	UserId       int64
}

func (r *UnbanChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "unbanChatMember", r, response)
	return
}

func (r *UnbanChatMember) IsMultipart() (multipart bool) {
	return false
}

func (r *UnbanChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.OnlyIfBanned {
		values["only_if_banned"] = "1"
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
