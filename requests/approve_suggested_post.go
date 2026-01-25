package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type ApproveSuggestedPost struct {
	ChatId    int64
	MessageId int64
	SendDate  *int64
}

func (r *ApproveSuggestedPost) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "approveSuggestedPost", r, response)
	return
}

func (r *ApproveSuggestedPost) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = strconv.FormatInt(r.ChatId, 10)

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.SendDate != nil {
		values["send_date"] = strconv.FormatInt(*r.SendDate, 10)
	}

	return
}

func (r *ApproveSuggestedPost) GetFiles() (files map[string]io.Reader) {
	return
}
