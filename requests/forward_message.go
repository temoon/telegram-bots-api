package requests

import (
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/temoon/telegram-bots-api"
)

type ForwardMessage struct {
	ChatId                  telegram.ChatId
	FromChatId              telegram.ChatId
	MessageId               int64
	DirectMessagesTopicId   *int64
	DisableNotification     *bool
	MessageEffectId         *string
	MessageThreadId         *int64
	ProtectContent          *bool
	SuggestedPostParameters *telegram.SuggestedPostParameters
	VideoStartTimestamp     *int64
}

func (r *ForwardMessage) Call(ctx context.Context, b *telegram.Bot) (response interface{}, err error) {
	response = new(telegram.Message)
	err = b.CallMethod(ctx, "forwardMessage", r, response)
	return
}

func (r *ForwardMessage) GetValues() (values map[string]interface{}, err error) {
	values = make(map[string]interface{})

	values["chat_id"] = r.ChatId.String()

	values["from_chat_id"] = r.FromChatId.String()

	values["message_id"] = strconv.FormatInt(r.MessageId, 10)

	if r.DirectMessagesTopicId != nil {
		values["direct_messages_topic_id"] = strconv.FormatInt(*r.DirectMessagesTopicId, 10)
	}

	if r.DisableNotification != nil {
		if *r.DisableNotification {
			values["disable_notification"] = "1"
		} else {
			values["disable_notification"] = "0"
		}
	}

	if r.MessageEffectId != nil {
		values["message_effect_id"] = *r.MessageEffectId
	}

	if r.MessageThreadId != nil {
		values["message_thread_id"] = strconv.FormatInt(*r.MessageThreadId, 10)
	}

	if r.ProtectContent != nil {
		if *r.ProtectContent {
			values["protect_content"] = "1"
		} else {
			values["protect_content"] = "0"
		}
	}

	if r.SuggestedPostParameters != nil {
		var dataSuggestedPostParameters []byte
		if dataSuggestedPostParameters, err = json.Marshal(r.SuggestedPostParameters); err != nil {
			return
		}

		values["suggested_post_parameters"] = string(dataSuggestedPostParameters)
	}

	if r.VideoStartTimestamp != nil {
		values["video_start_timestamp"] = strconv.FormatInt(*r.VideoStartTimestamp, 10)
	}

	return
}

func (r *ForwardMessage) GetFiles() (files map[string]io.Reader) {
	return
}
