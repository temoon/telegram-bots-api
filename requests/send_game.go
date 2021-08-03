package requests

import (
	"encoding/json"
	"strconv"
)

type SendGame struct {
	AllowSendingWithoutReply bool
	ChatId                   uint64
	DisableNotification      bool
	GameShortName            string
	ReplyMarkup              interface{}
	ReplyToMessageId         uint64
}

func (r *SendGame) IsMultipart() bool {
	return false
}

func (r *SendGame) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	if r.AllowSendingWithoutReply {
		values["allow_sending_without_reply"] = "1"
	}

	values["chat_id"] = strconv.FormatUint(r.ChatId, 10)

	if r.DisableNotification {
		values["disable_notification"] = "1"
	}

	values["game_short_name"] = r.GameShortName

	if r.ReplyMarkup != nil {
		var data []byte
		if data, err = json.Marshal(r.ReplyMarkup); err != nil {
			return
		}

		values["reply_markup"] = string(data)
	}

	if r.ReplyToMessageId != 0 {
		values["reply_to_message_id"] = strconv.FormatUint(r.ReplyToMessageId, 10)
	}

	return
}
