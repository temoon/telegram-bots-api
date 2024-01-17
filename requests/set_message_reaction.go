package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"io"
	"strconv"
)

type SetMessageReaction struct {
	ChatId    telegram.ChatId
	MessageId int64
	Reaction  []interface{}
	IsBig     *bool
}

func (r *SetMessageReaction) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMessageReaction", r, response)
	return
}

func (r *SetMessageReaction) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.Reaction != nil {
		var dataReaction []byte
		if dataReaction, err = json.Marshal(r.Reaction); err != nil {
			return
		}

		values["reaction"] = string(dataReaction)
	}

	if r.IsBig != nil {
		if *r.IsBig {
			values["is_big"] = "1"
		} else {
			values["is_big"] = "0"
		}
	}

	return
}

func (r *SetMessageReaction) GetFiles() (files map[string]io.Reader) {
	return
}
