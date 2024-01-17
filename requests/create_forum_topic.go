package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type CreateForumTopic struct {
	ChatId            telegram.ChatId
	Name              string
	IconColor         *int64
	IconCustomEmojiId *string
}

func (r *CreateForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.ForumTopic)
	err = b.CallMethod(ctx, "createForumTopic", r, response)
	return
}

func (r *CreateForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["name"] = r.Name

	if r.IconColor != nil {
		values["icon_color"] = strconv.FormatInt(*r.IconColor, 10)
	}

	if r.IconCustomEmojiId != nil {
		values["icon_custom_emoji_id"] = *r.IconCustomEmojiId
	}

	return
}

func (r *CreateForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
