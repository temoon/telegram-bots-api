package requests

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SendMediaGroup struct {
	ChatId              telegram.ChatId
	Media               interface{}
	DisableNotification *bool
	MessageThreadId     *int64
	ProtectContent      *bool
	ReplyParameters     *telegram.ReplyParameters
}

func (r *SendMediaGroup) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new([]telegram.Message)
	err = b.CallMethod(ctx, "sendMediaGroup", r, response)
	return
}

func (r *SendMediaGroup) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	switch value := r.Media.(type) {
	case []telegram.InputMediaAudio, []telegram.InputMediaDocument, []telegram.InputMediaPhoto, []telegram.InputMediaVideo:
		var data []byte
		if data, err = json.Marshal(value); err != nil {
			return
		}

		values["media"] = string(data)
	default:
		err = errors.New("unsupported media field type")
		return
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
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

	return
}

func (r *SendMediaGroup) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	switch value := r.Media.(type) {
	case []telegram.InputMediaAudio:
		for _, item := range value {
			if item.Thumbnail != nil && item.Thumbnail.HasFile() {
				files[item.Thumbnail.GetFormFieldName()] = item.Thumbnail.GetFile()
			}
			if item.Media.HasFile() {
				files[item.Media.GetFormFieldName()] = item.Media.GetFile()
			}
		}
	case []telegram.InputMediaDocument:
		for _, item := range value {
			if item.Media.HasFile() {
				files[item.Media.GetFormFieldName()] = item.Media.GetFile()
			}
			if item.Thumbnail != nil && item.Thumbnail.HasFile() {
				files[item.Thumbnail.GetFormFieldName()] = item.Thumbnail.GetFile()
			}
		}
	case []telegram.InputMediaPhoto:
		for _, item := range value {
			if item.Media.HasFile() {
				files[item.Media.GetFormFieldName()] = item.Media.GetFile()
			}
		}
	case []telegram.InputMediaVideo:
		for _, item := range value {
			if item.Media.HasFile() {
				files[item.Media.GetFormFieldName()] = item.Media.GetFile()
			}
			if item.Thumbnail != nil && item.Thumbnail.HasFile() {
				files[item.Thumbnail.GetFormFieldName()] = item.Thumbnail.GetFile()
			}
		}
	}
	return
}
