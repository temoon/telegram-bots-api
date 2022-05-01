package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetChatPermissions struct {
	ChatId      interface{}
	Permissions telegram.ChatPermissions
}

func (r *SetChatPermissions) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatPermissions", r, response)
	return
}

func (r *SetChatPermissions) IsMultipart() (multipart bool) {
	return false
}

func (r *SetChatPermissions) GetValues() (values map[string]interface{}, err error) {
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

	return
}
