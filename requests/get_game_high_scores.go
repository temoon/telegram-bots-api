package requests

import (
	"strconv"
)

type GetGameHighScores struct {
	ChatId          uint64
	InlineMessageId string
	MessageId       uint64
	UserId          uint64
}

func (r *GetGameHighScores) IsMultipart() bool {
	return false
}

func (r *GetGameHighScores) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ChatId != 0 {
		values["chat_id"] = strconv.FormatUint(r.ChatId, 10)
	}

	if r.InlineMessageId != "" {
		values["inline_message_id"] = r.InlineMessageId
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatUint(r.MessageId, 10)
	}

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
