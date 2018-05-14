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

func (r *SendMediaGroup) GetValues() (values map[string][]interface{}, err error) {
    values = make(map[string][]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = []interface{}{strconv.FormatInt(r.ChatID.(int64), 10)}
    case string:
        values["chat_id"] = []interface{}{r.ChatID.(string)}
    default:
        return nil, errors.New("invalid chat_id")
    }

    var data []byte
    if data, err = json.Marshal(r.Media); err != nil {
        return
    }

    values["media"] = []interface{}{string(data)}

    if r.DisableNotification == true {
        values["disable_notification"] = []interface{}{"1"}
    }

    values["reply_to_message_id"] = []interface{}{strconv.Itoa(r.ReplyToMessageID)}

    return
}
