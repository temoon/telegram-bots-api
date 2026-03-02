package requests

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendAudio struct {
	Audio                   telegram.InputFile
	ChatId                  telegram.ChatId
	AllowPaidBroadcast      *bool
	BusinessConnectionId    *string
	Caption                 *string
	CaptionEntities         []telegram.MessageEntity
	DirectMessagesTopicId   *int64
	DisableNotification     *bool
	Duration                *int64
	MessageEffectId         *string
	MessageThreadId         *int64
	ParseMode               *string
	Performer               *string
	ProtectContent          *bool
	ReplyMarkup             interface{}
	ReplyParameters         *telegram.ReplyParameters
	SuggestedPostParameters *telegram.SuggestedPostParameters
	Thumbnail               *telegram.InputFile
	Title                   *string
}

func (r *SendAudio) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendAudio", r, response)
	return
}

func (r *SendAudio) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["audio"] = r.Audio.String()

	values["chat_id"] = r.ChatId.String()

	if r.AllowPaidBroadcast != nil {
		if *r.AllowPaidBroadcast {
			values["allow_paid_broadcast"] = "1"
		} else {
			values["allow_paid_broadcast"] = "0"
		}
	}

	if r.BusinessConnectionId != nil {
		values["business_connection_id"] = *r.BusinessConnectionId
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

	if r.DirectMessagesTopicId != nil {
		values["direct_messages_topic_id"] = strconv.FormatInt(*r.DirectMessagesTopicId, 10)
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

	if r.MessageEffectId != nil {
		values["message_effect_id"] = *r.MessageEffectId
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

	if r.SuggestedPostParameters != nil {
		var dataSuggestedPostParameters []byte
		if dataSuggestedPostParameters, err = json.Marshal(r.SuggestedPostParameters); err != nil {
			return
		}

		values["suggested_post_parameters"] = string(dataSuggestedPostParameters)
	}

	if r.Thumbnail != nil {
		values["thumbnail"] = r.Thumbnail.String()
	}

	if r.Title != nil {
		values["title"] = *r.Title
	}

	return
}

func (r *SendAudio) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	if r.Thumbnail != nil && r.Thumbnail.HasFile() {
		files[r.Thumbnail.GetFormFieldName()] = r.Thumbnail.GetFile()
	}
	if r.Audio.HasFile() {
		files[r.Audio.GetFormFieldName()] = r.Audio.GetFile()
	}

	return
}
