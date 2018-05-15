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
    ReplyToMessageID    int
}

func (r *SendMediaGroup) IsMultipart() bool {
    return false
}

func (r *SendMediaGroup) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = strconv.FormatInt(r.ChatID.(int64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
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

    values["reply_to_message_id"] = strconv.Itoa(r.ReplyToMessageID)

    return
}
