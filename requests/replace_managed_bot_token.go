package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type ReplaceManagedBotToken struct {
	UserId int64
}

func (r *ReplaceManagedBotToken) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "replaceManagedBotToken", r, response)
	return
}

func (r *ReplaceManagedBotToken) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *ReplaceManagedBotToken) GetFiles() (files map[string]io.Reader) {
	return
}
