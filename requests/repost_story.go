package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type RepostStory struct {
	ActivePeriod         int64
	BusinessConnectionId string
	FromChatId           int64
	FromStoryId          int64
	PostToChatPage       *bool
	ProtectContent       *bool
}

func (r *RepostStory) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Story)
	err = b.CallMethod(ctx, "repostStory", r, response)
	return
}

func (r *RepostStory) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["active_period"] = strconv.FormatInt(r.ActivePeriod, 10)

	values["business_connection_id"] = r.BusinessConnectionId

	values["from_chat_id"] = strconv.FormatInt(r.FromChatId, 10)

	values["from_story_id"] = strconv.FormatInt(r.FromStoryId, 10)

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

func (r *RepostStory) GetFiles() (files map[string]io.Reader) {
	return
}
