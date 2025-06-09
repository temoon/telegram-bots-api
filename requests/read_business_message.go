package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type ReadBusinessMessage struct {
	BusinessConnectionId string
	ChatId               int64
	MessageId            int64
}

func (r *ReadBusinessMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "readBusinessMessage", r, response)
	return
}

func (r *ReadBusinessMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	values["chat_id"] = strconv.FormatInt(r.ChatId, 10)

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	return
}

func (r *ReadBusinessMessage) GetFiles() (files map[string]io.Reader) {
	return
}
