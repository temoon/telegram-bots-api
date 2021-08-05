package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type ExportChatInviteLink struct {
	ChatId interface{}
}

func (r *ExportChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "exportChatInviteLink", r, response)
	return
}

func (r *ExportChatInviteLink) IsMultipart() (multipart bool) {
	return false
}

func (r *ExportChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	return
}
