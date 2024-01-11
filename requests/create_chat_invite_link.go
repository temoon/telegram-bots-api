package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type CreateChatInviteLink struct {
	ChatId             interface{}
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

func (r *CreateChatInviteLink) IsMultipart() bool {
	return false
}

func (r *CreateChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

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
