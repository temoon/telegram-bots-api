package requests

import (
    "encoding/json"
    "errors"
    "strconv"
)

type SendMediaGroup struct {
    ChatID              interface{}
    Media               []interface{}
    DisableNotification bool
    ReplyToMessageID    uint32
}

func (r *SendMediaGroup) IsMultipart() bool {
    return false
}

func (r *SendMediaGroup) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch chatID := r.ChatID.(type) {
    case uint64:
        values["chat_id"] = strconv.FormatUint(chatID, 10)
    case string:
        values["chat_id"] = chatID
    default:
        return nil, errors.New("invalid chat_id")
    }

    var data []byte
    if data, err = json.Marshal(r.Media); err != nil {
        return
    }

    values["media"] = string(data)

    if r.DisableNotification {
        values["disable_notification"] = "1"
    }

    values["reply_to_message_id"] = strconv.FormatUint(uint64(r.ReplyToMessageID), 10)

    return
}
