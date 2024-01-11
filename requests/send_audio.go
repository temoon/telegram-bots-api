package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendAudio struct {
	Audio               interface{}
	Caption             *string
	CaptionEntities     []telegram.MessageEntity
	ChatId              interface{}
	DisableNotification *bool
	Duration            *int64
	MessageThreadId     *int64
	ParseMode           *string
	Performer           *string
	ProtectContent      *bool
	ReplyMarkup         interface{}
	ReplyParameters     *telegram.ReplyParameters
	Thumbnail           interface{}
	Title               *string
}

func (r *SendAudio) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendAudio", r, response)
	return
}

func (r *SendAudio) IsMultipart() bool {
	return true
}

func (r *SendAudio) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.Audio.(type) {
	case io.Reader:
		values["audio"] = value
	case string:
		values["audio"] = value
	default:
		err = errors.New("invalid audio field type")
		return
	}

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

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.Performer != nil {
		values["performer"] = *r.Performer
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

	if r.Title != nil {
		values["title"] = *r.Title
	}

	return
}
