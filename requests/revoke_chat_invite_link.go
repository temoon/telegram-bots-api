package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type RevokeChatInviteLink struct {
	ChatId     telegram.ChatId
	InviteLink string
}

func (r *RevokeChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatInviteLink)
	err = b.CallMethod(ctx, "revokeChatInviteLink", r, response)
	return
}

func (r *RevokeChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["invite_link"] = r.InviteLink

	return
}

func (r *RevokeChatInviteLink) GetFiles() (files map[string]io.Reader) {
	return
}
