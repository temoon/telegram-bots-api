package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type DeleteStory struct {
	BusinessConnectionId string
	StoryId              int64
}

func (r *DeleteStory) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "deleteStory", r, response)
	return
}

func (r *DeleteStory) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["business_connection_id"] = r.BusinessConnectionId

	values["story_id"] = strconv.FormatInt(r.StoryId, 10)

	return
}

func (r *DeleteStory) GetFiles() (files map[string]io.Reader) {
	return
}
