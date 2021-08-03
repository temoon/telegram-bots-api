package requests

import (
	"strconv"
)

type SetGameScore struct {
	ChatId             uint64
	DisableEditMessage bool
	Force              bool
	InlineMessageId    string
	MessageId          uint64
	Score              uint64
	UserId             uint64
}

func (r *SetGameScore) IsMultipart() bool {
	return false
}

func (r *SetGameScore) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.ChatId != 0 {
		values["chat_id"] = strconv.FormatUint(r.ChatId, 10)
	}

	if r.DisableEditMessage {
		values["disable_edit_message"] = "1"
	}

	if r.Force {
		values["force"] = "1"
	}

	if r.InlineMessageId != "" {
		values["inline_message_id"] = r.InlineMessageId
	}

	if r.MessageId != 0 {
		values["message_id"] = strconv.FormatUint(r.MessageId, 10)
	}

	values["score"] = strconv.FormatUint(r.Score, 10)

	values["user_id"] = strconv.FormatUint(r.UserId, 10)

	return
}
