package requests

import (
    "errors"
    "strconv"
)

type SendChatAction struct {
    ChatID interface{}
    Action string
}

func (r *SendChatAction) IsMultipart() bool {
    return false
}

func (r *SendChatAction) GetValues() (values map[string]interface{}, err error) {
    values = make(map[string]interface{})

    switch r.ChatID.(type) {
    case int64:
        values["chat_id"] = strconv.FormatInt(r.ChatID.(int64), 10)
    case string:
        values["chat_id"] = r.ChatID.(string)
    default:
        return nil, errors.New("invalid chat_id")
    }

    values["action"] = r.Action

    return
}
