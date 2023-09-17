package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type RevokeChatInviteLink struct {
	ChatId     interface{}
	InviteLink string
}

func (r *RevokeChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "revokeChatInviteLink", r, response)
	return
}

func (r *RevokeChatInviteLink) IsMultipart() (multipart bool) {
	return false
}

func (r *RevokeChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["invite_link"] = r.InviteLink

	return
}
