package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type EditForumTopic struct {
	ChatId            telegram.ChatId
	MessageThreadId   int64
	IconCustomEmojiId *string
	Name              *string
}

func (r *EditForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "editForumTopic", r, response)
	return
}

func (r *EditForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["message_thread_id"] = strconv.FormatInt(r.MessageThreadId, 10)

	if r.IconCustomEmojiId != nil {
		values["icon_custom_emoji_id"] = *r.IconCustomEmojiId
	}

	if r.Name != nil {
		values["name"] = *r.Name
	}

	return
}

func (r *EditForumTopic) GetFiles() (files map[string]io.Reader) {
	return
}
