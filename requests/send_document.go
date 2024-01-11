package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendDocument struct {
	Caption                     *string
	CaptionEntities             []telegram.MessageEntity
	ChatId                      interface{}
	DisableContentTypeDetection *bool
	DisableNotification         *bool
	Document                    interface{}
	MessageThreadId             *int64
	ParseMode                   *string
	ProtectContent              *bool
	ReplyMarkup                 interface{}
	ReplyParameters             *telegram.ReplyParameters
	Thumbnail                   interface{}
}

func (r *SendDocument) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendDocument", r, response)
	return
}

func (r *SendDocument) IsMultipart() bool {
	return true
}

func (r *SendDocument) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Caption != nil {
		values["caption"] = *r.Caption
	}

	if r.CaptionEntities != nil {
		var dataCaptionEntities []byte
		if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(dataCaptionEntities)
	}

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

	if r.DisableContentTypeDetection != nil {
		if *r.DisableContentTypeDetection {
			values["disable_content_type_detection"] = "1"
		} else {
			values["disable_content_type_detection"] = "0"
		}
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	switch value := r.Document.(type) {
	case io.Reader:
		values["document"] = value
	case string:
		values["document"] = value
	default:
		err = errors.New("invalid document field type")
		return
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

	if r.Thumbnail != nil {
		switch value := r.Thumbnail.(type) {
		case io.Reader:
			values["thumbnail"] = value
		case string:
			values["thumbnail"] = value
		default:
			err = errors.New("invalid thumbnail field type")
			return
		}
	}

	return
}
