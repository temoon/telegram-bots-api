package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type SetChatDescription struct {
	ChatId      interface{}
	Description *string
}

func (r *SetChatDescription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setChatDescription", r, response)
	return
}

func (r *SetChatDescription) IsMultipart() (multipart bool) {
	return false
}

func (r *SetChatDescription) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.Description != nil {
		values["description"] = *r.Description
	}

	return
}
