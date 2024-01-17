package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type GetChatMemberCount struct {
	ChatId telegram.ChatId
}

func (r *GetChatMemberCount) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(int64)
	err = b.CallMethod(ctx, "getChatMemberCount", r, response)
	return
}

func (r *GetChatMemberCount) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *GetChatMemberCount) GetFiles() (files map[string]io.Reader) {
	return
}
