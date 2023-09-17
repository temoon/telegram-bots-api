package requests

import (
	"context"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type EditForumTopic struct {
	ChatId            interface{}
	IconCustomEmojiId *string
	MessageThreadId   int32
	Name              *string
}

func (r *EditForumTopic) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "editForumTopic", r, response)
	return
}

func (r *EditForumTopic) IsMultipart() (multipart bool) {
	return false
}

func (r *EditForumTopic) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.IconCustomEmojiId != nil {
		values["icon_custom_emoji_id"] = *r.IconCustomEmojiId
	}

	values["message_thread_id"] = strconv.FormatInt(int64(r.MessageThreadId), 10)

	if r.Name != nil {
		values["name"] = *r.Name
	}

	return
}
