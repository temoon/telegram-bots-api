package requests

import (
	"context"
	"io"

	"github.com/temoon/telegram-bots-api"
)

type EditChatSubscriptionInviteLink struct {
	ChatId     telegram.ChatId
	InviteLink string
	Name       *string
}

func (r *EditChatSubscriptionInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatInviteLink)
	err = b.CallMethod(ctx, "editChatSubscriptionInviteLink", r, response)
	return
}

func (r *EditChatSubscriptionInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["invite_link"] = r.InviteLink

	if r.Name != nil {
		values["name"] = *r.Name
	}

	return
}

func (r *EditChatSubscriptionInviteLink) GetFiles() (files map[string]io.Reader) {
	return
}
