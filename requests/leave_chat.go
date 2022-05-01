package requests

import (
	"context"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type LeaveChat struct {
	ChatId interface{}
}

func (r *LeaveChat) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "leaveChat", r, response)
	return
}

func (r *LeaveChat) IsMultipart() (multipart bool) {
	return false
}

func (r *LeaveChat) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
