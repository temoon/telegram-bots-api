package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetChatMemberTag struct {
	ChatId telegram.ChatId
	UserId int64
	Tag    *string
}

func (r *SetChatMemberTag) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatMemberTag", r, response)
	return
}

func (r *SetChatMemberTag) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.Tag != nil {
		values["tag"] = *r.Tag
	}

	return
}

func (r *SetChatMemberTag) GetFiles() (files map[string]io.Reader) {
	return
}
