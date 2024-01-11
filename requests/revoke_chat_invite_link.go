package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type RevokeChatInviteLink struct {
	ChatId     interface{}
	InviteLink string
}

func (r *RevokeChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatInviteLink)
	err = b.CallMethod(ctx, "revokeChatInviteLink", r, response)
	return
}

func (r *RevokeChatInviteLink) IsMultipart() bool {
	return false
}

func (r *RevokeChatInviteLink) GetValues() (values map[string]interface{}, err error) {
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

	values["invite_link"] = r.InviteLink

	return
}
