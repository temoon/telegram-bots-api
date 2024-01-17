package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type CopyMessage struct {
	ReplyMarkup         interface{}
	ChatId              telegram.ChatId
	MessageThreadId     *int64
	ProtectContent      *bool
	ReplyParameters     *telegram.ReplyParameters
	CaptionEntities     []telegram.MessageEntity
	DisableNotification *bool
	FromChatId          telegram.ChatId
	MessageId           int64
	Caption             *string
	ParseMode           *string
}

func (r *CopyMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.MessageId)
	err = b.CallMethod(ctx, "copyMessage", r, response)
	return
}

func (r *CopyMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ReplyMarkup != nil {
		switch value := r.ReplyMarkup.(type) {
		case telegram.InlineKeyboardMarkup, telegram.ReplyKeyboardMarkup, telegram.ReplyKeyboardRemove, telegram.ForceReply:
			var data []byte
			if data, err = json.Marshal(value); err != nil {
				return
			}

			values["reply_markup"] = string(data)
		default:
			err = errors.New("unsupported reply_markup field type")
			return
		}
	}

	values["chat_id"] = r.ChatId.String()

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	if r.CaptionEntities != nil {
		var dataCaptionEntities []byte
		if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(dataCaptionEntities)
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	values["from_chat_id"] = r.FromChatId.String()

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.Caption != nil {
		values["caption"] = *r.Caption
	}

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	return
}

func (r *CopyMessage) GetFiles() (files map[string]io.Reader) {
	return
}
