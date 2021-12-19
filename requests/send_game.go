package requests

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/temoon/go-telegram-bots-api"
)

type SendGame struct {
	AllowSendingWithoutReply *bool
	ChatId                   int64
	DisableNotification      *bool
	GameShortName            string
	ReplyMarkup              *telegram.InlineKeyboardMarkup
	ReplyToMessageId         *int32
}

func (r *SendGame) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendGame", r, response)
	return
}

func (r *SendGame) IsMultipart() (multipart bool) {
	return false
}

func (r *SendGame) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply != nil {
		if *r.AllowSendingWithoutReply {
			values["allow_sending_without_reply"] = "1"
		} else {
			values["allow_sending_without_reply"] = "0"
		}
	}

	values["chat_id"] = strconv.FormatInt(r.ChatId, 10)

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	values["game_short_name"] = r.GameShortName

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	if r.ReplyToMessageId != nil {
		values["reply_to_message_id"] = strconv.FormatInt(int64(*r.ReplyToMessageId), 10)
	}

	return
}
