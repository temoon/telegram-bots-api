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
	Title               *string
	ReplyMarkup         interface{}
	ReplyParameters     *telegram.ReplyParameters
	ChatId              telegram.ChatId
	MessageThreadId     *int64
	CaptionEntities     []telegram.MessageEntity
	Performer           *string
	Audio               telegram.InputFile
	ParseMode           *string
	Duration            *int64
	Thumbnail           *telegram.InputFile
	Caption             *string
	DisableNotification *bool
	ProtectContent      *bool
}

func (r *SendAudio) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendAudio", r, response)
	return
}

func (r *SendAudio) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.Title != nil {
		values["title"] = *r.Title
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

	if r.ReplyParameters != nil {
		var dataReplyParameters []byte
		if dataReplyParameters, err = json.Marshal(r.ReplyParameters); err != nil {
			return
		}

		values["reply_parameters"] = string(dataReplyParameters)
	}

	values["chat_id"] = r.ChatId.String()

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.CaptionEntities != nil {
		var dataCaptionEntities []byte
		if dataCaptionEntities, err = json.Marshal(r.CaptionEntities); err != nil {
			return
		}

		values["caption_entities"] = string(dataCaptionEntities)
	}

	if r.Performer != nil {
		values["performer"] = *r.Performer
	}

	values["audio"] = r.Audio.GetValue()

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.Duration != nil {
		values["duration"] = strconv.FormatInt(*r.Duration, 10)
	}

	if r.Thumbnail != nil {
		values["thumbnail"] = r.Thumbnail.GetValue()
	}

	if r.Caption != nil {
		values["caption"] = *r.Caption
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	return
}

func (r *SendAudio) GetFiles() (files map[string]io.Reader) {
	return
}
