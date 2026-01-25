package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendChecklist struct {
	BusinessConnectionId string
	ChatId               int64
	Checklist            telegram.InputChecklist
	DisableNotification  *bool
	MessageEffectId      *string
	ProtectContent       *bool
	ReplyMarkup          *telegram.InlineKeyboardMarkup
	ReplyParameters      *telegram.ReplyParameters
}

func (r *SendChecklist) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendChecklist", r, response)
	return
}

func (r *SendChecklist) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["business_connection_id"] = r.BusinessConnectionId

	values["chat_id"] = strconv.FormatInt(r.ChatId, 10)

	var dataChecklist []byte
	if dataChecklist, err = json.Marshal(r.Checklist); err != nil {
		return
	}

	values["checklist"] = string(dataChecklist)

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.MessageEffectId != nil {
		values["message_effect_id"] = *r.MessageEffectId
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	return
}

func (r *SendChecklist) GetFiles() (files map[string]io.Reader) {
	return
}
