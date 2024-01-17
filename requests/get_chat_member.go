package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type GetChatMember struct {
	ChatId telegram.ChatId
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
		err = errors.New("unsupported response type")
	}

	return
}

func (r *GetChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}

func (r *GetChatMember) GetFiles() (files map[string]io.Reader) {
	return
}
