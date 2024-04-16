package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditMessageMedia struct {
	Media           interface{}
	ChatId          *telegram.ChatId
	InlineMessageId *string
	MessageId       *int64
	ReplyMarkup     *telegram.InlineKeyboardMarkup
}

func (r *EditMessageMedia) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "editMessageMedia", r, response)
	return
}

func (r *EditMessageMedia) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	var dataMedia []byte
	if dataMedia, err = json.Marshal(r.Media); err != nil {
		return
	}

	values["media"] = string(dataMedia)

	if r.ChatId != nil {
		values["chat_id"] = r.ChatId.String()
	}

	if r.InlineMessageId != nil {
		values["inline_message_id"] = *r.InlineMessageId
	}

	if r.MessageId != nil {
		values["message_id"] = strconv.FormatInt(*r.MessageId, 10)
	}

	if r.ReplyMarkup != nil {
		var dataReplyMarkup []byte
		if dataReplyMarkup, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(dataReplyMarkup)
	}

	return
}

func (r *EditMessageMedia) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	switch value := r.Media.(type) {
	case telegram.InputMediaAnimation:
		if value.Media.HasFile() {
			files[value.Media.GetFormFieldName()] = value.Media.GetFile()
		}
		if value.Thumbnail != nil && value.Thumbnail.HasFile() {
			files[value.Thumbnail.GetFormFieldName()] = value.Thumbnail.GetFile()
		}
	case telegram.InputMediaDocument:
		if value.Media.HasFile() {
			files[value.Media.GetFormFieldName()] = value.Media.GetFile()
		}
		if value.Thumbnail != nil && value.Thumbnail.HasFile() {
			files[value.Thumbnail.GetFormFieldName()] = value.Thumbnail.GetFile()
		}
	case telegram.InputMediaAudio:
		if value.Media.HasFile() {
			files[value.Media.GetFormFieldName()] = value.Media.GetFile()
		}
		if value.Thumbnail != nil && value.Thumbnail.HasFile() {
			files[value.Thumbnail.GetFormFieldName()] = value.Thumbnail.GetFile()
		}
	case telegram.InputMediaPhoto:
		if value.Media.HasFile() {
			files[value.Media.GetFormFieldName()] = value.Media.GetFile()
		}
	case telegram.InputMediaVideo:
		if value.Media.HasFile() {
			files[value.Media.GetFormFieldName()] = value.Media.GetFile()
		}
		if value.Thumbnail != nil && value.Thumbnail.HasFile() {
			files[value.Thumbnail.GetFormFieldName()] = value.Thumbnail.GetFile()
		}
	}

	return
}
