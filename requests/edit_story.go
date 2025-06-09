package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditStory struct {
	BusinessConnectionId string
	Content              interface{}
	StoryId              int64
	Areas                []telegram.StoryArea
	Caption              *string
	CaptionEntities      []telegram.MessageEntity
	ParseMode            *string
}

func (r *EditStory) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Story)
	err = b.CallMethod(ctx, "editStory", r, response)
	return
}

func (r *EditStory) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	var dataContent []byte
	if dataContent, err = json.Marshal(r.Content); err != nil {
		return
	}

	values["content"] = string(dataContent)

	values["story_id"] = strconv.FormatInt(r.StoryId, 10)

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

	return
}

func (r *EditStory) GetFiles() (files map[string]io.Reader) {
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
