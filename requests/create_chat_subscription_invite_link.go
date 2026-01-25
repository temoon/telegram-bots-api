package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type CreateChatSubscriptionInviteLink struct {
	ChatId             telegram.ChatId
	SubscriptionPeriod int64
	SubscriptionPrice  int64
	Name               *string
}

func (r *CreateChatSubscriptionInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatInviteLink)
	err = b.CallMethod(ctx, "createChatSubscriptionInviteLink", r, response)
	return
}

func (r *CreateChatSubscriptionInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["subscription_period"] = strconv.FormatInt(r.SubscriptionPeriod, 10)

	values["subscription_price"] = strconv.FormatInt(r.SubscriptionPrice, 10)

	if r.Name != nil {
		values["name"] = *r.Name
	}

	return
}

func (r *CreateChatSubscriptionInviteLink) GetFiles() (files map[string]io.Reader) {
	return
}
