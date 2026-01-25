package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SetMessageReaction struct {
	ChatId    telegram.ChatId
	MessageId int64
	IsBig     *bool
	Reaction  []interface{}
}

func (r *SetMessageReaction) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMessageReaction", r, response)
	return
}

func (r *SetMessageReaction) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	values["chat_id"] = r.ChatId.String()

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.IsBig != nil {
		if *r.IsBig {
			values["is_big"] = "1"
		} else {
			values["is_big"] = "0"
		}
	}

	if r.Reaction != nil {
		var dataReaction []byte
		if dataReaction, err = json.Marshal(r.Reaction); err != nil {
			return
		}

		values["reaction"] = string(dataReaction)
	}

	return
}

func (r *SetMessageReaction) GetFiles() (files map[string]io.Reader) {
	return
}
