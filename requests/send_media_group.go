package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SendMediaGroup struct {
	AllowSendingWithoutReply *bool
	ChatId                   interface{}
	DisableNotification      *bool
	Media                    []interface{}
	ReplyToMessageId         *int32
}

func (r *SendMediaGroup) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Message)
	err = b.CallMethod(ctx, "sendMediaGroup", r, response)
	return
}

func (r *SendMediaGroup) IsMultipart() (multipart bool) {
	return true
}

func (r *SendMediaGroup) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply != nil {
		if *r.AllowSendingWithoutReply {
			values["allow_sending_without_reply"] = "1"
		} else {
			values["allow_sending_without_reply"] = "0"
		}
	}

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	var dataMedia []byte
	if dataMedia, err = json.Marshal(r.Media); err != nil {
		return
	}

	values["media"] = string(dataMedia)

	if r.ReplyToMessageId != nil {
		values["reply_to_message_id"] = strconv.FormatInt(int64(*r.ReplyToMessageId), 10)
	}

	return
}
