package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendAnimation struct {
	MessageThreadId     *int64
	Duration            *int64
	Caption             *string
	ParseMode           *string
	CaptionEntities     []telegram.MessageEntity
	Width               *int64
	HasSpoiler          *bool
	ProtectContent      *bool
	ReplyParameters     *telegram.ReplyParameters
	ReplyMarkup         interface{}
	Animation           telegram.InputFile
	Height              *int64
	DisableNotification *bool
	ChatId              telegram.ChatId
	Thumbnail           *telegram.InputFile
}

func (r *SendAnimation) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendAnimation", r, response)
	return
}

func (r *SendAnimation) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.Duration != nil {
		values["duration"] = strconv.FormatInt(*r.Duration, 10)
	}

	if r.Caption != nil {
		values["caption"] = *r.Caption
	}

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.CaptionEntities != nil {
		var dataCaptionEntities []byte
		if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(dataCaptionEntities)
	}

	if r.Width != nil {
		values["width"] = strconv.FormatInt(*r.Width, 10)
	}

	if r.HasSpoiler != nil {
		if *r.HasSpoiler {
			values["has_spoiler"] = "1"
		} else {
			values["has_spoiler"] = "0"
		}
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

	values["animation"] = r.Animation.GetValue()

	if r.Height != nil {
		values["height"] = strconv.FormatInt(*r.Height, 10)
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	values["chat_id"] = r.ChatId.String()

	if r.Thumbnail != nil {
		values["thumbnail"] = r.Thumbnail.GetValue()
	}

	return
}

func (r *SendAnimation) GetFiles() (files map[string]io.Reader) {
	return
}
