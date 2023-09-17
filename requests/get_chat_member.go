package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type GetChatMember struct {
	ChatId interface{}
	UserId int64
}

func (r *GetChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(interface{})
	err = b.CallMethod(ctx, "getChatMember", r, response)
	return
}

func (r *GetChatMember) CallWithResponse(ctx context.Context, b *telegram.Bot, response interface{}) (err error) {
	switch response.(type) {
	case *telegram.ChatMemberOwner, *telegram.ChatMemberAdministrator, *telegram.ChatMemberMember, *telegram.ChatMemberRestricted, *telegram.ChatMemberLeft, *telegram.ChatMemberBanned:
		err = b.CallMethod(ctx, "getChatMember", r, response)
	default:
		err = errors.New("unexpected response type")
	}

	return
}

func (r *GetChatMember) IsMultipart() (multipart bool) {
	return false
}

func (r *GetChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
