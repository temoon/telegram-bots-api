package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type RestrictChatMember struct {
	ChatId      interface{}
	Permissions telegram.ChatPermissions
	UntilDate   *int32
	UserId      int64
}

func (r *RestrictChatMember) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "restrictChatMember", r, response)
	return
}

func (r *RestrictChatMember) IsMultipart() (multipart bool) {
	return false
}

func (r *RestrictChatMember) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	var dataPermissions []byte
	if dataPermissions, err = json.Marshal(r.Permissions); err != nil {
		return
	}

	values["permissions"] = string(dataPermissions)

	if r.UntilDate != nil {
		values["until_date"] = strconv.FormatInt(int64(*r.UntilDate), 10)
	}

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	return
}
