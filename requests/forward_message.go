package requests

import (
    "errors"
    "strconv"
)

type ForwardMessage struct {
    ChatID              interface{}
    FromChatID          interface{}
    DisableNotification bool
    MessageID           uint32
}

func (r *ForwardMessage) IsMultipart() bool {
    return false
}

func (r *ForwardMessage) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(r.ChatID.(uint64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    switch r.FromChatID.(type) {
    case uint64:
        values["from_chat_id"] = strconv.FormatUint(r.FromChatID.(uint64), 10)
    case string:
        values["from_chat_id"] = r.FromChatID.(string)
    default:
        return nil, errors.New("invalid from_chat_id")
    }

    if r.DisableNotification {
        values["disable_notification"] = "1"
    }

    values["message_id"] = strconv.FormatUint(uint64(r.MessageID), 10)

    return
}
