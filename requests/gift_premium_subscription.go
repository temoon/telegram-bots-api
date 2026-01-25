package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type GiftPremiumSubscription struct {
	MonthCount    int64
	StarCount     int64
	UserId        int64
	Text          *string
	TextEntities  []telegram.MessageEntity
	TextParseMode *string
}

func (r *GiftPremiumSubscription) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "giftPremiumSubscription", r, response)
	return
}

func (r *GiftPremiumSubscription) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["month_count"] = strconv.FormatInt(r.MonthCount, 10)

	values["star_count"] = strconv.FormatInt(r.StarCount, 10)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

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

	return
}

func (r *GiftPremiumSubscription) GetFiles() (files map[string]io.Reader) {
	return
}
