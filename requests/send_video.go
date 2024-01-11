package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendVideo struct {
	Caption             *string
	CaptionEntities     []telegram.MessageEntity
	ChatId              interface{}
	DisableNotification *bool
	Duration            *int64
	HasSpoiler          *bool
	Height              *int64
	MessageThreadId     *int64
	ParseMode           *string
	ProtectContent      *bool
	ReplyMarkup         interface{}
	ReplyParameters     *telegram.ReplyParameters
	SupportsStreaming   *bool
	Thumbnail           interface{}
	Video               interface{}
	Width               *int64
}

func (r *SendVideo) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendVideo", r, response)
	return
}

func (r *SendVideo) IsMultipart() bool {
	return true
}

func (r *SendVideo) GetValues() (values map[string]interface{}, err error) {
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

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.Duration != nil {
		values["duration"] = strconv.FormatInt(*r.Duration, 10)
	}

	if r.HasSpoiler != nil {
		if *r.HasSpoiler {
			values["has_spoiler"] = "1"
		} else {
			values["has_spoiler"] = "0"
		}
	}

	if r.Height != nil {
		values["height"] = strconv.FormatInt(*r.Height, 10)
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

	if r.SupportsStreaming != nil {
		if *r.SupportsStreaming {
			values["supports_streaming"] = "1"
		} else {
			values["supports_streaming"] = "0"
		}
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

	switch value := r.Video.(type) {
	case io.Reader:
		values["video"] = value
	case string:
		values["video"] = value
	default:
		err = errors.New("invalid video field type")
		return
	}

	if r.Width != nil {
		values["width"] = strconv.FormatInt(*r.Width, 10)
	}

	return
}
