package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type PostStory struct {
	ActivePeriod         int64
	BusinessConnectionId string
	Content              interface{}
	Areas                []telegram.StoryArea
	Caption              *string
	CaptionEntities      []telegram.MessageEntity
	ParseMode            *string
	PostToChatPage       *bool
	ProtectContent       *bool
}

func (r *PostStory) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Story)
	err = b.CallMethod(ctx, "postStory", r, response)
	return
}

func (r *PostStory) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["active_period"] = strconv.FormatInt(r.ActivePeriod, 10)

	values["business_connection_id"] = r.BusinessConnectionId

	var dataContent []byte
	if dataContent, err = json.Marshal(r.Content); err != nil {
		return
	}

	values["content"] = string(dataContent)

	if r.Areas != nil {
		var dataAreas []byte
		if dataAreas, err = json.Marshal(r.Areas); err != nil {
			return
		}

		values["areas"] = string(dataAreas)
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

	if r.ParseMode != nil {
		values["parse_mode"] = *r.ParseMode
	}

	if r.PostToChatPage != nil {
		if *r.PostToChatPage {
			values["post_to_chat_page"] = "1"
		} else {
			values["post_to_chat_page"] = "0"
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

func (r *PostStory) GetFiles() (files map[string]io.Reader) {
	files = make(map[string]io.Reader)

	switch value := r.Content.(type) {
	case telegram.InputStoryContentPhoto:
		if value.Photo.HasFile() {
			files[value.Photo.GetFormFieldName()] = value.Photo.GetFile()
		}
	case telegram.InputStoryContentVideo:
		if value.Video.HasFile() {
			files[value.Video.GetFormFieldName()] = value.Video.GetFile()
		}
	}

	return
}
