package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type SendMessage struct {
	ChatId              interface{}
	DisableNotification *bool
	Entities            []telegram.MessageEntity
	LinkPreviewOptions  *telegram.LinkPreviewOptions
	MessageThreadId     *int64
	ParseMode           *string
	ProtectContent      *bool
	ReplyMarkup         interface{}
	ReplyParameters     *telegram.ReplyParameters
	Text                string
}

func (r *SendMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendMessage", r, response)
	return
}

func (r *SendMessage) IsMultipart() bool {
	return false
}

func (r *SendMessage) GetValues() (values map[string]interface{}, err error) {
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

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.Entities != nil {
		var dataEntities []byte
		if dataEntities, err = json.Marshal(r.Entities); err != nil {
			return
		}

		values["entities"] = string(dataEntities)
	}

	if r.LinkPreviewOptions != nil {
		var dataLinkPreviewOptions []byte
		if dataLinkPreviewOptions, err = json.Marshal(r.LinkPreviewOptions); err != nil {
			return
		}

		values["link_preview_options"] = string(dataLinkPreviewOptions)
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.ReplyMarkup != nil {
		switch value := r.ReplyMarkup.(type) {
		case telegram.InlineKeyboardMarkup, telegram.ReplyKeyboardMarkup, telegram.ReplyKeyboardRemove, telegram.ForceReply:
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		default:
			err = errors.New("invalid reply_markup field type")
			return
		}
	}

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	values["text"] = r.Text

	return
}
