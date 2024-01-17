package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type SetChatPermissions struct {
	ChatId                        telegram.ChatId
	Permissions                   telegram.ChatPermissions
	UseIndependentChatPermissions *bool
}

func (r *SetChatPermissions) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatPermissions", r, response)
	return
}

func (r *SetChatPermissions) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

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

	return
}

func (r *SetChatPermissions) GetFiles() (files map[string]io.Reader) {
	return
}
