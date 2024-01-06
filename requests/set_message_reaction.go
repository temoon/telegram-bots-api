package requests

import (
	"context"
	"encoding/json"
	"github.com/temoon/telegram-bots-api"
	"strconv"
)

type SetMessageReaction struct {
	ChatId    interface{}
	IsBig     *bool
	MessageId int32
	Reaction  interface{}
}

func (r *SetMessageReaction) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(bool)
	err = b.CallMethod(ctx, "setMessageReaction", r, response)
	return
}

func (r *SetMessageReaction) IsMultipart() (multipart bool) {
	return false
}

func (r *SetMessageReaction) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	switch value := r.ChatId.(type) {
	case int64:
		values["chat_id"] = strconv.FormatInt(value, 10)
	case string:
		values["chat_id"] = value
	}

	if r.IsBig != nil {
		if *r.IsBig {
			values["is_big"] = "1"
		} else {
			values["is_big"] = "0"
		}
	}

	values["message_id"] = strconv.FormatInt(int64(r.MessageId), 10)

	switch value := r.Reaction.(type) {
	case *telegram.ReactionTypeEmoji:
		if value != nil {
			var dataReactionTypeEmoji []byte
			if dataReactionTypeEmoji, err = json.Marshal(value); err != nil {
				return
			}

			values["reaction"] = string(dataReactionTypeEmoji)
		}
	case *telegram.ReactionTypeCustomEmoji:
		if value != nil {
			var dataReactionTypeCustomEmoji []byte
			if dataReactionTypeCustomEmoji, err = json.Marshal(value); err != nil {
				return
			}

			values["reaction"] = string(dataReactionTypeCustomEmoji)
		}
	}

	return
}
