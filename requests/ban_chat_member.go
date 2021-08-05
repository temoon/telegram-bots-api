package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type BanChatMember struct {
	ChatId         interface{}
	RevokeMessages bool
	UntilDate      int32
	UserId         int64
}

func (r *BanChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "banChatMember", r, response)
	return
}

func (r *BanChatMember) IsMultipart() (multipart bool) {
	return false
}

func (r *BanChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.RevokeMessages {
		values["revoke_messages"] = "1"
	}

	if r.UntilDate != 0 {
		values["until_date"] = strconv.FormatInt(int64(r.UntilDate), 10)
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
