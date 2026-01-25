package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type LeaveChat struct {
	ChatId telegram.ChatId
}

func (r *LeaveChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "leaveChat", r, response)
	return
}

func (r *LeaveChat) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *LeaveChat) GetFiles() (files map[string]io.Reader) {
	return
}
