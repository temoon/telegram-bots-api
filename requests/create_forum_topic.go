package requests

import (
	"context"
	"errors"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type CreateForumTopic struct {
	ChatId            interface{}
	IconColor         *int64
	IconCustomEmojiId *string
	Name              string
}

func (r *CreateForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ForumTopic)
	err = b.CallMethod(ctx, "createForumTopic", r, response)
	return
}

func (r *CreateForumTopic) IsMultipart() bool {
	return false
}

func (r *CreateForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	default:
		err = errors.New("invalid chat_id field type")
		return
	}

	if r.IconColor != nil {
		values["icon_color"] = strconv.FormatInt(*r.IconColor, 10)
	}

	if r.IconCustomEmojiId != nil {
		values["icon_custom_emoji_id"] = *r.IconCustomEmojiId
	}

	values["name"] = r.Name

	return
}
