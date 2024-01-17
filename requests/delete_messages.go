package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
)

type DeleteMessages struct {
	ChatId     telegram.ChatId
	MessageIds []int64
}

func (r *DeleteMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteMessages", r, response)
	return
}

func (r *DeleteMessages) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	var dataMessageIds []byte
	if dataMessageIds, err = json.Marshal(r.MessageIds); err != nil {
		return
	}

	values["message_ids"] = string(dataMessageIds)

	return
}

func (r *DeleteMessages) GetFiles() (files map[string]io.Reader) {
	return
}
