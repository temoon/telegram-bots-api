package requests

import (
	"context"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetUserEmojiStatus struct {
	UserId                    int64
	EmojiStatusCustomEmojiId  *string
	EmojiStatusExpirationDate *int64
}

func (r *SetUserEmojiStatus) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setUserEmojiStatus", r, response)
	return
}

func (r *SetUserEmojiStatus) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.EmojiStatusCustomEmojiId != nil {
		values["emoji_status_custom_emoji_id"] = *r.EmojiStatusCustomEmojiId
	}

	if r.EmojiStatusExpirationDate != nil {
		values["emoji_status_expiration_date"] = strconv.FormatInt(*r.EmojiStatusExpirationDate, 10)
	}

	return
}

func (r *SetUserEmojiStatus) GetFiles() (files map[string]io.Reader) {
	return
}
