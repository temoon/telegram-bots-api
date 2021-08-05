package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/go-telegram-bots-api"
	"strconv"
)

type EditMessageMedia struct {
	ChatId          interface{}
	InlineMessageId string
	Media           interface{}
	MessageId       int32
	ReplyMarkup     *telegram.InlineKeyboardMarkup
}

func (r *EditMessageMedia) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(interface{})
	err = b.CallMethod(ctx, "editMessageMedia", r, response)
	return
}

func (r *EditMessageMedia) IsMultipart() (multipart bool) {
	return true
}

func (r *EditMessageMedia) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.InlineMessageId != "" {
		values["inline_message_id"] = r.InlineMessageId
	}

	switch value := r.Media.(type) {
	case telegram.InputMediaAnimation:
		var dataInputMediaAnimation []byte
		if dataInputMediaAnimation, err = json.Marshal(value); err != nil {
			return
		}

		values["media"] = string(dataInputMediaAnimation)
	case telegram.InputMediaDocument:
		var dataInputMediaDocument []byte
		if dataInputMediaDocument, err = json.Marshal(value); err != nil {
			return
		}

		values["media"] = string(dataInputMediaDocument)
	case telegram.InputMediaAudio:
		var dataInputMediaAudio []byte
		if dataInputMediaAudio, err = json.Marshal(value); err != nil {
			return
		}

		values["media"] = string(dataInputMediaAudio)
	case telegram.InputMediaPhoto:
		var dataInputMediaPhoto []byte
		if dataInputMediaPhoto, err = json.Marshal(value); err != nil {
			return
		}

		values["media"] = string(dataInputMediaPhoto)
	case telegram.InputMediaVideo:
		var dataInputMediaVideo []byte
		if dataInputMediaVideo, err = json.Marshal(value); err != nil {
			return
		}

		values["media"] = string(dataInputMediaVideo)
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)
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
