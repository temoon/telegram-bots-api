package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type CreateChatInviteLink struct {
	ChatId             telegram.ChatId
	Name               *string
	ExpireDate         *int64
	MemberLimit        *int64
	CreatesJoinRequest *bool
}

func (r *CreateChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatInviteLink)
	err = b.CallMethod(ctx, "createChatInviteLink", r, response)
	return
}

func (r *CreateChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	if r.Name != nil {
		values["name"] = *r.Name
	}

	if r.ExpireDate != nil {
		values["expire_date"] = strconv.FormatInt(*r.ExpireDate, 10)
	}

	if r.MemberLimit != nil {
		values["member_limit"] = strconv.FormatInt(*r.MemberLimit, 10)
	}

	if r.CreatesJoinRequest != nil {
		if *r.CreatesJoinRequest {
			values["creates_join_request"] = "1"
		} else {
			values["creates_join_request"] = "0"
		}
	}

	return
}

func (r *CreateChatInviteLink) GetFiles() (files map[string]io.Reader) {
	return
}
