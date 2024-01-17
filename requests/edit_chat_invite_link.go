package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditChatInviteLink struct {
	ChatId             telegram.ChatId
	InviteLink         string
	CreatesJoinRequest *bool
	ExpireDate         *int64
	MemberLimit        *int64
	Name               *string
}

func (r *EditChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ChatInviteLink)
	err = b.CallMethod(ctx, "editChatInviteLink", r, response)
	return
}

func (r *EditChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["invite_link"] = r.InviteLink

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

func (r *EditChatInviteLink) GetFiles() (files map[string]io.Reader) {
	return
}
