package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type SavePreparedInlineMessage struct {
	Result            interface{}
	UserId            int64
	AllowBotChats     *bool
	AllowChannelChats *bool
	AllowGroupChats   *bool
	AllowUserChats    *bool
}

func (r *SavePreparedInlineMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.PreparedInlineMessage)
	err = b.CallMethod(ctx, "savePreparedInlineMessage", r, response)
	return
}

func (r *SavePreparedInlineMessage) GetValues() (values map[string]string, err error) {
	values = make(map[string]string)

	var dataResult []byte
	if dataResult, err = json.Marshal(r.Result); err != nil {
		return
	}

	values["result"] = string(dataResult)

	values["user_id"] = strconv.FormatInt(r.UserId, 10)

	if r.AllowBotChats != nil {
		if *r.AllowBotChats {
			values["allow_bot_chats"] = "1"
		} else {
			values["allow_bot_chats"] = "0"
		}
	}

	if r.AllowChannelChats != nil {
		if *r.AllowChannelChats {
			values["allow_channel_chats"] = "1"
		} else {
			values["allow_channel_chats"] = "0"
		}
	}

	if r.AllowGroupChats != nil {
		if *r.AllowGroupChats {
			values["allow_group_chats"] = "1"
		} else {
			values["allow_group_chats"] = "0"
		}
	}

	if r.AllowUserChats != nil {
		if *r.AllowUserChats {
			values["allow_user_chats"] = "1"
		} else {
			values["allow_user_chats"] = "0"
		}
	}

	return
}

func (r *SavePreparedInlineMessage) GetFiles() (files map[string]io.Reader) {
	return
}
