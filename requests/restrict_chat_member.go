package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type RestrictChatMember struct {
	UserId                        int64
	Permissions                   telegram.ChatPermissions
	UseIndependentChatPermissions *bool
	UntilDate                     *int64
	ChatId                        telegram.ChatId
}

func (r *RestrictChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "restrictChatMember", r, response)
	return
}

func (r *RestrictChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	var dataPermissions []byte
	if dataPermissions, err = json.Marshal(r.Permissions); err != nil {
		return
	}

	values["permissions"] = string(dataPermissions)

	if r.UseIndependentChatPermissions != nil {
		if *r.UseIndependentChatPermissions {
			values["use_independent_chat_permissions"] = "1"
		} else {
			values["use_independent_chat_permissions"] = "0"
		}
	}

	if r.UntilDate != nil {
		values["until_date"] = strconv.FormatInt(*r.UntilDate, 10)
	}

	values["chat_id"] = r.ChatId.String()

	return
}

func (r *RestrictChatMember) GetFiles() (files map[string]io.Reader) {
	return
}
