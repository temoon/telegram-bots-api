package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type BanChatMember struct {
	UntilDate      *int64
	RevokeMessages *bool
	ChatId         telegram.ChatId
	UserId         int64
}

func (r *BanChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "banChatMember", r, response)
	return
}

func (r *BanChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.UntilDate != nil {
		values["until_date"] = strconv.FormatInt(*r.UntilDate, 10)
	}

	if r.RevokeMessages != nil {
		if *r.RevokeMessages {
			values["revoke_messages"] = "1"
		} else {
			values["revoke_messages"] = "0"
		}
	}

	values["chat_id"] = r.ChatId.String()

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *BanChatMember) GetFiles() (files map[string]io.Reader) {
	return
}
