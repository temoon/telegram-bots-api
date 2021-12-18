package requests

import (
	"context"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type EditChatInviteLink struct {
	ChatId             interface{}
	CreatesJoinRequest bool
	ExpireDate         int32
	InviteLink         string
	MemberLimit        int32
	Name               string
}

func (r *EditChatInviteLink) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(string)
	err = b.CallMethod(ctx, "editChatInviteLink", r, response)
	return
}

func (r *EditChatInviteLink) IsMultipart() (multipart bool) {
	return false
}

func (r *EditChatInviteLink) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.CreatesJoinRequest {
		values["creates_join_request"] = "1"
	}

	if r.ExpireDate != 0 {
		values["expire_date"] = strconv.FormatInt(int64(r.ExpireDate), 10)
	}

	values["invite_link"] = r.InviteLink

	if r.MemberLimit != 0 {
		values["member_limit"] = strconv.FormatInt(int64(r.MemberLimit), 10)
	}

	if r.Name != "" {
		values["name"] = r.Name
	}

	return
}
