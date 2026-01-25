package requests

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendAnimation struct {
	Animation               telegram.InputFile
	ChatId                  telegram.ChatId
	AllowPaidBroadcast      *bool
	BusinessConnectionId    *string
	Caption                 *string
	CaptionEntities         []telegram.MessageEntity
	DirectMessagesTopicId   *int64
	DisableNotification     *bool
	Duration                *int64
	HasSpoiler              *bool
	Height                  *int64
	MessageEffectId         *string
	MessageThreadId         *int64
	ParseMode               *string
	ProtectContent          *bool
	ReplyMarkup             interface{}
	ReplyParameters         *telegram.ReplyParameters
	ShowCaptionAboveMedia   *bool
	SuggestedPostParameters *telegram.SuggestedPostParameters
	Thumbnail               *telegram.InputFile
	Width                   *int64
}

func (r *SendAnimation) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "sendAnimation", r, response)
	return
}

func (r *SendAnimation) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["animation"] = r.Animation.String()

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

	if r.MessageEffectId != nil {
		values["message_effect_id"] = *r.MessageEffectId
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

	if r.ShowCaptionAboveMedia != nil {
		if *r.ShowCaptionAboveMedia {
			values["show_caption_above_media"] = "1"
		} else {
			values["show_caption_above_media"] = "0"
		}
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

	if r.Width != nil {
		values["width"] = strconv.FormatInt(*r.Width, 10)
	}

	return
}

func (r *SendAnimation) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	if r.Animation.HasFile() {
		files[r.Animation.GetFormFieldName()] = r.Animation.GetFile()
	}
	if r.Thumbnail != nil && r.Thumbnail.HasFile() {
		files[r.Thumbnail.GetFormFieldName()] = r.Thumbnail.GetFile()
	}

	return
}
