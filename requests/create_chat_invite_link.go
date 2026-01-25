package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type CreateChatInviteLink struct {
	ChatId             telegram.ChatId
	CreatesJoinRequest *bool
	ExpireDate         *int64
	MemberLimit        *int64
	Name               *string
}

func (r *CreateChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatInviteLink)
	err = b.CallMethod(ctx, "createChatInviteLink", r, response)
	return
}

func (r *CreateChatInviteLink) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	if r.CreatesJoinRequest != nil {
		if *r.CreatesJoinRequest {
			values["creates_join_request"] = "1"
		} else {
			values["creates_join_request"] = "0"
		}
	}

	if r.ExpireDate != nil {
		values["expire_date"] = strconv.FormatInt(*r.ExpireDate, 10)
	}

	if r.MemberLimit != nil {
		values["member_limit"] = strconv.FormatInt(*r.MemberLimit, 10)
	}

	if r.Name != nil {
		values["name"] = *r.Name
	}

	return
}

func (r *CreateChatInviteLink) GetFiles() (files map[string]io.Reader) {
	return
}
