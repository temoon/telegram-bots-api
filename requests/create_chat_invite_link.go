package requests

import (
	"context"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type CreateChatInviteLink struct {
	ChatId      interface{}
	ExpireDate  int32
	MemberLimit int32
}

func (r *CreateChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "createChatInviteLink", r, response)
	return
}

func (r *CreateChatInviteLink) IsMultipart() (multipart bool) {
	return false
}

func (r *CreateChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.ExpireDate != 0 {
		values["expire_date"] = strconv.FormatInt(int64(r.ExpireDate), 10)
	}

	if r.MemberLimit != 0 {
		values["member_limit"] = strconv.FormatInt(int64(r.MemberLimit), 10)
	}

	return
}
