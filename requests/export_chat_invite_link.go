package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
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

func (r *ExportChatInviteLink) IsMultipart() bool {
	return false
}

func (r *ExportChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

	return
}
