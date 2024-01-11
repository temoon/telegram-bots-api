package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type DeleteMessages struct {
	ChatId     interface{}
	MessageIds []int64
}

func (r *DeleteMessages) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteMessages", r, response)
	return
}

func (r *DeleteMessages) IsMultipart() bool {
	return false
}

func (r *DeleteMessages) GetValues() (values map[string]interface{}, err error) {
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

	var dataMessageIds []byte
	if dataMessageIds, err = json.Marshal(r.MessageIds); err != nil {
		return
	}

	values["message_ids"] = string(dataMessageIds)

	return
}
