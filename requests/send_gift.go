package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SendGift struct {
	GiftId        string
	ChatId        *telegram.ChatId
	PayForUpgrade *bool
	Text          *string
	TextEntities  []telegram.MessageEntity
	TextParseMode *string
	UserId        *int64
}

func (r *SendGift) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "sendGift", r, response)
	return
}

func (r *SendGift) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["gift_id"] = r.GiftId

	if r.ChatId != nil {
		values["chat_id"] = r.ChatId.String()
	}

	if r.PayForUpgrade != nil {
		if *r.PayForUpgrade {
			values["pay_for_upgrade"] = "1"
		} else {
			values["pay_for_upgrade"] = "0"
		}
	}

	if r.Text != nil {
		values["text"] = *r.Text
	}

	if r.TextEntities != nil {
		var dataTextEntities []byte
		if dataTextEntities, err = json.Marshal(r.TextEntities); err != nil {
			return
		}

		values["text_entities"] = string(dataTextEntities)
	}

	if r.TextParseMode != nil {
		values["text_parse_mode"] = *r.TextParseMode
	}

	if r.UserId != nil {
		values["user_id"] = strconv.FormatInt(*r.UserId, 10)
	}

	return
}

func (r *SendGift) GetFiles() (files map[string]io.Reader) {
	return
}
