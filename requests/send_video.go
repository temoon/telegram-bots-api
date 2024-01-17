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
	Thumbnail           *telegram.InputFile
	CaptionEntities     []telegram.MessageEntity
	HasSpoiler          *bool
	Width               *int64
	Height              *int64
	Duration            *int64
	SupportsStreaming   *bool
	ProtectContent      *bool
	ChatId              telegram.ChatId
	MessageThreadId     *int64
	ReplyMarkup         interface{}
	ParseMode           *string
	DisableNotification *bool
	ReplyParameters     *telegram.ReplyParameters
	Video               telegram.InputFile
	Caption             *string
}

func (r *SendVideo) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendVideo", r, response)
	return
}

func (r *SendVideo) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Thumbnail != nil {
		values["thumbnail"] = r.Thumbnail.GetValue()
	}

	if r.CaptionEntities != nil {
		var dataCaptionEntities []byte
		if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(dataCaptionEntities)
	}

	if r.HasSpoiler != nil {
		if *r.HasSpoiler {
			values["has_spoiler"] = "1"
		} else {
			values["has_spoiler"] = "0"
		}
	}

	if r.Width != nil {
		values["width"] = strconv.FormatInt(*r.Width, 10)
	}

	if r.Height != nil {
		values["height"] = strconv.FormatInt(*r.Height, 10)
	}

	if r.Duration != nil {
		values["duration"] = strconv.FormatInt(*r.Duration, 10)
	}

	if r.SupportsStreaming != nil {
		if *r.SupportsStreaming {
			values["supports_streaming"] = "1"
		} else {
			values["supports_streaming"] = "0"
		}
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	values["chat_id"] = r.ChatId.String()

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
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

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	values["video"] = r.Video.GetValue()

	if r.Caption != nil {
		values["caption"] = *r.Caption
	}

	return
}

func (r *SendVideo) GetFiles() (files map[string]io.Reader) {
	return
}
